package vangogh_types

var supportedImageTypes = map[ProductType][]ImageType{
	StoreProducts:    {Image, Screenshots},
	AccountProducts:  {Image},
	WishlistProducts: {Image},
	Details:          {Background},
	ApiProductsV1:    {Icon, Background, Screenshots},
	ApiProductsV2:    {Image, BoxArt, Logo, Icon, Background, GalaxyBackground, Screenshots},
}

func ProductTypesSupportingImageType(imageType ImageType) []ProductType {
	pts := make([]ProductType, 0)
	for pt, its := range supportedImageTypes {
		for _, it := range its {
			if it == imageType {
				pts = append(pts, pt)
				break
			}
		}
	}
	return pts
}

func SupportsImageType(pt ProductType, it ImageType) bool {
	if !ValidProductType(pt) ||
		!ValidImageType(it) {
		return false
	}

	supportedIts, ok := supportedImageTypes[pt]
	if !ok {
		return false
	}

	for _, sit := range supportedIts {
		if sit == it {
			return true
		}
	}

	return false
}
