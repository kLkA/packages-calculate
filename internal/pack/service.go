package pack

import (
	"context"

	"homework/internal/pack/domain"
)

type Service interface {
	Calc(context.Context, domain.PacksCalcRequest) (*domain.CalculatePacksResponse, error)
}
