package octaviusErrors

import (
	"fmt"
	"octavius/pkg/constant"
)

type errorStruct struct {
	errorCode int
	err       error
}

func New(errorCode int, err error) error {
	return &errorStruct{errorCode, err}
}
func (e *errorStruct) Error() string {
	return fmt.Sprintf("Error code=%d Becasue of %s error=%s", e.errorCode, constant.ErrorCode[e.errorCode], e.err)
}