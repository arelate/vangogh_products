package vangogh_products

var supportsCopyFromTo = map[ProductType]ProductType{
	StoreProducts: WishlistProducts,
}

func SupportsCopy(pt1, pt2 ProductType) bool {
	return supportsCopyFromTo[pt1] == pt2
}
