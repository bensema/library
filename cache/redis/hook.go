package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

type Hook interface {
	BeforeProcess(ctx context.Context, cmd Cmder) (context.Context, error)
	AfterProcess(ctx context.Context, cmd Cmder) error
}

type hooks struct {
	hooks []Hook
}

func (hs *hooks) AddHook(hook Hook) {
	hs.hooks = append(hs.hooks, hook)
}

func (hs *hooks) beforeProcess(ctx context.Context, cmd Cmder) (context.Context, error) {
	for _, h := range hs.hooks {
		var err error
		ctx, err = h.BeforeProcess(ctx, cmd)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

func (hs hooks) afterProcess(ctx context.Context, cmd Cmder) error {
	var firstErr error
	for _, h := range hs.hooks {
		err := h.AfterProcess(ctx, cmd)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

type OpenTelemetryHook struct{}

func (OpenTelemetryHook) BeforeProcess(ctx context.Context, cmd Cmder) (context.Context, error) {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx, nil
	}

	tracer := otel.Tracer("github.com/bensema/library/redis")
	ctx, span := tracer.Start(ctx, cmd.Name())
	span.SetAttributes(
		label.String("db.system", "redis"),
		label.String("redis.cmd", fmt.Sprintf("%s %s", cmd.Name(), cmd.Args())),
	)

	return ctx, nil
}

func (OpenTelemetryHook) AfterProcess(ctx context.Context, cmd Cmder) error {
	span := trace.SpanFromContext(ctx)
	if err := cmd.Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.End()
	return nil
}

func recordError(ctx context.Context, span trace.Span, err error) {
	if err != redis.Nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
