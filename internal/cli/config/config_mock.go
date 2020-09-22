package config

import (
	"github.com/stretchr/testify/mock"
)

//MockLoader is a mock for config
type MockLoader struct {
	mock.Mock
}

//Load mock
func (m *MockLoader) Load() (OctaviusConfig, ConfigError) {
	args := m.Called()
	return args.Get(0).(OctaviusConfig), args.Get(1).(ConfigError)
}
