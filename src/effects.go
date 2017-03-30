package main

import (
	"image"

	"github.com/disintegration/imaging"
)

/*Effect is an interface with 'apply' method with pointer receivers
to the effect that implements it*/
type Effect interface {
	apply(image.Image) *image.NRGBA
}

type blur struct {
	value float64
}

type rotate struct {
	angle float64
}

type contrast struct {
	value float64
}

type grayscale struct {
}

type sharpness struct {
	value float64
}

func (b blur) apply(destImg image.Image) *image.NRGBA {
	return imaging.Blur(destImg, b.value)
}

func (r rotate) apply(destImg image.Image) *image.NRGBA {
	if r.angle == 90 {
		return imaging.Rotate90(destImg)
	} else if r.angle == 270 {
		return imaging.Rotate270(destImg)
	}
	return imaging.Rotate180(destImg)
}

func (c contrast) apply(destImg image.Image) *image.NRGBA {
	return imaging.AdjustContrast(destImg, c.value)
}

func (g grayscale) apply(destImg image.Image) *image.NRGBA {
	return imaging.Grayscale(destImg)
}

func (s sharpness) apply(destImg image.Image) *image.NRGBA {
	return imaging.Sharpen(destImg, s.value)
}
