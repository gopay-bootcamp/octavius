package metadata

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"octavius/internal/pkg/constant"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"testing"

	"github.com/golang/protobuf/proto"
)

func init() {
	log.Init("info", "", false)
}

func Test_metadataRepository_Save(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	metadataVal := &clientCPproto.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val, err := proto.Marshal(metadataVal)
	if err != nil {
		t.Error("error in marshalling metadata")
	}
	mockClient.On("PutValue", "metadata/test data", string(val)).Return(nil)
	mockClient.On("GetValue", "metadata/test data").Return("", nil)

	testMetadataRepo := NewMetadataRepository(mockClient)
	ctx := context.Background()
	sr, err := testMetadataRepo.Save(ctx, "test data", metadataVal)

	if err != nil {
		t.Error(err, "saving metadata failed")
	}
	if sr.Name != "test data" {
		t.Errorf("expected %s, got %s", "test data", sr.Name)
	}

	mockClient.AssertExpectations(t)
}

func Test_metadataRepository_Save_KeyAlreadyPresent(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	metadataVal := &clientCPproto.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val, err := proto.Marshal(metadataVal)
	if err != nil {
		t.Error("error in marshalling metadata")
	}
	mockClient.On("PutValue", "metadata/test data", string(val)).Return(nil)
	mockClient.On("GetValue", "metadata/test data").Return("some key", nil)

	testMetadataRepo := NewMetadataRepository(mockClient)
	ctx := context.Background()
	_, err = testMetadataRepo.Save(ctx, "test data", metadataVal)

	if err.Error() != status.Error(codes.AlreadyExists, errors.New(constant.Etcd+constant.KeyAlreadyPresent).Error()).Error() {
		t.Error("key already present error expected")
	}
}

func Test_metadataRepository_Save_GetValueError(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	metadataVal := &clientCPproto.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val, err := proto.Marshal(metadataVal)
	if err != nil {
		t.Error("error in marshalling metadata")
	}
	mockClient.On("PutValue", "metadata/test data", string(val)).Return(nil)
	mockClient.On("GetValue", "metadata/test data").Return("", errors.New("some error"))

	testMetadataRepo := NewMetadataRepository(mockClient)
	ctx := context.Background()
	_, err = testMetadataRepo.Save(ctx, "test data", metadataVal)

	if err.Error() != status.Error(codes.Internal, "some error").Error() {
		t.Error("get value error expected")
	}
}

func Test_metadataRepository_Save_PutValueError(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	metadataVal := &clientCPproto.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val, err := proto.Marshal(metadataVal)
	if err != nil {
		t.Error("error in marshalling metadata")
	}
	mockClient.On("PutValue", "metadata/test data", string(val)).Return(errors.New("some error"))
	mockClient.On("GetValue", "metadata/test data").Return("", nil)

	testMetadataRepo := NewMetadataRepository(mockClient)
	ctx := context.Background()
	_, err = testMetadataRepo.Save(ctx, "test data", metadataVal)

	if err.Error() != status.Error(codes.Internal, "some error").Error() {
		t.Error("put value error expected")
	}
}

func Test_metadataRepository_GetAll(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	metadataArr := make([]string, 3)

	metadataVal1 := &clientCPproto.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val1, err := proto.Marshal(metadataVal1)
	metadataArr = append(metadataArr, string(val1))

	metadataVal2 := &clientCPproto.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val2, err := proto.Marshal(metadataVal2)
	metadataArr = append(metadataArr, string(val2))

	mockClient.On("GetAllValues").Return(metadataArr, nil)

	testMetadataRepo := NewMetadataRepository(mockClient)
	ctx := context.Background()
	_, err = testMetadataRepo.GetAll(ctx)

	if err != nil {
		t.Error(err, "saving metadata failed")
	}

	mockClient.AssertExpectations(t)
}
