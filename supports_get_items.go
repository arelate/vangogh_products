package vangogh_products

var supportsGetItems = []ProductType{
	Details,
	ApiProductsV1,
	ApiProductsV2,
	Licences,
}

func SupportsGetItems(pt ProductType) bool {
	return containsProductType(supportsGetItems, pt)
}
