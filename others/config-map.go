package main

import (
	"errors"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Endpoint string
}

type ConfigProvider interface {
	GetServiceConfig(serviceName string) (*Config, error)
}




type YamlConfigProvider struct {
	Serivces map[string]*Config `yaml:"services"`
}

func NewYamlConfigProvider(filepath string) (*YamlConfigProvider, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return  nil, err
	}

	ycp := &YamlConfigProvider{}
	err = yaml.Unmarshaler(content, err)
	return ycp, err
}

func (y *YamlConfigProvider) GetServiceConfig(serviceName string) (*main2.Config, error) {
	cfg, ok := y.Serivces[serviceName]
	if !ok {
		return nil,  errors.New("service not found")
	}
	return  cfg, nil
}
