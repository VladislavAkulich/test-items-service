COMPOSE_FILE = ./docker/docker-compose.yaml

.PHONY: db_up
db_up:
	docker-compose -f $(COMPOSE_FILE) up --build -d

.PHONY: db_stop
db_stop:
	docker-compose -f $(COMPOSE_FILE) stop
