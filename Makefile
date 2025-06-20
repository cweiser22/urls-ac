dev:
	sudo docker compose -f docker-compose.yml -f docker-compose-dev.yml up app prometheus grafana postgres redis

build-prod:
	sudo docker compose -f docker-compose.yml build

prod:
	sudo docker compose -f docker-compose.yml up

build-dev:
	sudo docker compose -f docker-compose.yml -f docker-compose-dev.yml build

clean:
	sudo docker compose -f docker-compose.yml -f docker-compose-dev.yml down --remove-orphans

test:
	cd app && go test -v ./...

build-frontend:
	cd frontend && npm install && npm run build && cd ..

make migrate-up:
	migrate -database "postgresql://postgres:postgres@localhost:5433/url_shortener?sslmode=disable" -source file://./migrations up
