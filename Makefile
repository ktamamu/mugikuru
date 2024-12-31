.PHONY: clean test

apprun-cli: go.* *.go
	go build -o $@ main.go

clean:
	rm -rf mugikuru dist/

test:
	go test -v ./...

install:
	go install github.com/ktamamu/mugikuru

dist:
	goreleaser build --snapshot --clean