package randomIdGenerator

import "github.com/stretchr/testify/mock"

type IdGeneratorMock struct {
	mock.Mock
}

func (m *IdGeneratorMock) Generate() (uint64,error){
	args:= m.Called()
	return args.Get(0).(uint64), args.Error(1)
}
