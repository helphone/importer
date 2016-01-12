all: build

create-builder:
	docker build -t helphone/importer-builder -f scripts/Dockerfile.build .

build:
	docker run --name builder -v $$(pwd):/usr/lib/go/src/github.com/helphone/importer helphone/importer-builder
	@docker cp builder:/usr/lib/go/src/github.com/helphone/importer/importer ./importer
	@docker rm builder
	docker build -t helphone/importer -f scripts/Dockerfile .
	@rm importer

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
	-docker run -it --rm --env-file ./.env --link db_importer_test:db --name importer_test -v $$(pwd):/usr/lib/go/src/github.com/helphone/importer helphone/importer-builder /bin/sh -c "./scripts/test.sh"

test: mount-test cleanup

cleanup:
	@echo "Cleanup in progress..."
	@-docker rm -f db_importer db_importer_test > /dev/null 2>&1
