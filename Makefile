OSFLAG
ARCHFLAG

build:
	ifeq ($(OS),Windows_NT)
		OSFLAG += windows
	else ifeq ($(OS),Darwin)
		OSFLAG += darwin
	else
		OSFLAG += linux
	endif

	ifeq ($(PROCESSOR_ARCHITECTURE),x86_64)
		ARCHFLAG += amd64
	else ifeq ($(OS),x86)
		ARCHFLAG += 386
	else
		OSFLAG += arm
	endif

	GOOS=$(OSFLAG) GOARCH=$(ARCHFLAG) go build -o file_conv main.go
