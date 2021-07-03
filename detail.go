package vangogh_products

func IsDetail(pt ProductType) bool {
	return containsProductType(Detail(), pt)
}
