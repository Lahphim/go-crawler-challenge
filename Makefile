-include .env
export

.PHONY: envsetup dev db/setup db/migrate db/rollback install/package test test/run

install-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/ddollar/forego
	go mod tidy

envsetup:
	make db/setup
	make install/package

dev:
	forego start

db/setup:
	docker-compose -f docker-compose.dev.yml up -d
	make db/migrate

db/migrate:
	bee migrate -driver=postgres -conn="$(DATABASE_URL)"

db/rollback:
	bee migrate rollback -driver=postgres -conn="$(DATABASE_URL)"

install/package:
	npm i

test:
	make test/run ENV=test

test/run:
	docker-compose -f docker-compose.test.yml up -d
	APP_RUN_MODE=test go test -v -p 1 ./...
	docker-compose -f docker-compose.test.yml down
