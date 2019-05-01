ALL: install

install:
	go install ./...

build:
	go build ./...

test:
	go test ./...

clean:
	rm $(GOPATH)/bin/upload-string
