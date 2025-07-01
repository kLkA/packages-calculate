package usecase

import (
	"context"
	"math"
	"sort"

	"homework/internal/pack"
	"homework/internal/pack/domain"
	"homework/internal/shared/logger"
)

// packService is an unexported struct used as the implementation of the pack.Service interface.
type packService struct {
}

func NewPackService() pack.Service {
	return &packService{}
}

// Calc computes the optimal packaging solution for a given total amount using specified pack sizes and returns the result.
func (s *packService) Calc(ctx context.Context, request domain.PacksCalcRequest) (*domain.CalculatePacksResponse, error) {
	ctx = logger.WithData(ctx, "svc", "pack")
	logger.Debugf(ctx, "calculating packs for total amount: %d with pack sizes: %v", request.TotalAmount, request.PackSizes)

	var packsMapResult map[int]int
	var totalAmount, totalPacks = math.MaxInt, math.MaxInt
	sort.Sort(sort.Reverse(sort.IntSlice(request.PackSizes)))
	var backtrack func(int, int, int, map[int]int)

	backtrack = func(index int, curAmount int, curPacks int, packsMap map[int]int) {
		if index == len(request.PackSizes) {
			return
		}

		size := request.PackSizes[index]
		maxCount := (request.TotalAmount - curAmount + size - 1) / size
		if maxCount < 0 {
			maxCount = 0
		}

		for count := maxCount; count >= 0; count-- {
			nextAmount := curAmount + count*size
			nextPacks := curPacks + count

			if nextAmount > totalAmount || (nextAmount == totalAmount && nextPacks >= totalPacks) {
				continue
			}

			nextPacksMap := make(map[int]int, len(packsMap))
			for k, v := range packsMap {
				nextPacksMap[k] = v
			}
			if count > 0 {
				nextPacksMap[size] += count
			}

			if nextAmount >= request.TotalAmount {
				if nextAmount < totalAmount || (nextAmount == totalAmount && nextPacks < totalPacks) {
					totalAmount = nextAmount
					totalPacks = nextPacks
					packsMapResult = nextPacksMap
				}
			} else {
				backtrack(index+1, nextAmount, nextPacks, nextPacksMap)
			}
		}
	}

	backtrack(0, 0, 0, map[int]int{})

	var packs []domain.CalculatePackItem
	for i := 0; i < len(request.PackSizes); i++ {
		if packsMapResult[request.PackSizes[i]] > 0 {
			packs = append(packs, domain.CalculatePackItem{
				Size:  request.PackSizes[i],
				Count: packsMapResult[request.PackSizes[i]],
			})
		}
	}
	res := &domain.CalculatePacksResponse{
		Packs: packs,
	}

	logger.Infof(ctx, "calculation complete - total amount: %d, packs: %v", totalAmount, packs)
	return res, nil
}
