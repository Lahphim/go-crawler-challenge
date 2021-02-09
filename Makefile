.PHONY: dev test

install-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/ddollar/forego
	go mod tidy

dev:
	forego start

test:
	go test -v -p 1 ./...
