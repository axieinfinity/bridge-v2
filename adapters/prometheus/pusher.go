package prometheus

import (
	"context"
	"time"

	"github.com/axieinfinity/bridge-v2/configs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Pusher struct {
	counters map[string]prometheus.Counter
	gauges   map[string]prometheus.Gauge
	registry *prometheus.Registry
	pusher   *push.Pusher
}

func (p *Pusher) AddCounter(name string, description string) *Pusher {
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

func (p *Pusher) IncrCounter(name string, value int) {
	if _, ok := p.counters[name]; !ok {
		return
	}
	p.counters[name].Add(float64(value))
}

func (p *Pusher) AddGauge(name string, description string) *Pusher {
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

func (p *Pusher) IncrGauge(name string, value int) {
	if _, ok := p.gauges[name]; !ok {
		return
	}

	p.gauges[name].Add(float64(value))
}

func (p *Pusher) Push() error {
	return p.pusher.Push()
}

func (p *Pusher) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Second * time.Duration(configs.AppConfig.Prometheus.PushInterval))
	select {
	case <-ticker.C:
		p.pusher.Push()
	case <-ctx.Done():
		ticker.Stop()
	}
}

func NewPusher() *Pusher {
	pusher := push.New(configs.AppConfig.Prometheus.PushURL, configs.AppConfig.Prometheus.PushJob)
	return &Pusher{
		pusher:   pusher,
		counters: make(map[string]prometheus.Counter),
		gauges:   make(map[string]prometheus.Gauge),
	}
}
