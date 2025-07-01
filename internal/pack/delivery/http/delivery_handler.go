package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"homework/internal/pack"
	"homework/internal/pack/delivery/mapper"
	"homework/internal/shared/logger"
)

// NewServer initializes and returns a new PackServer with the provided gin engine and pack service.
func NewServer(r *gin.Engine, service pack.Service) *PackServer {
	srv := &PackServer{
		service: service,
	}
	rg := r.Group("/v1/pack")
	rg.POST("/calc", srv.Calc)
	return srv
}

// PackServer is a HTTP server handler that manages pack-related API endpoints using a provided pack.Service.
// It serves as the interface for handling HTTP requests, such as pack calculation operations.
type PackServer struct {
	service pack.Service
}

// Calc handles the HTTP POST request to calculate packaging distribution based on the provided request payload.
func (s PackServer) Calc(ctx *gin.Context) {
	// Bind JSON request body to the PacksCalcRequest struct
	var in mapper.PacksCalcRequest
	if err := ctx.ShouldBindJSON(&in); err != nil {
		// If the input JSON is invalid, return an error response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input JSON"})
		return
	}

	request := mapper.ToDomainPackCalculateRequest(in)
	response, err := s.service.Calc(ctx.Request.Context(), request)

	if err != nil {
		logger.Errorf(ctx, "error to get pack by id: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if response == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	out := mapper.ToHttpPackCalculate(response)
	ctx.JSON(http.StatusOK, out)
}
