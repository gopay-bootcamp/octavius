// Package metadata implements metadata related functions
package metadata

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"octavius/internal/controller/server/repository/metadata"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"
	"testing"
)

func init() {
	log.Init("info", "", true, 1)
}

func TestGetMetadata(t *testing.T) {
	metadataRepoMock := new(metadata.MetadataMock)

	testExec := NewMetadataExec(metadataRepoMock)
	testClientInfo := &protofiles.ClientInfo{
		ClientEmail: "test@gmail.com",
		AccessToken: "random",
	}
	testRequestForDescribe := &protofiles.RequestToDescribe{
		JobName:    "testJobName",
		ClientInfo: testClientInfo,
	}
	var testMetadata = &protofiles.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
	}
	metadataRepoMock.On("GetValue", testRequestForDescribe.JobName).Return(testMetadata, nil)
	resultMetadata, getMetadataErr := testExec.GetMetadata(context.Background(), testRequestForDescribe)
	assert.Equal(t, testMetadata, resultMetadata)
	assert.Nil(t, getMetadataErr)
	metadataRepoMock.AssertExpectations(t)

}

func TestGetJobList(t *testing.T) {
	metadataRepoMock := new(metadata.MetadataMock)

	testExec := NewMetadataExec(metadataRepoMock)

	var jobList []string
	jobList = append(jobList, "demo-image-name")
	jobList = append(jobList, "demo-image-name-1")

	testResponse := &protofiles.JobList{
		Jobs: jobList,
	}

	metadataRepoMock.On("GetAvailableJobList").Return(testResponse, nil)

	res, err := testExec.GetJobList(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, testResponse, res)
}

func TestGetJobListForGetAvailableJobListFunctionErr(t *testing.T) {
	metadataRepoMock := new(metadata.MetadataMock)

	testExec := NewMetadataExec(metadataRepoMock)

	metadataRepoMock.On("GetAvailableJobList").Return(&protofiles.JobList{}, errors.New("error in GetAvailableJobList function"))

	_, err := testExec.GetJobList(context.Background())
	assert.Equal(t, "error in GetAvailableJobList function", err.Error())
}

func TestSaveMetadata(t *testing.T) {
	metadataRepoMock := new(metadata.MetadataMock)

	testExec := NewMetadataExec(metadataRepoMock)
	testMetadata := protofiles.Metadata{
		Name:        "testJobName",
		Description: "This is a test image",
		ImageName:   "images/test-image",
	}
	testMetadataName := protofiles.MetadataName{
		Name: "testJobName",
	}
	metadataRepoMock.On("Save", "testJobName", &testMetadata).Return(&testMetadataName, nil).Once()
	res, err := testExec.SaveMetadata(context.Background(), &testMetadata)
	assert.Nil(t, err)
	assert.Equal(t, &testMetadataName, res)
	metadataRepoMock.AssertExpectations(t)
}

func TestReadAllMetadata(t *testing.T) {
	metadataRepoMock := new(metadata.MetadataMock)

	testExec := NewMetadataExec(metadataRepoMock)
	testMetadata1 := protofiles.Metadata{
		Name:        "testJobName1",
		Description: "This is a test image1",
		ImageName:   "images/test-image1",
	}
	testMetadata2 := protofiles.Metadata{
		Name:        "testJobName2",
		Description: "This is a test image2",
		ImageName:   "images/test-image2",
	}
	var metadataArray []*protofiles.Metadata
	metadataArray = append(metadataArray, &testMetadata1)
	metadataArray = append(metadataArray, &testMetadata2)
	testMetadataArray := protofiles.MetadataArray{
		Values: metadataArray,
	}

	metadataRepoMock.On("GetAll").Return(&testMetadataArray, nil).Once()
	res, err := testExec.ReadAllMetadata(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, &testMetadataArray, res)
	metadataRepoMock.AssertExpectations(t)
}
