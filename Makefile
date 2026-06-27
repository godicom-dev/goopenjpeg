GOHOSTOS ?= $(shell go env GOHOSTOS)
GOHOSTARCH ?= $(shell go env GOHOSTARCH)

.PHONY: embed-name
embed-name:
ifeq ($(GOHOSTOS),windows)
	@echo goopenjpeg_$(GOHOSTARCH).dll
else ifeq ($(GOHOSTOS),darwin)
	@echo goopenjpeg_darwin_$(GOHOSTARCH).dylib
else
	@echo goopenjpeg_linux_$(GOHOSTARCH).so
endif

.PHONY: build-native
build-native:
	git submodule update --init --recursive lib/openjpeg
	cmake -S lib -B lib/build \
	  -DCMAKE_BUILD_TYPE=Release \
	  -DCMAKE_POSITION_INDEPENDENT_CODE=ON
	cmake --build lib/build --config Release
	mkdir -p native/libs
	cp lib/build/libgoopenjpeg.* native/libs/

.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf lib/build lib/build_nix
