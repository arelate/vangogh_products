package vangogh_types

import (
	"github.com/arelate/gog_media"
)

// unsupported is used instead of supported in similar cases to
// avoid all, but one repetitive data
var unsupportedMedia = map[ProductType][]gog_media.Media{
	ApiProductsV2: {gog_media.Movie},
}

func SupportsMedia(pt ProductType, mt gog_media.Media) bool {
	if !gog_media.Valid(mt) {
		return false
	}
	if !ValidProductType(pt) {
		return false
	}

	ums, ok := unsupportedMedia[pt]
	if !ok {
		return true
	}

	for _, um := range ums {
		if um == mt {
			return false
		}
	}

	return true
}
