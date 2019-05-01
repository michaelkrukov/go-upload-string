ALL: install

install:
	go install ./...

build:
	go build ./...

clean:
	rm $(GOPATH)/bin/upload-string
