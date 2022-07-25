package prometheus

import (
	"context"
	"sync"
	"time"

	"github.com/axieinfinity/bridge-v2/configs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type PusherAdapter struct {
	counters map[string]prometheus.Counter
	gauges   map[string]prometheus.Gauge
	registry *prometheus.Registry
	pusher   *push.Pusher
	lock     sync.RWMutex
}

func (p *PusherAdapter) AddCounter(name string, description string) *PusherAdapter {
	p.lock.Lock()
	defer p.lock.Unlock()

	if _, ok := p.counters[name]; ok {
		return p
	}

	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: description,
	})
	p.counters[name] = counter
	p.pusher.Collector(counter)
	return p
}

func (p *PusherAdapter) IncrCounter(name string, value int) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	if _, ok := p.counters[name]; !ok {
		return
	}
	p.counters[name].Add(float64(value))
}

func (p *PusherAdapter) AddGauge(name string, description string) *PusherAdapter {
	p.lock.Lock()
	defer p.lock.Unlock()

	if _, ok := p.gauges[name]; ok {
		return p
	}

	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: description,
	})
	p.pusher.Collector(gauge)
	return p
}

func (p *PusherAdapter) IncrGauge(name string, value int) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	if _, ok := p.gauges[name]; !ok {
		p.AddCounter(name, "")
	}

	p.gauges[name].Add(float64(value))
}

func (p *PusherAdapter) Push() error {
	return p.pusher.Push()
}

func (p *PusherAdapter) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Second * time.Duration(configs.AppConfig.Prometheus.PushInterval))
	select {
	case <-ticker.C:
		p.pusher.Push()
	case <-ctx.Done():
		ticker.Stop()
	}
}

func NewPusher() *PusherAdapter {
	pusher := push.New(configs.AppConfig.Prometheus.PushURL, configs.AppConfig.Prometheus.PushJob)
	return &PusherAdapter{
		pusher:   pusher,
		counters: make(map[string]prometheus.Counter),
		gauges:   make(map[string]prometheus.Gauge),
	}
}
