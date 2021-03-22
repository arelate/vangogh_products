module github.com/arelate/vangogh_products

go 1.16

require (
	github.com/arelate/gog_media v0.1.2-alpha
	github.com/arelate/vangogh_images v0.1.0-alpha
)

replace (
	github.com/arelate/gog_media => ../gog_media
	github.com/arelate/vangogh_images => ../vangogh_images
)
