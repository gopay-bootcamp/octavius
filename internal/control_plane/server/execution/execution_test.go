package execution

import (
	"context"
	"octavius/internal/control_plane/server/metadata/repository"
	"octavius/pkg/protobuf"
	"reflect"
	"testing"
)


func Test_execution_SaveMetadataToDb(t *testing.T) {
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
		want   *protobuf.MetadataID
	}{
		// TODO: Add test cases.
		
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
