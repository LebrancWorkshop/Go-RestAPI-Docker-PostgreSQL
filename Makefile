build:
	docker compose build 

run:
	docker compose up 

bash:
	docker compose run --service-ports web bash 