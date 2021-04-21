package vangogh_products

import "github.com/arelate/vangogh_images"

var supportedImageTypes = map[ProductType][]vangogh_images.ImageType{
	StoreProducts: {
		vangogh_images.Image,
		vangogh_images.Screenshots,
	},
	AccountProducts:  {vangogh_images.Image},
	WishlistProducts: {vangogh_images.Image},
	Details:          {vangogh_images.Background},
	ApiProductsV1: {
		vangogh_images.Icon,
		vangogh_images.Background,
		vangogh_images.Screenshots,
	},
	ApiProductsV2: {
		vangogh_images.Image,
		vangogh_images.BoxArt,
		vangogh_images.Logo,
		vangogh_images.Icon,
		vangogh_images.Background,
		vangogh_images.GalaxyBackground,
		vangogh_images.Screenshots,
	},
}

func AllSupportingImageType(imageType vangogh_images.ImageType) []ProductType {
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

func SupportsImageType(pt ProductType, it vangogh_images.ImageType) bool {
	if !Valid(pt) ||
		!vangogh_images.Valid(it) {
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
