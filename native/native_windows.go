//go:build windows

package native

import "golang.org/x/sys/windows"

func loadLibrary(path string) (uintptr, error) {
	handle, err := windows.LoadLibrary(path)
	return uintptr(handle), err
}
