package vangogh_products

var supportsCopyFromTo = map[ProductType]ProductType{
	StoreProducts: WishlistProducts,
}

func SupportsCopy(from, to ProductType) bool {
	return supportsCopyFromTo[from] == to
}
