include .env

ifndef APP_VERSION
	APP_VERSION := latest
endif

IMAGE_TAG=ymanshur/go-restful:$(APP_VERSION)

run:
	clear && air
test:
	clear
	go test ./... -coverprofile=cover.out -p 1
	go tool cover -func=cover.out
ifeq ($(mode), html)
	go tool cover -html=cover.out
endif
up:
	docker compose up -d
stop:
	docker compose stop
down: stop
	docker compose down --rmi local --remove-orphans -v 
build:
	docker build . -t ${IMAGE_TAG}
push:
	docker push ${IMAGE_TAG}
	