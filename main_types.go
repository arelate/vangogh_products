package vangogh_products

var detailMainProductTypes = map[ProductType][]ProductType{
	Details: {LicenceProducts, AccountProducts},
	ApiProductsV1: {
		StoreProducts,
		AccountProducts,
		ApiProductsV2,
	},
	ApiProductsV2: {
		StoreProducts,
		AccountProducts,
		ApiProductsV2, // includes-games, is-included-in-games, requires-games, is-required-by-games
	},
}

func MainTypes(pt ProductType) []ProductType {
	return detailMainProductTypes[pt]
}
