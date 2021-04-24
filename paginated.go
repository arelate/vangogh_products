package vangogh_products

func IsPaged(pt ProductType) bool {
	for _, ppt := range Paged() {
		if ppt == pt {
			return true
		}
	}
	return false
}
