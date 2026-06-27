package goopenjpeg

import (
	"fmt"
	"io"
	"os"
)

// ReadStream reads JPEG 2000 data from bytes, a file path, or an io.Reader.
func ReadStream(stream any) ([]byte, error) {
	switch v := stream.(type) {
	case nil:
		return nil, fmt.Errorf(
			"invalid type 'nil' - must be the path to a JPEG 2000 file, a buffer containing the data or an open file-like",
		)
	case string:
		return os.ReadFile(v)
	case []byte:
		return v, nil
	case io.Reader:
		return io.ReadAll(v)
	default:
		return nil, fmt.Errorf(
			"invalid type '%T' - must be the path to a JPEG 2000 file, a buffer containing the data or an open file-like",
			stream,
		)
	}
}
