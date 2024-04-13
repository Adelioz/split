package logging

import (
	"context"

	"go.uber.org/zap"
)

type loggerCtxKey struct{}

// WithLogger adds logger to context
func WithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, l)
}

func L(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(loggerCtxKey{}).(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return l
}

func S(ctx context.Context) *zap.SugaredLogger {
	return L(ctx).Sugar()
}
