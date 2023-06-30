ifeq ($(OS),Windows_NT)
	OSFLAG := windows
	ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
		ARCHFLAG := -D amd64
	else
		ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
			ARCHFLAG := amd64
		endif
		ifeq ($(PROCESSOR_ARCHITECTURE),x86)
			ARCHFLAG := 386
		endif
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSFLAG := linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OSFLAG := darwin
	endif
	UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		ARCHFLAG := amd64
	endif
	ifneq ($(filter %86,$(UNAME_P)),)
		ARCHFLAG := 386
	endif
	ifneq ($(filter aarch64%,$(UNAME_P)),)
		ARCHFLAG := arm64
	endif
endif

MAKEPATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAIN := $(patsubst %/,%,$(dir $(MAKEPATH)))/main

all:
	cd main; GOOS=$(OSFLAG) GOARCH=$(ARCHFLAG) go build -o ../bin/file_conv