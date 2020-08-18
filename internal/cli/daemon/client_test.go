package daemon

import (
	"github.com/stretchr/testify/suite"
	"octavius/internal/cli/config"
	"testing"
)

type ClientTestSuite struct {
	suite.Suite
	testClient       Client
	mockConfig 		config.MockConfig
//	mockConfigLoader *config.MockLoader
}

func (s *ClientTestSuite) SetupTest() {
	s.testClient = NewClient(config.NewLoader())
}

func (s *ClientTestSuite) TestCreateMetadata() {
	s.mockConfig.On("Load").Once()
	s.testClient.CreateMetadata("../../../test/metadata/metadata.json")

}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

