//go:build windows && amd64

package native

import _ "embed"

//go:embed libs/goopenjpeg_amd64.dll
var libData []byte
