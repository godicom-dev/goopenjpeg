# goopenjpeg

Go JPEG 2000 decoder — **no CGO** for callers (`purego` + embedded native library).

Aligned with [pylibjpeg-openjpeg](https://github.com/pydicom/pylibjpeg-openjpeg) for DICOM transfer syntaxes:

| UID | Description |
|-----|-------------|
| 1.2.840.10008.1.2.4.90 | JPEG 2000 Lossless Only |
| 1.2.840.10008.1.2.4.91 | JPEG 2000 |
| 1.2.840.10008.1.2.4.201–203 | HTJ2K |

## Status

**Phase 1 (current):** project scaffold + decode API + CI native builds.

- Done: `DecodeImage`, `GetImageParameters`, `DecodePixelData`, purego loader
- Next: compliance tests from `ref/pylibjpeg-openjpeg`, `encode` API (Phase 2)

## API

```go
func DecodeImage(stream any, codec Codec) (*Image, error)
func GetImageParameters(stream any, codec Codec) (*Params, error)
func DecodePixelData(src []byte, opts PixelDataOptions) ([]byte, error)
func OpenJPEGVersion() (string, error)
```

`Codec`: `CodecJ2K` (0), `CodecJPT` (1), `CodecJP2` (2).

## Layout

```
goopenjpeg/           # public Go API
native/               # purego + go:embed prebuilt libs
lib/
  openjpeg/           # submodule → uclouvain/openjpeg
  interface/          # decode glue (from pylibjpeg-openjpeg, memory streams)
  capi/               # C ABI for purego
ref/pylibjpeg-openjpeg/
```

## Development

```bash
git clone --recurse-submodules https://github.com/godicom-dev/goopenjpeg.git
cd goopenjpeg
go test ./...          # uses stub or CI-built libs in native/libs/
make build-native      # requires CMake
```

CI (`build.yml`): build-native → commit `native/libs/` on main → test → release on tags.

## References

- [golibjpeg](https://github.com/godicom-dev/golibjpeg) — same purego architecture for ISO 10918 / JPEG-LS
- [pylibjpeg-openjpeg](https://github.com/pydicom/pylibjpeg-openjpeg) — behaviour and tests reference
