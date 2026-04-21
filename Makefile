BINARY=kenvctl

all: build

build:
	go build -o $(BINARY).exe ./cmd

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY) ./cmd

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY) ./cmd

clean:
	rm -f $(BINARY) $(BINARY).exe

