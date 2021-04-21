package vangogh_products

func Paged() []ProductType {
	return []ProductType{
		StorePage,
		AccountPage,
		WishlistPage}
}

func Detail() []ProductType {
	return []ProductType{
		Details,
		ApiProductsV1,
		ApiProductsV2,
	}
}

func Remote() []ProductType {
	remote := make([]ProductType, 0)
	remote = append(remote, Paged()...)
	return append(remote, Detail()...)
}

func Local() []ProductType {
	return []ProductType{
		StoreProducts,
		AccountProducts,
		WishlistProducts,
		Details,
		ApiProductsV1,
		ApiProductsV2,
	}
}
