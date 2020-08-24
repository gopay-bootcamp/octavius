package execution

import (
	"context"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/metadata/repository"
	protobuf "octavius/pkg/protofiles/client_CP"
	"reflect"
	"testing"
)

func init() {
	logger.Setup("info")
}

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
	metadataRepoMock.On("Save", "test data", metadataVal).Return(metadataResp, nil)
	type fields struct {
		metadata repository.MetadataRepository
	}
	type args struct {
		ctx      context.Context
		metadata *protobuf.Metadata
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *protobuf.MetadataName
		wantErr bool
	}{
		{
			fields: fields{
				metadata: metadataRepoMock,
			},
			args: args{
				ctx:      context.Background(),
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
			}
			got, err := e.SaveMetadataToDb(tt.args.ctx, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("execution.SaveMetadataToDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execution.SaveMetadataToDb() = %v, want %v", got, tt.want)
			}
		})
	}
}
