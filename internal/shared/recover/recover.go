package recover

import (
	"context"
	"runtime/debug"

	"homework/internal/shared/logger"
)

func Recover() {
	err := recover()

	if err != nil {
		logger.Errorf(context.Background(), "%+v\nstacktrace from panic:\n%s", err, string(debug.Stack()))
	}
}
