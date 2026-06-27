package goopenjpeg

import "github.com/godicom-dev/goopenjpeg/native"

// OpenJPEGVersion returns the linked openjpeg library version.
func OpenJPEGVersion() (string, error) {
	return native.Version()
}

// DecodeImage decodes JPEG 2000 data (pylibjpeg-openjpeg decode()).
func DecodeImage(stream any, codec Codec) (*Image, error) {
	data, err := ReadStream(stream)
	if err != nil {
		return nil, err
	}
	res, err := native.Decode(data, int(codec))
	if err != nil {
		return nil, err
	}
	return &Image{
		Pixels:      res.Output,
		Width:       res.Width,
		Height:      res.Height,
		Components:  res.Components,
		Precision:   res.Precision,
		IsSigned:    res.IsSigned,
		ColourSpace: res.ColourSpace,
	}, nil
}

// GetImageParameters reads JPEG 2000 parameters without decoding pixels.
func GetImageParameters(stream any, codec Codec) (*Params, error) {
	data, err := ReadStream(stream)
	if err != nil {
		return nil, err
	}
	p, err := native.GetParameters(data, int(codec))
	if err != nil {
		return nil, err
	}
	return &Params{
		Width:       p.Width,
		Height:      p.Height,
		Components:  p.Components,
		Precision:   p.Precision,
		IsSigned:    p.IsSigned,
		ColourSpace: p.ColourSpace,
	}, nil
}

// Decode is a shorthand for DecodeImage with CodecJ2K.
func Decode(data []byte) (*Image, error) {
	return DecodeImage(data, CodecJ2K)
}

// GetParameters is a shorthand for GetImageParameters with CodecJ2K.
func GetParameters(data []byte) (*Params, error) {
	return GetImageParameters(data, CodecJ2K)
}
