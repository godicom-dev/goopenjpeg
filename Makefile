GOHOSTOS ?= $(shell go env GOHOSTOS)
GOHOSTARCH ?= $(shell go env GOHOSTARCH)

BUILD_DIR := lib/build_nix
ifeq ($(GOHOSTOS),windows)
BUILD_DIR := lib/build
endif

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
	git submodule update --init --recursive
	cmake -S lib -B $(BUILD_DIR) \
	  -DCMAKE_BUILD_TYPE=Release \
	  -DCMAKE_POSITION_INDEPENDENT_CODE=ON
	cmake --build $(BUILD_DIR) --target goopenjpeg
	mkdir -p native/libs
ifeq ($(GOHOSTOS),windows)
	cp $(BUILD_DIR)/goopenjpeg.dll native/libs/goopenjpeg_$(GOHOSTARCH).dll
else ifeq ($(GOHOSTOS),darwin)
	cp $(BUILD_DIR)/libgoopenjpeg.dylib native/libs/goopenjpeg_darwin_$(GOHOSTARCH).dylib
else
	cp $(BUILD_DIR)/libgoopenjpeg.so native/libs/goopenjpeg_linux_$(GOHOSTARCH).so
endif

.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf lib/build lib/build_nix
