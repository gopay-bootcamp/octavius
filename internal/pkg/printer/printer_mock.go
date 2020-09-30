package printer

import (
	"github.com/fatih/color"
	"github.com/stretchr/testify/mock"
)

// MockPrinter mock struct
type MockPrinter struct {
	mock.Mock
}

// Println mock
func (m *MockPrinter) Println(s string, attr ...color.Attribute) {
	m.Called(s)
}
