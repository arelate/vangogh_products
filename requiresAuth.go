package vangogh_types

var productTypeRequiresAuth = []ProductType{
	AccountPage,
	WishlistPage,
	Details,
}

func ProductTypeRequiresAuth(pt ProductType) bool {
	for _, ra := range productTypeRequiresAuth {
		if ra == pt {
			return true
		}
	}
	return false
}

var imageTypeRequiresAuth []ImageType

func ImageTypeRequiresAuth(it ImageType) bool {
	for _, itra := range imageTypeRequiresAuth {
		if itra == it {
			return true
		}
	}
	return false
}
