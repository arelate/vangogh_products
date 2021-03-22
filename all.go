package vangogh_products

func AllPaged() []ProductType {
	return []ProductType{
		StorePage,
		AccountPage,
		WishlistPage}
}

func AllDetail() []ProductType {
	return []ProductType{
		Details,
		ApiProductsV1,
		ApiProductsV2,
	}
}

func AllLocal() []ProductType {
	return []ProductType{
		StoreProducts,
		AccountProducts,
		WishlistProducts,
		Details,
		ApiProductsV1,
		ApiProductsV2,
	}
}
