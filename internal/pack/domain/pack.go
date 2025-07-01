package domain

type CalculatePackItem struct {
	Size  int
	Count int
}

type CalculatePacksResponse struct {
	Packs []CalculatePackItem
}

type PacksCalcRequest struct {
	TotalAmount int
	PackSizes   []int
}
