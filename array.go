package vangogh_products

func IsArray(pt ProductType) bool {
	return containsProductType(Array(), pt)
}
