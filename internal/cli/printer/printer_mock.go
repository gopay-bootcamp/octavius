package printer

import (
	"github.com/fatih/color"
	"github.com/stretchr/testify/mock"
)

type MockPrinter struct {
	mock.Mock
}

func (m *MockPrinter) Println(s string, attr ...color.Attribute) {
	m.Called(s)
}
