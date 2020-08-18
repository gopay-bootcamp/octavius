package daemon

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"octavius/internal/cli/client"
	"octavius/internal/cli/config"
	"os"
	"testing"
	"time"
)

type ClientTestSuite struct {
	suite.Suite
	testClient       Client
	mockLoader 		config.MockLoader
	mockGrpcClient 	client.MockGrpcClient
}

func (s *ClientTestSuite) SetupTest() {
	s.mockGrpcClient=client.MockGrpcClient{}
	s.mockLoader=config.MockLoader{}


	s.testClient=NewClient(&s.mockLoader)


}

func (s *ClientTestSuite) TestCreateMetadata() {

	s.mockLoader.On("Load").Return(
		config.OctaviusConfig{
			Host           :  "localhost:5050",
			Email                 : "jaimin.rathod@go-jek.com",
			AccessToken           :	"AllowMe",
			ConnectionTimeoutSecs : time.Second,
		},
		config.ConfigError{},
		).Once()

	s.mockGrpcClient.On("NewGrpcClient","localhost:5050").Once()

	metadataFileHandler, err := os.Open("../../../test/metadata/metadata.json")
	if err != nil {
		fmt.Println("Error opening the file given")
		return
	}
	defer metadataFileHandler.Close()
	s.testClient.CreateMetadata(metadataFileHandler,&s.mockGrpcClient)

}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}


