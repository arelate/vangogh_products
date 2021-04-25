package vangogh_products

func IsPaged(pt ProductType) bool {
	return containsProductType(Paged(), pt)
}
