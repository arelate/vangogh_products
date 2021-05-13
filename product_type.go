package vangogh_products

type ProductType int

const (
	Unknown ProductType = iota
	StorePage
	StoreProducts
	AccountPage
	AccountProducts
	WishlistPage
	WishlistProducts
	Details
	ApiProductsV1
	ApiProductsV2
	Licences
	LicenceProducts
	OrderPage
	Order
)

var productTypeStrings = map[ProductType]string{
	Unknown:          "unknown-product-type",
	StorePage:        "store-page",
	StoreProducts:    "store-products",
	AccountPage:      "account-page",
	AccountProducts:  "account-products",
	WishlistPage:     "wishlist-page",
	WishlistProducts: "wishlist-products",
	Details:          "details",
	ApiProductsV1:    "api-products-v1",
	ApiProductsV2:    "api-products-v2",
	Licences:         "licences",
	LicenceProducts:  "licence-products",
	OrderPage:        "order-page",
	Order:            "order",
}

func (pt ProductType) String() string {
	str, ok := productTypeStrings[pt]
	if ok {
		return str
	}

	return productTypeStrings[Unknown]
}

func Parse(productType string) ProductType {
	for pt, str := range productTypeStrings {
		if str == productType {
			return pt
		}
	}
	return Unknown
}

func Valid(pt ProductType) bool {
	_, ok := productTypeStrings[pt]
	return ok && pt != Unknown
}
