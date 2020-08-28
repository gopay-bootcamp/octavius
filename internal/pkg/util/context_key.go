package util

import "context"

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var (
	ContextKeyUUID ContextKey
)

// GetJobIDFromContext gets the jobID value from the context.
func GetUUIDFromContext(ctx context.Context) (string, bool) {
	UUID, ok := ctx.Value(ContextKeyUUID).(string)
	return UUID, ok
}
