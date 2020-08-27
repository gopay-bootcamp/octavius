package file

import (
	"io"

	"github.com/stretchr/testify/mock"
)

// MockFileUtil mock
type MockFileUtil struct {
	mock.Mock
}

// GetUserInput mock
func (m *MockFileUtil) GetUserInput() (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

// GetIoReader mock
func (m *MockFileUtil) GetIoReader(filePath string) (io.Reader, error) {
	args := m.Called(filePath)
	return args.Get(0).(io.Reader), args.Error(1)
}

// IsFileExist mock
func (m *MockFileUtil) IsFileExist(filePath string) bool {
	args := m.Called(filePath)
	return args.Get(0).(bool)
}

// ReadFile mock
func (m *MockFileUtil) ReadFile(filePath string) (string, error) {
	args := m.Called(filePath)
	return args.Get(0).(string), args.Error(1)
}

// CreateDirIfNotExist mock
func (m *MockFileUtil) CreateDirIfNotExist(filePath string) error {
	args := m.Called(filePath)
	return args.Error(0)
}

// CreateFile mock
func (m *MockFileUtil) CreateFile(filepath string) error {
	args := m.Called(filepath)
	return args.Error(0)
}

// WriteFile mock
func (m *MockFileUtil) WriteFile(filepath string, content string) error {
	args := m.Called(filepath, content)
	return args.Error(0)
}
