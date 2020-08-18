package config

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type MockLoader struct {
	mock.Mock
}



func (m *MockLoader) Load() (OctaviusConfig, ConfigError) {
	fmt.Printf("Called load mock")
	args:= m.Called()

	return args.Get(0).(OctaviusConfig), args.Get(1).(ConfigError)
}
