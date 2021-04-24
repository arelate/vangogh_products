package vangogh_products

var productTypeRequiresAuth = []ProductType{
	AccountPage,
	WishlistPage,
	Details,
	Licences,
}

func RequiresAuth(pt ProductType) bool {
	for _, ra := range productTypeRequiresAuth {
		if ra == pt {
			return true
		}
	}
	return false
}
