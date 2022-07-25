package configs

import "github.com/kelseyhightower/envconfig"

var AppConfig Config

type Config struct {
	Prometheus Prometheus
}

type Prometheus struct {
	PushURL      string `default:"localhost:9091" envconfig:"PUSH_URL"`
	PushJob      string `default:"sm-bridge-v2" envconfig:"PUSH_JOB"`
	PushInterval int    `default:"15" envconfig:"PUSH_INTERVAl"`
}

func New() (*Config, error) {
	if err := envconfig.Process("", &AppConfig); err != nil {
		return nil, err
	}
	return &AppConfig, nil
}
