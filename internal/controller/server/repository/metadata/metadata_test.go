package metadata

import (
	"context"
	"octavius/internal/pkg/db/etcd"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
)

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
	type fields struct {
		etcdClient etcd.Client
	}
	type args struct {
		ctx      context.Context
		key      string
		metadata *clientCPproto.Metadata
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *clientCPproto.MetadataName
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				etcdClient: mockClient,
			},
			args: args{
				ctx:      context.Background(),
				key:      "test data",
				metadata: metadataVal,
			},
			want: &clientCPproto.MetadataName{
				Name: "test data",
				Err: &clientCPproto.Error{
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
			got, err := c.Save(tt.args.ctx, tt.args.key, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("metadataRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("metadataRepository.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
