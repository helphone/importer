all: build

build: export GOOS=linux
build:
	$(shell echo $$GOPATH)/bin/godep go build -o importer main.go
	docker build -t helphone/importer .
	@rm importer

build-for-test: export GO15VENDOREXPERIMENT=1
build-for-test:
	go vet $$(go list ./...|grep -v vendor)
	@docker build -t helphone/importer_test -f Dockerfile.test .

up:
	@echo "Mount the database"
	docker run -d --name db_importer -p 5432:5432 helphone/database
	@sleep 8
	@echo "Start the importer"
	-docker run --rm --env-file ./.env --link db_importer:db helphone/importer
	@docker rm -f db_importer

up-with-build: build up

up-test: export GO15VENDOREXPERIMENT=1
up-test:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db_importer_test helphone/database > /dev/null 2>&1
	@sleep 8
	@echo "Launch tests"
	-docker run -it --rm --name importer_test --env-file ./.env --link db_importer_test:db helphone/importer_test

test: build-for-test up-test cleanup

cleanup:
	@-docker rm -f db_importer db_importer_test > /dev/null 2>&1
