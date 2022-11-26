package conf

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configYaml string

type Config struct {
	TcpProtocolAddr       string `yaml:"tcp_protocol_addr"`
	CustomTcpProtocolAddr string `yaml:"custom_tcp_protocol_addr"`
	Mqtt                  struct {
		Addr     string `yaml:"addr"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Qos      int    `yaml:"qos"`
		Retain   bool   `yaml:"retain"`
	} `yaml:"mqtt"`
	Topic struct {
		Subscribe        string `yaml:"subscribe"`
		SubscribeRawData string `yaml:"subscribe_raw_data"`
		PublishRawData   string `yaml:"publish_raw_data"`
	} `yaml:"topic"`
	Api struct {
		CustomAddr string `yaml:"custom_addr"`
		SelfAddr   string `yaml:"self_addr"`
	} `yaml:"api"`
}

func LoadConfig(key string) (*Config, error) {
	m := make(map[string]*Config)
	err := yaml.Unmarshal([]byte(configYaml), &m)
	if err != nil {
		return nil, err
	}

	return m[key], nil
}
