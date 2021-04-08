package vangogh_products

var paginatedProductTypes = []ProductType{
	StorePage,
	AccountPage,
	WishlistPage,
}

func Paginated(pt ProductType) bool {
	for _, ppt := range paginatedProductTypes {
		if ppt == pt {
			return true
		}
	}
	return false
}
