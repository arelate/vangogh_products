package vangogh_products

var detailMainProductTypes = map[ProductType][]ProductType{
	Details:       {AccountProducts},
	ApiProductsV1: {StoreProducts, AccountProducts},
	ApiProductsV2: {
		StoreProducts,
		AccountProducts,
		ApiProductsV2, // .Links.IncludesGames
	},
}

func MainTypes(pt ProductType) []ProductType {
	return detailMainProductTypes[pt]
}
