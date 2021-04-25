package vangogh_products

var requireAuth = []ProductType{
	AccountPage,
	WishlistPage,
	Details,
	Licences,
}

func RequiresAuth(pt ProductType) bool {
	return containsProductType(requireAuth, pt)
}
