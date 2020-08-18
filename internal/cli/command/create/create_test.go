package create

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"octavius/internal/cli/daemon"
	"testing"
)

type CreateCmdTestSuite struct {
	suite.Suite
	mockOctaviusDClient *daemon.MockClient
	testCreateCmd        *cobra.Command
}

func (s *CreateCmdTestSuite) SetupTest() {
	s.mockOctaviusDClient = &daemon.MockClient{}
	s.testCreateCmd = NewCmd(s.mockOctaviusDClient)
}

func (s *CreateCmdTestSuite) TestCreateCmdUsage() {
	assert.Equal(s.T(), "create", s.testCreateCmd.Use)
}

func (s *CreateCmdTestSuite) TestCreateCmdHelp() {
	assert.Equal(s.T(), "Create new octavius job metadata", s.testCreateCmd.Short)
	assert.Equal(s.T(), "This command helps create new jobmetadata to your CP host with proper metadata.json file", s.testCreateCmd.Long)
	assert.Equal(s.T(), "octavius create PATH=<filepath>/metadata.json", s.testCreateCmd.Example)
}

func (s *CreateCmdTestSuite) TestCreateCmd() {

	s.mockOctaviusDClient.On("CreateMetadata","../../../../test/metadata/metadata.json").Return(nil).Once()

	args := []string{ "PATH=../../../../test/metadata/metadata.json"}
	s.testCreateCmd.Run(&cobra.Command{},args)

	s.mockOctaviusDClient.AssertExpectations(s.T())

}

func (s *CreateCmdTestSuite) TestCreateCmdWithWrongArguments() {

	s.mockOctaviusDClient.On("CreateMetadata","../../../../test/metadata/metadata.json").Return(nil).Once()

	args := []string{ "PAT=../../../../test/metadata/metadata.json"}
	s.testCreateCmd.Run(&cobra.Command{},args)

	s.mockOctaviusDClient.AssertNotCalled(s.T(),"CreateMetadata","../../../../test/metadata/metadata.json")

}

func (s *CreateCmdTestSuite) TestCreateCmdWithReturnedError() {

	s.mockOctaviusDClient.On("CreateMetadata","../../../../test/metadata/metadata.json").Return(errors.New("GOT SOME ERROR")).Once()

	args := []string{ "PATH=../../../../test/metadata/metadata.json"}
	s.testCreateCmd.Run(&cobra.Command{},args)

	s.mockOctaviusDClient.AssertExpectations(s.T())
}

func TestExecutionCmdTestSuite(t *testing.T) {
	suite.Run(t, new(CreateCmdTestSuite))
}