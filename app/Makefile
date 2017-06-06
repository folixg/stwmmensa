BINARY = stwmmensa
GOARCH = amd64
GOARCH_PI = arm
GOOS = linux

all: test linux

test:
	go test -v .

linux:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o ${BINARY} . ;

rpi:
	GOOS=${GOOS} GOARCH=${GOARCH_PI} go build -v -o ${BINARY}-${GOARCH_PI} . ;
