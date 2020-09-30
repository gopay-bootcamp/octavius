package idgen

import (
	"github.com/sony/sonyflake"
)

// RandomIdGenerator interface
type RandomIdGenerator interface {
	Generate() (uint64, error)
}

type randomIdGenerator struct {
	sonyFlake *sonyflake.Sonyflake
}

// NewRandomIdGenerator used to return RandomIdGenerator interface
func NewRandomIdGenerator() RandomIdGenerator {
	sonyFlake := sonyflake.NewSonyflake(sonyflake.Settings{})
	return &randomIdGenerator{
		sonyFlake: sonyFlake,
	}
}

// Generate used to generate random ID
func (r *randomIdGenerator) Generate() (uint64, error) {
	randomId, err := r.sonyFlake.NextID()
	if err != nil {
		return 0, err
	}
	return randomId, nil
}
