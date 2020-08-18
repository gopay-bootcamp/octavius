package execution

import (
	"context"
	"octavius/internal/control_plane/server/metadata/repository"
	"octavius/pkg/protobuf"
	"reflect"
	"testing"
)

func Test_execution_SaveMetadataToDb(t *testing.T) {
	metadataRepoMock := new(repository.MetadataMock)
	
	metadataVal := &protobuf.Metadata{
					Author:      "littlestar642",
					ImageName:   "demo image",
					Name:        "test data",
					Description: "sample test metadata",
				}
	metadataResp := &protobuf.MetadataName{
				Name: "test data",
				Err: &protobuf.Error{
					ErrorCode:    0,
					ErrorMessage: "no error",
				},
			}
	metadataRepoMock.On("Save","test data",metadataVal).Return(metadataResp,nil)
	type fields struct {
		metadata repository.MetadataRepository
		ctx      context.Context
		cancel   context.CancelFunc
	}
	type args struct {
		ctx      context.Context
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
				metadata: metadataRepoMock,
				ctx:      context.Background(),
			},
			args: args{
				ctx: context.Background(),
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
			e := &execution{
				metadata: tt.fields.metadata,
				ctx:      tt.fields.ctx,
				cancel:   tt.fields.cancel,
			}
			if got := e.SaveMetadataToDb(tt.args.ctx, tt.args.metadata); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execution.SaveMetadataToDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

