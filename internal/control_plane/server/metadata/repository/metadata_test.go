package repository

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"octavius/internal/control_plane/db/etcd"
	"octavius/pkg/protobuf"
	"reflect"
	"testing"
)

func Test_metadataRepository_Save(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	metadataVal := &protobuf.Metadata{
		Author:      "littlestar642",
		ImageName:   "demo image",
		Name:        "test data",
		Description: "sample test metadata",
	}
	val, err := proto.Marshal(metadataVal)
	if err != nil {
		t.Error("error in marshalling metadata")
	}
	mockClient.On("PutValue", "metadata/test data", string(val)).Return("metadata/test data", nil)

	type fields struct {
		etcdClient etcd.EtcdClient
	}
	type args struct {
		ctx      context.Context
		key      string
		metadata *protobuf.Metadata
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *protobuf.MetadataName
	}{
		{
			fields: fields{
				etcdClient: mockClient,
			},
			args: args{
				ctx:      context.Background(),
				key:      "test data",
				metadata: metadataVal,
			},
			want: &protobuf.MetadataName{
				Name: "test data",
				Err: &protobuf.Error{
					ErrorCode:    0,
					ErrorMessage: "no error",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &metadataRepository{
				etcdClient: tt.fields.etcdClient,
			}
			if got := c.Save(tt.args.ctx, tt.args.key, tt.args.metadata); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("metadataRepository.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
