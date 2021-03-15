package vangogh_types

type ImageType int

const (
	UnknownImageType ImageType = iota
	Image
	BoxArt
	Logo
	Icon
	Background
	GalaxyBackground
	Screenshots
)

var imageTypeStrings = map[ImageType]string{
	UnknownImageType: "unknown-image-type",
	Image:            "image",
	BoxArt:           "box-art",
	Logo:             "logo",
	Icon:             "icon",
	Background:       "background",
	GalaxyBackground: "galaxy-background",
	Screenshots:      "screenshots",
}

func (it ImageType) String() string {
	str, ok := imageTypeStrings[it]
	if ok {
		return str
	}

	return imageTypeStrings[UnknownImageType]
}

func ParseImageType(imageType string) ImageType {
	for it, str := range imageTypeStrings {
		if str == imageType {
			return it
		}
	}
	return UnknownImageType
}

func ValidImageType(it ImageType) bool {
	_, ok := imageTypeStrings[it]
	return ok && it != UnknownImageType
}
