.PHONY: envsetup dev test

install-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/ddollar/forego
	go mod tidy

envsetup:
	docker-compose -f docker-compose.dev.yml up -d
	npm i

dev:
	forego start

test:
	go test -v -p 1 ./...
