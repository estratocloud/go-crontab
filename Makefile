go ?= 1.21

.PHONY: up run test coverage

help:
	@echo "Available options:"
	@echo "    make up"
	@echo "    make up go=$(go)"
	@echo "    make test"
	@echo "    make coverage"

up:
	docker rm -f go-crontab > /dev/null 2>&1
	docker build . -t go-crontab-image --build-arg GO_VERSION=$(go)
	docker run --interactive --detach --volume $(shell pwd):/app --name go-crontab go-crontab-image

test:
	docker exec -ti go-crontab go test ./crontab/...

coverage:
	docker exec -ti go-crontab go test ./crontab/... -coverprofile=dev/coverage.txt
	docker exec -ti go-crontab go tool cover -html=dev/coverage.txt -o=dev/coverage.html
	sudo chown $$USER dev/coverage.html
	firefox dev/coverage.html

update:
	echo "Check the GO version in the Dockerfile"
	docker exec -ti go-crontab go get -u ./crontab
	docker exec -ti go-crontab go mod tidy
