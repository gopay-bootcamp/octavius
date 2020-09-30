package util

import "context"

type ContextKey string

// String used to stringify key value
func (c ContextKey) String() string {
	return string(c)
}

var (
	// ContextKeyUUID is the key uuid
	ContextKeyUUID ContextKey
)

// GetUUIDFromContext gets the jobID value from the context.
func GetUUIDFromContext(ctx context.Context) (string, bool) {
	UUID, ok := ctx.Value(ContextKeyUUID).(string)
	return UUID, ok
}
