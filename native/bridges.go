package native

import "unsafe"

type DecodeResult struct {
	Output      []byte
	Width       int
	Height      int
	Components  int
	Precision   int
	IsSigned    bool
	ColourSpace int
}

type ImageParams struct {
	Width       int
	Height      int
	Components  int
	Precision   int
	IsSigned    bool
	ColourSpace int
}

var decodingErrors = map[int32]string{
	1: "failed to create the input stream",
	2: "failed to setup the decoder",
	3: "failed to read the header",
	4: "failed to set the component indices",
	5: "failed to set the decoded area",
	6: "failed to decode image",
	7: "support for more than 32-bits per component is not implemented",
	8: "failed to upscale subsampled components",
}

func Decode(data []byte, codec int) (*DecodeResult, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	var outputPtr unsafe.Pointer
	var outputLen int32
	var width, height, components, precision, isSigned int32

	code := decodeFn(
		unsafe.Pointer(&data[0]),
		int32(len(data)),
		int32(codec),
		&outputPtr,
		&outputLen,
		&width,
		&height,
		&components,
		&precision,
		&isSigned,
	)
	if code != 0 {
		return nil, errWithCode("Decode()", code)
	}
	if outputPtr == nil || outputLen == 0 {
		return nil, errEmptyOutput
	}

	out := make([]byte, outputLen)
	copy(out, unsafe.Slice((*byte)(outputPtr), outputLen))
	freeFn(outputPtr)

	return &DecodeResult{
		Output:      out,
		Width:       int(width),
		Height:      int(height),
		Components:  int(components),
		Precision:   int(precision),
		IsSigned:    isSigned != 0,
		ColourSpace: 0,
	}, nil
}

func GetParameters(data []byte, codec int) (*ImageParams, error) {
	if len(data) == 0 {
		return nil, errEmptyInput
	}

	var width, height, components, precision, isSigned, colourspace int32

	code := getParamsFn(
		unsafe.Pointer(&data[0]),
		int32(len(data)),
		int32(codec),
		&width,
		&height,
		&components,
		&precision,
		&isSigned,
		&colourspace,
	)
	if code != 0 {
		return nil, errWithCode("GetParameters()", code)
	}

	return &ImageParams{
		Width:       int(width),
		Height:      int(height),
		Components:  int(components),
		Precision:   int(precision),
		IsSigned:    isSigned != 0,
		ColourSpace: int(colourspace),
	}, nil
}

func Version() (string, error) {
	buf := make([]byte, 64)
	code := versionFn(unsafe.Pointer(&buf[0]), int32(len(buf)))
	if code != 0 {
		return "", errWithCode("Version()", code)
	}
	n := 0
	for n < len(buf) && buf[n] != 0 {
		n++
	}
	return string(buf[:n]), nil
}
