package octaviusErrors

import (
	"fmt"
	"octavius/internal/pkg/constant"
)

type errorStruct struct {
	errorCode int
	err       error
}

func New(errorCode int, err error) error {
	return &errorStruct{errorCode, err}
}
func (e *errorStruct) Error() string {
	return fmt.Sprintf("%s{error code=%d(%s)}", e.err, e.errorCode, constant.ErrorCode[e.errorCode])
}
