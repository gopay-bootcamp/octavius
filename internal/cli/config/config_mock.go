package config

import (
	"github.com/stretchr/testify/mock"
)

type MockLoader struct {
	mock.Mock
}

func (m *MockLoader) Load() (OctaviusConfig, ConfigError) {
	args := m.Called()
	return args.Get(0).(OctaviusConfig), args.Get(1).(ConfigError)
}
