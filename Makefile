# Docker Makefile
.PHONY: start
up:
	docker-compose down
	docker-compose build
	docker-compose up -d

down:
	docker-compose down