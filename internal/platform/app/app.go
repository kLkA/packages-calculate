package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	packDeliveryHttp "homework/internal/pack/delivery/http"
	"homework/internal/shared/config"
	"homework/internal/shared/logger"
)

type Application struct {
	cfg      *config.Config
	registry *Registry

	httpServer *http.Server
	packServer *packDeliveryHttp.PackServer
}

// NewApplication initializes a new Application instance with the provided context and configuration.
// It sets up the required components such as the registry, HTTP server, and pack server.
// It returns an Application pointer and an error if any issues occur during initialization.
func NewApplication(ctx context.Context, c *config.Config) (*Application, error) {
	registry := NewRegistry(context.Background(), c)

	// err is skipped as it will be available in registry.GetErrors()
	packServer, _ := registry.GetPackServer()

	app := Application{
		registry: registry,

		cfg:        c,
		httpServer: registry.GetHttpServer(),
		packServer: packServer,
	}
	errorsList := registry.GetErrors()
	if len(errorsList) > 0 {
		for _, err := range errorsList {
			logger.Fatal(ctx, err.Error())
		}
		return nil, errorsList[len(errorsList)-1]
	}
	return &app, nil
}

// Run starts the HTTP server and begins listening for incoming requests on the specified port in the configuration.
func (app *Application) Run(ctx context.Context) error {
	// Start to listen to events
	logger.Info(ctx, fmt.Sprintf("app running to listen http handler on %s", app.cfg.Http.Port))
	if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(ctx, fmt.Sprintf("fail to start http server: %s", err))
	}

	return nil
}

// StopAPI gracefully stops the HTTP server by shutting it down within a 10-second timeout. Returns an error if shutdown fails.
func (app *Application) StopAPI(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("stopping listen http server on %s", app.cfg.Http.Port))

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := app.httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error(ctx, fmt.Sprintf("fail to gracefully shutdown http server: %s", err))
	}

	return nil
}
