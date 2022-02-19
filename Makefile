.PHONY: up
up:
	docker build -t go-docker-template . && docker run -itd --rm --name go-docker-template go-docker-template

.PHONY: down
down:
	docker stop go-docker-template