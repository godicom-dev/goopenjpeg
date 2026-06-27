//go:build linux && amd64

package native

import _ "embed"

//go:embed libs/goopenjpeg_linux_amd64.so
var libData []byte
