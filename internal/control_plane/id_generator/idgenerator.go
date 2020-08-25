package id_generator

import (
	"github.com/sony/sonyflake"
)

//used for generating next random id (for associated with request number)
func NextID() (uint64, error) {
	var idGenerator = sonyflake.NewSonyflake(sonyflake.Settings{})
	uid, err := idGenerator.NextID()
	return uid, err
}
