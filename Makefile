dev:
	sudo docker compose -f docker-compose.yml -f docker-compose-dev.yml up app prometheus grafana

build-prod:
	sudo docker compose -f docker-compose.yml build

prod:
	sudo docker compose -f docker-compose.yml up

build-dev:
	sudo docker compose -f docker-compose.yml -f docker-compose-dev.yml build

clean:
	sudo docker compose -f docker-compose.yml -f docker-compose-dev.yml down --remove-orphans