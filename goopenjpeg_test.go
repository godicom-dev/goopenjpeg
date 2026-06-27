package goopenjpeg

import "testing"

func TestReadStreamBytes(t *testing.T) {
	data := []byte{0xff, 0x4f, 0xff, 0x51}
	got, err := ReadStream(data)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(data) {
		t.Fatalf("unexpected data")
	}
}

func TestDecodeInvalidStreamType(t *testing.T) {
	_, err := DecodeImage(nil, CodecJ2K)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestOpenJPEGVersion(t *testing.T) {
	version, err := OpenJPEGVersion()
	if err != nil {
		t.Fatal(err)
	}
	if version == "" {
		t.Fatal("expected version string")
	}
}

func TestDecodePixelDataMissingPhotometric(t *testing.T) {
	_, err := DecodePixelData([]byte{1, 2, 3}, PixelDataOptions{Version: PixelDataV1})
	if err == nil {
		t.Fatal("expected error")
	}
}
