package vangogh_products

var detailMainProductTypes = map[ProductType][]ProductType{
	Details:       {AccountProducts},
	ApiProductsV1: {StoreProducts, AccountProducts},
	ApiProductsV2: {StoreProducts, AccountProducts},
}

func MainTypes(pt ProductType) []ProductType {
	return detailMainProductTypes[pt]
}
