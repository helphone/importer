all: build
up: build mount cleanup
test: mount-test cleanup

create-builder:
	@echo "The builder is in construction..."
	@docker build -t helphone/importer-builder -f scripts/Dockerfile.build .

build:
	@echo "Importer build started..."
	@docker run --rm -v $$(pwd):/go/src/github.com/helphone/importer helphone/importer-builder
	@echo "Build finished, the final docker image is in construction"
	@docker build -t helphone/importer -f scripts/Dockerfile .
	@rm importer
	@echo "Image construction and cleanup are finished"

mount:
	@echo "Mount the database..."
	@docker run -d --name db_importer -p 5432:5432 helphone/database
	@sleep 10
	@echo "Start the importer"
	-docker run --rm --env DB_USERNAME=postgres --env DB_PASSWORD=postgres --link db_importer:db helphone/importer

mount-test:
	@echo "Setup the environnement..."
	@echo "Mount the database"
	@docker run -d --name db_importer_test helphone/database > /dev/null 2>&1
	@sleep 10
	@echo "Launch tests"
	-docker run -it --rm --env DB_USERNAME=postgres --env DB_PASSWORD=postgres --link db_importer_test:db -v $$(pwd):/go/src/github.com/helphone/importer helphone/importer-builder /bin/sh -c "./scripts/test.sh"

cleanup:
	@echo "Cleanup in progress..."
	@-docker rm -f db_importer db_importer_test > /dev/null 2>&1
