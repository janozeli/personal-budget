BINARY_NAME=personal-budget
RELEASE_DIR=release

GOARCH = amd64
BUILD_MODE=debug
GO_GCFLAGS=-N -l
ifeq ($(shell uname -s),Linux)
	GOOS := linux
else
	ifeq ($(OS),Windows_NT)
		GOOS := windows
	else
		$(error Sistema operacional não suportado)
	endif
endif

.PHONY: run clean release debug

run: clean build
	@$(RELEASE_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH)-$(BUILD_MODE)

build: clean
	@go build -gcflags="$(GO_GCFLAGS)" -ldflags="$(GO_LDFLAGS)" -o $(RELEASE_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH)-$(BUILD_MODE)

clean:
	@rm -f $(RELEASE_DIR)/*