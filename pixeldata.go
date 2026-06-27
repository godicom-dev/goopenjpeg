package goopenjpeg

import "fmt"

// PixelDataVersion selects decode_pixel_data behaviour.
type PixelDataVersion int

const (
	PixelDataV1 PixelDataVersion = 1
	PixelDataV2 PixelDataVersion = 2
)

// PixelDataOptions configures DecodePixelData for DICOM handlers.
type PixelDataOptions struct {
	Version                   PixelDataVersion
	Codec                     Codec
	PhotometricInterpretation string
}

// DecodePixelData decodes encapsulated JPEG 2000 pixel data for DICOM.
// Version 2 returns raw decoded bytes; version 1 matches pylibjpeg v1 behaviour.
func DecodePixelData(src []byte, opts PixelDataOptions) ([]byte, error) {
	if opts.Version == 0 {
		opts.Version = PixelDataV1
	}
	if opts.Codec == 0 && opts.Version == PixelDataV1 {
		opts.Codec = CodecJ2K
	}

	switch opts.Version {
	case PixelDataV1:
		if opts.PhotometricInterpretation == "" {
			return nil, fmt.Errorf(
				"The (0028,0004) Photometric Interpretation element is missing from the dataset",
			)
		}
		img, err := DecodeImage(src, opts.Codec)
		if err != nil {
			return nil, err
		}
		return img.Pixels, nil
	case PixelDataV2:
		img, err := DecodeImage(src, opts.Codec)
		if err != nil {
			return nil, err
		}
		return img.Pixels, nil
	default:
		return nil, fmt.Errorf("goopenjpeg: unsupported pixel data version %d", opts.Version)
	}
}
