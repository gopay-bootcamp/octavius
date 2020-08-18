package config

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) Load() (OctaviusConfig, ConfigError) {
	m.Called()
	OctaviusConfig:=OctaviusConfig{
		Host: "localhost:5050",
		Email: "xyz@gmail.com",
		AccessToken: "Token",
		ConnectionTimeoutSecs: time.Second,
	}
	return OctaviusConfig , ConfigError{}

}
