.PHONY: clean test

apprun-cli: go.* *.go
	go build -o $@ cmd/mugikuru/main.go

clean:
	rm -rf mugikuru dist/

test:
	go test -v ./...

install:
	go install github.com/ktamamu/mugikuru/cmd/mugikuru

dist:
	goreleaser build --snapshot --clean