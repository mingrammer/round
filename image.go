package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func encode(fm string, f *os.File, m image.Image) error {
	switch fm {
	case "png":
		return png.Encode(f, m)
	case "jpg", "jpeg":
		return jpeg.Encode(f, m, nil)
	default:
		return errInvalidFormat
	}
}

func decode(f *os.File) (image.Image, string, error) {
	m, fm, err := image.Decode(f)
	switch fm {
	case "png", "jpg", "jpeg":
		return m, fm, err
	default:
		return nil, "", errInvalidFormat
	}
}

func yCbCrToRGBA(m image.Image) image.Image {
	b := m.Bounds()
	nm := image.NewRGBA(b)
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			nm.Set(x, y, m.At(x, y))
		}
	}
	return nm
}
