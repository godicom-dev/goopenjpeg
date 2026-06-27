Native build tree for goopenjpeg.

- `openjpeg/` — git submodule ([uclouvain/openjpeg](https://github.com/uclouvain/openjpeg))
- `interface/` — decode glue adapted from pylibjpeg-openjpeg (memory streams, no Python)
- `capi/` — thin C ABI consumed by Go via purego

```bash
cmake -S lib -B lib/build -DCMAKE_BUILD_TYPE=Release
cmake --build lib/build
```

CI builds prebuilt artifacts into `native/libs/`.
