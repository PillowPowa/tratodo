package context

import (
	"context"
	"fmt"
)

type contextKey string

const contextUserIdKey contextKey = "userID"

func NewAuthContext(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, contextUserIdKey, id)
}

func InferAuthContext(ctx context.Context) (int64, error) {
	userId, ok := ctx.Value(contextUserIdKey).(int64)
	if !ok {
		return 0, fmt.Errorf(`
			context.InferAuthContext:
			cannot infer user id from context (make sure you use NewAuthContext to create context with user id)
		`)
	}
	return userId, nil
}
