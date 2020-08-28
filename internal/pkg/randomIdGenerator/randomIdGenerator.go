package randomIdGenerator

import (
	"github.com/sony/sonyflake"
)

type RandomIdGenerator interface {
	Generate() (uint64, error)
}

type randomIdGenerator struct {
	sonyFlake *sonyflake.Sonyflake
}

func NewRandomIdGenerator() RandomIdGenerator {
	sonyFlake := sonyflake.NewSonyflake(sonyflake.Settings{})
	return &randomIdGenerator{
		sonyFlake: sonyFlake,
	}
}

func (r *randomIdGenerator) Generate() (uint64, error) {
	randomId, err := r.sonyFlake.NextID()
	if err != nil {
		return 0, err
	}
	return randomId, nil
}
