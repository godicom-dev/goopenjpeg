package native

import (
	"os"
	"path/filepath"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

var (
	decodeFn func(
		data unsafe.Pointer, dataLen int32, codec int32,
		output *unsafe.Pointer, outputLen *int32,
		width, height, components, precision *int32,
		isSigned *int32,
	) int32
	getParamsFn func(
		data unsafe.Pointer, dataLen int32, codec int32,
		width, height, components, precision *int32,
		isSigned *int32, colourspace *int32,
	) int32
	versionFn func(buf unsafe.Pointer, bufLen int32) int32
	freeFn    func(p unsafe.Pointer)
)

func extractAndLoad(path string) (uintptr, error) {
	if err := os.WriteFile(path, libData, 0o755); err != nil {
		return 0, err
	}
	handle, err := loadLibrary(path)
	if err != nil {
		return 0, err
	}
	if runtime.GOOS != "windows" {
		_ = os.Remove(path)
	}
	return handle, nil
}

func init() {
	tmp := filepath.Join(os.TempDir(), "goopenjpeg."+libExt())
	handle, err := extractAndLoad(tmp)
	if err != nil {
		panic("goopenjpeg: failed to load native library: " + err.Error())
	}
	purego.RegisterLibFunc(&decodeFn, uintptr(handle), "goopenjpeg_decode")
	purego.RegisterLibFunc(&getParamsFn, uintptr(handle), "goopenjpeg_get_parameters")
	purego.RegisterLibFunc(&versionFn, uintptr(handle), "goopenjpeg_version")
	purego.RegisterLibFunc(&freeFn, uintptr(handle), "goopenjpeg_free")
}

func libExt() string {
	switch runtime.GOOS {
	case "windows":
		return "dll"
	case "darwin":
		return "dylib"
	default:
		return "so"
	}
}
