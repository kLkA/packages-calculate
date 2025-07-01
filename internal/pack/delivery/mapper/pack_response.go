package mapper

import (
	"homework/internal/pack/domain"
)

type CalculatePackItem struct {
	Size  int `json:"size,omitempty"`
	Count int `json:"count,omitempty"`
}

type CalculatePacksResponse struct {
	Packs []CalculatePackItem `json:"packs,omitempty"`
}

// ToHttpPackCalculate transforms a domain.CalculatePacksResponse object into an HTTP response representation.
// Converts each pack item within the input into a CalculatePackItem with size and count fields.
// Returns nil if the input parameter resp is nil.
func ToHttpPackCalculate(resp *domain.CalculatePacksResponse) *CalculatePacksResponse {
	if resp == nil {
		return nil
	}
	var packs []CalculatePackItem
	for _, p := range resp.Packs {
		packs = append(packs, CalculatePackItem{
			Size:  p.Size,
			Count: p.Count,
		})
	}
	out := &CalculatePacksResponse{
		Packs: packs,
	}
	return out
}
