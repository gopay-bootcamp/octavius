package util

import "context"

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	ContextKeyUUID contextKey
)

// GetJobIDFromContext gets the jobID value from the context.
func GetUUIDFromContext(ctx context.Context) (string, bool) {
	UUID, ok := ctx.Value(ContextKeyUUID).(string)
	return UUID, ok
}
