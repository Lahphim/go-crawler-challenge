.PHONY: dev test

install-dependencies:
	go get github.com/beego/bee/v2
	go mod tidy

dev:
	bee run

test:
	go test -v -p 1 ./...
