package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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

func encodeBase64(fm string, m image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := errInvalidFormat
	var prefix string
	switch fm {
	case "png":
		prefix = "data:image/png;base64,"
		err = png.Encode(buf, m)
	case "jpg", "jpeg":
		prefix = "data:image/jpeg;base64,"
		err = jpeg.Encode(buf, m, nil)
	}
	if err != nil {
		return "", err
	}
	return prefix + base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func decode(r io.Reader) (image.Image, string, error) {
	m, fm, err := image.Decode(r)
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
