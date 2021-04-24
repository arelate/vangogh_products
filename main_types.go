package vangogh_products

var detailMainProductTypes = map[ProductType][]ProductType{
	AccountProducts: {LicenceProducts},
	Details:         {AccountProducts},
	ApiProductsV1:   {StoreProducts, AccountProducts},
	ApiProductsV2: {
		StoreProducts,
		AccountProducts,
		ApiProductsV2, // .Links.IncludesGames
	},
}

func MainTypesPriorityOrder() []ProductType {
	return []ProductType{
		LicenceProducts,
		AccountProducts,
		StoreProducts,
	}
}

func MainTypes(pt ProductType) []ProductType {
	return detailMainProductTypes[pt]
}
