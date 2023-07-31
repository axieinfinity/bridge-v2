package stats

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type NodeInfo struct {
	Organization  string `json:"organization,omitempty" mapstructure:"organization"`
	Coinbase      string `json:"coinbase" mapstructure:"coinbase"`
	Name          string `json:"name" mapstructure:"name"`
	Node          string `json:"node" mapstructure:"node"`
	Port          int    `json:"port" mapstructure:"port"`
	Network       string `json:"net" mapstructure:"net"`
	Protocol      string `json:"protocol" mapstructure:"protocol"`
	API           string `json:"api" mapstructure:"api"`
	Os            string `json:"os" mapstructure:"os"`
	OsVer         string `json:"os_v" mapstructure:"os_v"`
	Client        string `json:"client" mapstructure:"client"`
	History       bool   `json:"canUpdateHistory" mapstructure:"canUpdateHistory"`
	Operator      string `json:"operator" mapstructure:"operator"`
	BridgeVersion string `json:"bridgeVersion" mapstructure:"bridgeVersion"`
}

var BridgeStats *Service

func SendErrorToStats(listener bridgeCore.Listener, err error) {
	if err != nil && BridgeStats != nil {
		BridgeStats.SendError(listener.GetName(), err.Error())
	}
}

// connWrapper is a wrapper to prevent concurrent-write or concurrent-read on the
// websocket.
//
// From Gorilla websocket docs:
//
//	Connections support one concurrent reader and one concurrent writer.
//	Applications are responsible for ensuring that no more than one goroutine calls the write methods
//	  - NextWriter, SetWriteDeadline, WriteMessage, WriteJSON, EnableWriteCompression, SetCompressionLevel
//	concurrently and that no more than one goroutine calls the read methods
//	  - NextReader, SetReadDeadline, ReadMessage, ReadJSON, SetPongHandler, SetPingHandler
//	concurrently.
//	The Close and WriteControl methods can be called concurrently with all other methods.
type connWrapper struct {
	conn *websocket.Conn

	rlock sync.Mutex
	wlock sync.Mutex
}

func newConnectionWrapper(conn *websocket.Conn) *connWrapper {
	return &connWrapper{conn: conn}
}

// WriteJSON wraps corresponding method on the websocket but is safe for concurrent calling
func (w *connWrapper) WriteJSON(v interface{}) error {
	w.wlock.Lock()
	defer w.wlock.Unlock()

	return w.conn.WriteJSON(v)
}

// ReadJSON wraps corresponding method on the websocket but is safe for concurrent calling
func (w *connWrapper) ReadJSON(v interface{}) error {
	w.rlock.Lock()
	defer w.rlock.Unlock()

	return w.conn.ReadJSON(v)
}

// Close wraps corresponding method on the websocket but is safe for concurrent calling
func (w *connWrapper) Close() error {
	// The Close and WriteControl methods can be called concurrently with all other methods,
	// so the mutex is not used here
	return w.conn.Close()
}

type errorMessage struct {
	Listener string `json:"listener"`
	Err      string `json:"error"`
}

type processedBlockMessage struct {
	Listener       string `json:"listener"`
	ProcessedBlock uint64 `json:"processedBlock"`
}

type BridgeInfo struct {
	Node           string            `json:"node"`
	Operator       string            `json:"bridgeOperatorAddress"`
	Version        string            `json:"version"`
	LastError      map[string]string `json:"lastError"`
	ProcessedBlock map[string]uint64 `json:"processedBlock"`
	PendingTasks   int               `json:"pendingTasks"`
	FailedTasks    int               `json:"failedTasks"`
}

type authMsg struct {
	ID     string    `json:"id"`
	Info   *NodeInfo `json:"info"`
	Secret string    `json:"secret"`
}

type Service struct {
	lock             sync.Mutex
	chainId          string
	node             string
	operator         string
	version          string
	secret           string
	host             string
	lastError        map[string]string
	processedBlock   map[string]uint64
	errCh            chan errorMessage
	processedBlockCh chan processedBlockMessage
	quitCh           chan struct{}
	store            stores.TaskStore
}

func NewService(node, chainId, operator, host, secret string, db *gorm.DB) {
	BridgeStats = &Service{
		node:             node,
		chainId:          chainId,
		operator:         operator,
		secret:           secret,
		host:             host,
		store:            stores.NewTaskStore(db),
		lastError:        make(map[string]string),
		processedBlock:   make(map[string]uint64),
		errCh:            make(chan errorMessage, 1),
		processedBlockCh: make(chan processedBlockMessage, 1),
		quitCh:           make(chan struct{}, 1),
	}
}

