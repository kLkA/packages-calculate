package mapper

import (
	"homework/internal/pack/domain"
)

// PacksCalcRequest represents the input payload for calculating pack distribution based on the total amount and pack sizes.
type PacksCalcRequest struct {
	TotalAmount int   `json:"total_amount,omitempty"`
	PackSizes   []int `json:"pack_sizes,omitempty"`
}

// ToDomainPackCalculateRequest converts a PacksCalcRequest object into a domain.PacksCalcRequest.
func ToDomainPackCalculateRequest(in PacksCalcRequest) domain.PacksCalcRequest {
	request := domain.PacksCalcRequest{
		PackSizes:   in.PackSizes,
		TotalAmount: in.TotalAmount,
	}
	return request
}
