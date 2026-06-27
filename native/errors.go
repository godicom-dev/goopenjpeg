package native

import "fmt"

var errEmptyInput = fmt.Errorf("goopenjpeg: empty input data")

var errEmptyOutput = fmt.Errorf("goopenjpeg: decode returned empty output")

type StatusError struct {
	Op   string
	Code int
}

func (e *StatusError) Error() string {
	code := e.Code
	if code < 0 {
		if msg, ok := decodingErrors[int32(-code)]; ok {
			return fmt.Sprintf("openjpeg error code '%d' returned from %s: %s", code, e.Op, msg)
		}
	}
	if e.Code < 0 {
		switch e.Code {
		case -1:
			return fmt.Sprintf("goopenjpeg: invalid parameter in %s", e.Op)
		case -2:
			return fmt.Sprintf("goopenjpeg: memory allocation failed in %s", e.Op)
		case -3:
			return fmt.Sprintf("goopenjpeg: decode failed in %s", e.Op)
		}
	}
	return fmt.Sprintf("openjpeg error code '%d' returned from %s", e.Code, e.Op)
}

func errWithCode(op string, code int32) error {
	if code == 0 {
		return nil
	}
	c := int(code)
	if c > 0 {
		c = -c
	}
	return &StatusError{Op: op, Code: c}
}
