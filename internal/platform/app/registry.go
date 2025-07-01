package app

import (
	"context"
	"net/http"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"homework/internal/pack"
	_packHTTP "homework/internal/pack/delivery/http"
	"homework/internal/shared/config"
	"homework/internal/shared/middleware"
)

// Registry is a struct responsible for managing application dependencies and configurations.
// It encapsulates context, configuration, HTTP server, router, middleware, and service instances.
// This type also collects errors that may occur during initialization or runtime operations.
type Registry struct {
	ctx context.Context
	cfg *config.Config

	httpServer        *http.Server
	router            *gin.Engine
	tracingMiddleware gin.HandlerFunc

	packServer  *_packHTTP.PackServer
	packService pack.Service

	errors []error
}

func NewRegistry(ctx context.Context, cfg *config.Config) *Registry {
	return &Registry{
		ctx: ctx,
		cfg: cfg,
	}
}

func (r *Registry) GetHttpServer() *http.Server {
	if r.httpServer == nil {
		r.httpServer = &http.Server{
			Addr:    r.cfg.Http.Port,
			Handler: r.GetRouter(),
		}
	}
	return r.httpServer
}

func (r *Registry) GetRouter() *gin.Engine {
	if r.router == nil {
		if r.cfg.Env == "prod" {
			gin.SetMode(gin.ReleaseMode)
		}
		r.router = gin.New()
		r.router.Use(r.GetTracingMiddleware())
		// simple router logging, have to be adapted to fit the same structure of main logs or disabled when not needed
		r.router.Use(logger.SetLogger())
	}
	return r.router
}

func (r *Registry) GetErrors() []error {
	return r.errors
}

func (r *Registry) GetTracingMiddleware() gin.HandlerFunc {
	if r.tracingMiddleware == nil {
		r.tracingMiddleware = middleware.NewTracingMiddleware()
	}
	return r.tracingMiddleware
}

func (r *Registry) GetDefaultMuxHandler(mux *runtime.ServeMux) gin.HandlerFunc {
	return func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
	}
}
