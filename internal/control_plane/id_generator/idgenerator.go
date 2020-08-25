package id_generator

import (
	"github.com/sony/sonyflake"
	"octavius/internal/control_plane/logger"
)

//used for generating next random id (for associated with request number)
func NextID() uint64 {
	var idGenerator = sonyflake.NewSonyflake(sonyflake.Settings{})
	uid, err := idGenerator.NextID()
	if err != nil {
		logger.Error(err, "unique id generation failed")
	}
	return uid
}
