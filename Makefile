BINARY_NAME=personal-finances
BUILD_DIR=build
RELEASE_DIR=release

.PHONY: build run clean release debug linux windows

build:
ifeq ($(shell uname -s), Linux)
	$(MAKE) linux
else
	$(MAKE) windows
endif

run:
	@$(RELEASE_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH)

clean:
	@rm -f $(BUILD_DIR)/*
	@rm -f $(RELEASE_DIR)/*

release:
	@$(MAKE) BUILD_MODE=release build

debug:
	@$(MAKE) BUILD_MODE=debug build

linux:
	@GOOS=linux GOARCH=amd64 $(MAKE) build

windows:
	@GOOS=windows GOARCH=amd64 $(MAKE) build

ifeq ($(BUILD_MODE),release)
GO_GCFLAGS=-trimpath=$(GOPATH)
GO_LDFLAGS="-s -w"
endif

ifeq ($(BUILD_MODE),debug)
GO_GCFLAGS=-N -l
endif
