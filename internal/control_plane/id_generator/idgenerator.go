package id_generator

import (
	id "github.com/sony/sonyflake"
	"octavius/internal/control_plane/logger"
)

func NextID() uint64 {
	var idGenerator = id.NewSonyflake(id.Settings{})
	uid, err := idGenerator.NextID()
	if err != nil {
		logger.Error(err, "unique id generation failed")
	}
	return uid
}
