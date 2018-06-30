.PHONY: install test

# install dependencies
install:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

# run unit tests
test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...