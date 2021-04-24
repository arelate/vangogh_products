package vangogh_products

var supportsGetItems = []ProductType{
	Details,
	ApiProductsV1,
	ApiProductsV2,
}

func SupportsGetItems(pt ProductType) bool {
	for _, spt := range supportsGetItems {
		if spt == pt {
			return true
		}
	}
	return false
}
