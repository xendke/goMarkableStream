package client

import (
	"crypto/md5"
	"image"
)

const (
	landscape byte = iota
	portrait
)

func compareSig(src []byte, sig [16]byte) bool {
	if len(src) != 16 {
		return false
	}
	for i := 0; i < 16; i++ {
		if src[i] != sig[i] {
			return false
		}
	}
	return true
}

func isPortraitLeft(content []byte) bool {
	sig := []byte{83, 234, 230, 173, 67, 108, 25, 219, 155, 106, 67, 4, 203, 188, 104, 255}
	return compareSig(sig, md5.Sum(content[2519641:2519679]))
}

func isLandscapeLeft(content []byte) bool {
	sig := []byte{55, 233, 72, 197, 15, 178, 181, 248, 108, 134, 188, 68, 38, 181, 70, 238}
	return compareSig(sig, md5.Sum(content[114121:114158]))
}

func isPortraitRight(content []byte) bool {
	sig := []byte{41, 84, 59, 72, 139, 237, 134, 28, 71, 156, 27, 161, 150, 96, 231, 125}
	return compareSig(sig, md5.Sum(content[112250:112287]))
}

func isLandscapeRight(content []byte) bool {
	sig := []byte{49, 255, 140, 250, 121, 227, 110, 1, 57, 124, 166, 197, 52, 4, 134, 253}
	return compareSig(sig, md5.Sum(content[112369:112407]))
}

type rotation struct {
	orientation byte
	isActive    bool
}

func (r *rotation) rotate(img *image.Gray) {
	if !r.isActive {
		return
	}
	switch {
	case (isPortraitLeft(img.Pix) || isPortraitRight(img.Pix)) && r.orientation != portrait:
		r.orientation = portrait
	case (isLandscapeLeft(img.Pix) || isLandscapeRight(img.Pix)) && r.orientation != landscape:
		r.orientation = landscape
	}
	if r.orientation == portrait {
		rotate(img)
	}
}

func rotate(img *image.Gray) {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	l := len(img.Pix)
	out := make([]uint8, l)
	for i := 0; i < l; i++ {
		j := w*(i%h+1) - i/h - 1
		out[i] = img.Pix[j]
	}
	(*img).Pix = out
	(*img).Rect = image.Rectangle{Max: image.Point{X: h, Y: w}}
	(*img).Stride = h

}
