backend_start:
	docker-compose -f backend/docker-compose.yml up -d

backend_stop:
	docker-compose -f backend/docker-compose.yml down