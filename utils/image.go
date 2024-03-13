package utils

import (
	"image/jpeg"
	"image/png"
	"io"
)

func CheckIsRealImage(file io.Reader) bool {
	_, err := png.Decode(file)
	if err == nil {
		return true
	}
	_, err = jpeg.Decode(file)
	if err == nil {
		return true
	}
	return false
}
