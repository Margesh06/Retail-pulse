package utils

import (
	"image"
	"io"
	"log"

	_ "image/jpeg"
	_ "image/png"
)

func GetImageDimensions(reader io.Reader) (int, int) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Printf("Failed to decode image: %v\n", err)
		return 0, 0
	}
	return img.Width, img.Height
}
