all: build

build:
	GOOS=linux GO15VENDOREXPERIMENT=1 go build -o importer main.go
	docker build -t helphone/importer .
	@rm ./importer

build-for-test:
	@docker build -t helphone/importer_test -f Dockerfile.test .

mount:
	@echo "Mount the database"
	docker run -d --name db_importer -p 5432:5432 helphone/database
	@sleep 8
	@echo "Start the importer"
	-docker run --rm --env-file ./.env --link db_importer:db helphone/importer
	@docker rm -f db_importer

up: build mount

mount-test:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db_importer_test helphone/database > /dev/null 2>&1
	@sleep 8
	@echo "Launch tests"
	-docker run -it --rm --name importer_test --env-file ./.env --link db_importer_test:db helphone/importer_test

test: build-for-test mount-test cleanup

cleanup:
	@echo "Cleanup in progress..."
	@-docker rm -f db_importer db_importer_test > /dev/null 2>&1
