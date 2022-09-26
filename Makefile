IMAGE_TAG=ymanshur/go-restful:latest

run:
	clear && air
test:
	clear
	go test ./... -coverprofile=cover.out
	go tool cover -func=cover.out
ifeq ($(mode), html)
	go tool cover -html=cover.out
endif
up:
	docker compose up -d
	docker compose start
stop:
	docker compose stop
down:
	docker compose down --rmi local --remove-orphans -v 
build:
	docker compose build
push:
	docker push ${IMAGE_TAG}
	