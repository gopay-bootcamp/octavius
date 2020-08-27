package errors

import (
	"fmt"
	"octavius/internal/pkg/constant"
)

type errorStruct struct {
	errorCode int
	err       error
}

//New intializes the error struct with error code and error
func New(errorCode int, err error) error {
	return &errorStruct{errorCode, err}
}
func (e *errorStruct) Error() string {
	return fmt.Sprintf("%s, code=%d, %v", e.err, e.errorCode, constant.ErrorCode[e.errorCode])
}
