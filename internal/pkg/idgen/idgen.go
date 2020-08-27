package idgen

// package naming is wrong.

import (
	"github.com/sony/sonyflake"
)

var flake *sonyflake.Sonyflake

func init() {
	if flake == nil {
		flake = sonyflake.NewSonyflake(sonyflake.Settings{})
	}
}

// NextID used for generating next random id (for associated with request number)
func NextID() (uint64, error) {
	uid, err := flake.NextID()
	return uid, err
}
