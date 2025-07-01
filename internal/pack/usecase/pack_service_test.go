package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"homework/internal/pack/domain"
)

var (
	defaultPackSizes = []int{250, 500, 1000, 2000, 5000}
	twoPackSize      = []int{1, 250}
)

// TestPackService_Calc tests the Calc function of the PackService to validate the calculation of optimal pack solutions.
func TestPackService_Calc(t *testing.T) {
	tests := []struct {
		name          string
		order         int
		packSizes     []int
		expectedPacks *domain.CalculatePacksResponse
	}{
		{
			name:      "1 item",
			order:     1,
			packSizes: defaultPackSizes,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  250,
						Count: 1,
					},
				},
			},
		},
		{
			name:      "250 items",
			order:     250,
			packSizes: defaultPackSizes,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  250,
						Count: 1,
					},
				},
			},
		},
		{
			name:      "251 items",
			order:     251,
			packSizes: defaultPackSizes,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  500,
						Count: 1,
					},
				},
			},
		},
		{
			name:      "501 items",
			order:     501,
			packSizes: defaultPackSizes,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  500,
						Count: 1,
					},
					{
						Size:  250,
						Count: 1,
					},
				},
			},
		},
		{
			name:      "12001 items",
			order:     12001,
			packSizes: defaultPackSizes,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  5000,
						Count: 2,
					},
					{
						Size:  2000,
						Count: 1,
					},
					{
						Size:  250,
						Count: 1,
					},
				},
			},
		},
		{
			name:      "3751 items",
			order:     3751,
			packSizes: defaultPackSizes,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  2000,
						Count: 2,
					},
				},
			},
		},
		{
			name:      "3751 items (2)",
			order:     3751,
			packSizes: twoPackSize,
			expectedPacks: &domain.CalculatePacksResponse{
				Packs: []domain.CalculatePackItem{
					{
						Size:  250,
						Count: 15,
					},
					{
						Size:  1,
						Count: 1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewPackService()
			got, err := s.Calc(context.Background(), domain.PacksCalcRequest{
				PackSizes:   tt.packSizes,
				TotalAmount: tt.order,
			})

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedPacks, got)
		})
	}
}
