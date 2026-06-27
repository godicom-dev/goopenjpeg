package goopenjpeg

import (
	"encoding/binary"
	"math"
)

// Image holds decoded pixel data in native precision, planar-interleaved (DICOM order).
type Image struct {
	Pixels     []byte
	Width      int
	Height     int
	Components int
	Precision  int
	IsSigned   bool
	ColourSpace int
}

// Params holds JPEG 2000 image parameters without decoding pixels.
type Params struct {
	Width       int
	Height      int
	Components  int
	Precision   int
	IsSigned    bool
	ColourSpace int
}

func (p *Params) Rows() int        { return p.Height }
func (p *Params) Columns() int     { return p.Width }
func (p *Params) NrComponents() int { return p.Components }

func (img *Image) BytesPerSample() int {
	prec := img.Precision
	if prec <= 8 {
		return 1
	}
	if prec <= 16 {
		return 2
	}
	if prec <= 24 {
		return 3
	}
	return int(math.Ceil(float64(prec) / 8))
}

func (img *Image) offset(y, x, c int) int {
	bps := img.BytesPerSample()
	return (y*img.Width+x)*img.Components*bps + c*bps
}

func (img *Image) ByteAt(y, x, c int) byte {
	return img.Pixels[img.offset(y, x, c)]
}

func (img *Image) Uint16At(y, x, c int) uint16 {
	off := img.offset(y, x, c)
	return binary.LittleEndian.Uint16(img.Pixels[off:])
}