// Start loop keeps trying to connect to ronin stats server, reporting bridge stats.
func (s *Service) Start() {
	errTimer := time.NewTimer(0)
	defer errTimer.Stop()
	// Loop reporting until termination
	for {
		select {
		case <-s.quitCh:
			return
		case <-errTimer.C:
			// Establish a websocket connection to the server on any supported URL
			var (
				conn *connWrapper
				err  error
			)
			dialer := websocket.Dialer{HandshakeTimeout: 5 * time.Second}
			header := make(http.Header)
			header.Set("origin", "http://localhost")
			log.Info("[Bridge stats] Dial to host ", "host", s.host)
			c, _, e := dialer.Dial(s.host, header)
			err = e
			if err == nil {
				conn = newConnectionWrapper(c)
			}
			if err != nil {
				log.Warn("Stats server unreachable", "err", err)
				errTimer.Reset(10 * time.Second)
				continue
			}

			// Authenticate the client with the server
			if err = s.login(conn); err != nil {
				log.Warn("Stats login failed", "err", err)
				conn.Close()
				errTimer.Reset(10 * time.Second)
				continue
			}
			go s.readLoop(conn)

			sendStatsTicker := time.NewTicker(10 * time.Second)
			for err == nil {
				select {
				case <-s.quitCh:
					sendStatsTicker.Stop()
					conn.Close()
					return
				case msg := <-s.errCh:
					err = s.setLastError(msg.Listener, msg.Err)
				case msg := <-s.processedBlockCh:
					err = s.setProcessedBlock(msg.Listener, msg.ProcessedBlock)
				case <-sendStatsTicker.C: // Checking stats interval
					if err = s.report(conn); err != nil {
						log.Warn("bridge stats report failed", "err", err)
						// When bridge stats report failed need to relogn after 10 seconds
					}
				}
			}
			sendStatsTicker.Stop()
			log.Warn("[Bridge stats] Redial connection")
			// Close the current connection and establish a new one
			conn.Close()
			errTimer.Reset(0)
		}
	}
}

func (s *Service) Stop() {
	close(s.quitCh)
}

func (s *Service) login(conn *connWrapper) error {
	// Construct and send the login authentication
	auth := &authMsg{
		ID: s.node,
		Info: &NodeInfo{
			Node:          s.node,
			Operator:      s.operator,
			BridgeVersion: s.version,
		},
		Secret: s.secret,
	}
	login := map[string][]interface{}{
		"emit": {"hello", auth},
	}
	if err := conn.WriteJSON(login); err != nil {
		return err
	}
	// Retrieve the remote ack or connection termination
	var ack map[string][]string
	if err := conn.ReadJSON(&ack); err != nil || len(ack["emit"]) != 1 || ack["emit"][0] != "ready" {
		return errors.New("unauthorized")
	}
	return nil
}

func (s *Service) report(conn *connWrapper) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	// collect bridge stats and send to ronin stats
	info := &BridgeInfo{
		Node:           s.node,
		Operator:       s.operator,
		Version:        "v2",
		LastError:      s.lastError,
		ProcessedBlock: s.processedBlock,
	}

	// count pending/failed tasks from database
	pending, err := s.store.CountTasks(s.chainId, "pending")
	if err != nil {
		log.Error("error while getting pending tasks", "err", err)
	} else {
		info.PendingTasks = int(pending)
	}

	failed, err := s.store.CountTasks(s.chainId, "failed")
	if err != nil {
		log.Error("error while getting failed tasks", "err", err)
	} else {
		info.FailedTasks = int(failed)
	}

	report := map[string][]interface{}{
		"emit": {"bridge-stats", info},
	}
	log.Info("[Bridge-Stats] Emit Bridge stats")
	return conn.WriteJSON(report)
}

// readLoop loops as long as the connection is alive and retrieves data packets
// from the network socket. If any of them match an active request, it forwards
// it, if they themselves are requests it initiates a reply, and lastly it drops
// unknown packets.
func (s *Service) readLoop(conn *connWrapper) {
	// If the read loop exits, close the connection
	defer conn.Close()
	for {
		select {
		case <-s.quitCh:
			return
		default:
		}
		// Retrieve the next generic network packet and bail out on error
		var blob json.RawMessage
		if err := conn.ReadJSON(&blob); err != nil {
			log.Warn("Failed to retrieve stats server message", "err", err)
			return
		}
		// If the network packet is a system ping, respond to it directly
		var ping string
		if err := json.Unmarshal(blob, &ping); err == nil && strings.HasPrefix(ping, "primus::ping::") {
			if err := conn.WriteJSON(strings.Replace(ping, "ping", "pong", -1)); err != nil {
				log.Warn("Failed to respond to system ping message", "err", err)
				return
			}
			continue
		}
	}
}

func (s *Service) setLastError(listener, err string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if err != "" {
		s.lastError[listener] = err
	}
	return nil
}

func (s *Service) setProcessedBlock(listener string, block uint64) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.processedBlock[listener] < block {
		s.processedBlock[listener] = block
	}
	return nil
}

func (s *Service) SendError(listener, err string) {
	s.errCh <- errorMessage{Listener: listener, Err: err}
}

func (s *Service) SendProcessedBlock(listener string, block uint64) {
	s.processedBlockCh <- processedBlockMessage{Listener: listener, ProcessedBlock: block}
}
