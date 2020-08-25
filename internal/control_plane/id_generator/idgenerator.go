package id_generator

import (
	id "github.com/sony/sonyflake"
	"octavius/internal/control_plane/logger"
)

func NextID()(int, error) {
	var idGenerator = id.NewSonyflake(id.Settings{})
	uid, err := idGenerator.NextID()
	if err != nil {
		logger.Error(err, "unique id generation failed")
	}
	return int(uid), err
}


