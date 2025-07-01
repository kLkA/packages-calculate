package logger

import (
	"context"
	"sync"

	"github.com/rs/zerolog"
)

type Trace struct {
	mux *sync.RWMutex

	Data map[string]string
}

func logWithTrace(ctx context.Context, e *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return e
	}

	if trace, ok := GetTraceFromContext(ctx); ok {
		trace.mux.RLock()
		if trace.Data != nil {
			for name, value := range trace.Data {
				e.Str(name, value)
			}
		}
		trace.mux.RUnlock()
	}

	return e
}

func WithData(ctx context.Context, name string, value string) (newCtx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	trace := GetOrCreateTraceFromContext(ctx)

	trace.mux.Lock()
	defer trace.mux.Unlock()
	newTrace := trace.Copy()
	newTrace.Data[name] = value

	return context.WithValue(ctx, traceKey, &newTrace)
}

func GetOrCreateTraceFromContext(ctx context.Context) Trace {
	u, ok := GetTraceFromContext(ctx)
	if !ok || u == nil {
		u = &Trace{
			Data: map[string]string{},
			mux:  &sync.RWMutex{},
		}
	}
	return *u
}

func GetTraceFromContext(ctx context.Context) (*Trace, bool) {
	u, ok := ctx.Value(traceKey).(*Trace)
	return u, ok
}

func (trace Trace) Copy() Trace {
	newTrace := trace
	newTrace.mux = &sync.RWMutex{}
	newTrace.Data = map[string]string{}

	for key, value := range trace.Data {
		newTrace.Data[key] = value
	}
	return newTrace
}
