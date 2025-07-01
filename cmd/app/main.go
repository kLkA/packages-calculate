package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"homework/internal/platform/app"
	"homework/internal/shared/config"
	"homework/internal/shared/logger"
	pkgRecover "homework/internal/shared/recover"
)

// main is the entry point for the application, initializing configuration, logger, and running the main application logic.
func main() {
	defer pkgRecover.Recover()

	// cancel for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configPath := flag.String("c", "config.toml", "Path to `toml` configuration file")
	flag.Parse()

	cfg, err := config.GetConfig(*configPath)
	if err != nil {
		logger.Fatalf(ctx, "couldn't run application: %s", err)
	}

	setupLogger(ctx, cfg)

	// wait until the processes stop
	wg := &sync.WaitGroup{}
	application, err := app.NewApplication(ctx, cfg)
	if err != nil {
		logger.Fatalf(ctx, "couldn't run application: %s", err)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := application.Run(ctx); err != nil && err != http.ErrServerClosed {
			logger.Fatalf(ctx, "service run error: %s", err)
			return
		}
	}()

	<-listenShutdownSignal()
	logger.Infof(ctx, "got signal starting shutdown")

	if err := application.StopAPI(ctx); err != nil {
		logger.Errorf(ctx, "api stop error: %s", err.Error())
	}

	cancel()
	wg.Wait()

	logger.Infof(ctx, "application is stopped")
}

// listenShutdownSignal listens for OS interrupt or termination signals and returns a channel to handle graceful shutdown.
func listenShutdownSignal() chan os.Signal {
	quit := make(chan os.Signal, 1)
	// interrupt signal sent from the terminal
	signal.Notify(quit, os.Interrupt)
	// sigterm signal sent from k8s
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	return quit
}

// setupLogger initializes the application's logging configuration and sets the global log level based on the provided config.
func setupLogger(ctx context.Context, cfg *config.Config) {
	logger.InitWithConfig(ctx, cfg.Env)

	const defaultLogLevel = "info"
	logLevel := cfg.LogLevel
	if logLevel == "" {
		logLevel = defaultLogLevel
	}
	zeroLogLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Printf("couldn't parse log level: %s\n", err)
		zeroLogLevel, _ = zerolog.ParseLevel(defaultLogLevel)
	}
	logger.SetGlobalLevel(zeroLogLevel)
}
