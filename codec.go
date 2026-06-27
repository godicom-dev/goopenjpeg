package goopenjpeg

// Codec selects the JPEG 2000 container format (pylibjpeg-openjpeg codec argument).
type Codec int

const (
	CodecJ2K Codec = 0 // codestream (.j2k, .jpc, .j2c)
	CodecJPT Codec = 1 // JPT-stream
	CodecJP2 Codec = 2 // JP2 file format (.jp2)
)
