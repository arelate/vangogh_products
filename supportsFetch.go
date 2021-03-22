package vangogh_products

var supportFetch = []ProductType{
	Details,
	ApiProductsV1,
	ApiProductsV2,
}

func SupportsFetch(pt ProductType) bool {
	for _, spt := range supportFetch {
		if spt == pt {
			return true
		}
	}
	return false
}
