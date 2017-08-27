all: build
	/tmp/splashmake -v

build:
	go build -o /tmp/splashmake ./splashmake/*.go

darwin: build
	/tmp/splashmake -v -os darwin

linux: build
	/tmp/splashmake -v -os linux

test:
	go run example/main.go -v

windows: build
	/tmp/splashmake -v -os windows