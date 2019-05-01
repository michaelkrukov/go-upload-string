ALL: install

install:
	go install ./...

build:
	go build ./...

test:
	go test ./... -race -coverprofile=cover.out -covermode=atomic -tags test

clean:
	rm $(GOPATH)/bin/upload-string
