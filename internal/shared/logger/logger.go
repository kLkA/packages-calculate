package logger

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type correlationIdType int

const (
	traceKey correlationIdType = iota

	graylogTimestampFormat = "2006-01-02T15:04:05.999Z"

	LogModeDev  = "dev"
	LogModeProd = "prod"
)

var gLogger zerolog.Logger

func InitWithConfig(ctx context.Context, env string) {
	zerolog.TimestampFieldName = "log_time"
	zerolog.TimeFieldFormat = graylogTimestampFormat
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true

	var logWriter io.Writer
	switch env {
	case LogModeDev:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		consoleWriter.TimeFormat = time.StampMilli
		logWriter = io.Writer(consoleWriter)
		log.Println("logging to user pretty console output in dev mode")
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logWriter = io.Writer(os.Stdout)
		log.Println("logging to os.Stdout in production Mode")
	}

	gLogger = zerolog.New(logWriter).With().Timestamp().Logger()

	switch env {
	case LogModeDev:
		Info(ctx, "logging to user pretty console output in dev mode")
	default:
		Info(ctx, "logging to os.Stdout in production Mode")
	}

}

func Info(ctx context.Context, msg string) {
	logWithTrace(ctx, gLogger.Info()).Msg(msg)
}

func Infof(ctx context.Context, msg string, v ...interface{}) {
	logWithTrace(ctx, gLogger.Info()).Msgf(msg, v...)
}

func Warn(ctx context.Context, msg string) {
	logWithTrace(ctx, gLogger.Warn()).Msg(msg)
}

func Warnf(ctx context.Context, msg string, v ...interface{}) {
	logWithTrace(ctx, gLogger.Warn()).Msgf(msg, v...)
}

func Error(ctx context.Context, msg string) {
	logWithTrace(ctx, gLogger.Error()).Msg(msg)
}

func Errorf(ctx context.Context, msg string, v ...interface{}) {
	logWithTrace(ctx, gLogger.Error()).Msgf(msg, v...)
}

func Fatalf(ctx context.Context, msg string, v ...interface{}) {
	logWithTrace(ctx, gLogger.Fatal()).Msgf(msg, v...)
}

func Fatal(ctx context.Context, msg string) {
	logWithTrace(ctx, gLogger.Fatal()).Msg(msg)
}

func Debug(ctx context.Context, msg string) {
	logWithTrace(ctx, gLogger.Debug()).Msg(msg)
}

func Debugf(ctx context.Context, msg string, v ...interface{}) {
	logWithTrace(ctx, gLogger.Debug()).Msgf(msg, v...)
}
func SetGlobalLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
	gLogger.Level(level)
}
