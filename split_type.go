package vangogh_products

var splitProductTypes = map[ProductType]ProductType{
	StorePage:    StoreProducts,
	AccountPage:  AccountProducts,
	WishlistPage: WishlistProducts,
	Licences:     LicenceProducts,
	OrderPage:    Order,
}

func SplitType(pt ProductType) ProductType {
	splitProductType, ok := splitProductTypes[pt]
	if ok {
		return splitProductType
	}

	return Unknown
}
