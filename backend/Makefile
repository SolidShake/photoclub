CUR_DIR = backend

backend_start:
	cd $(CUR_DIR) && docker-compose up -d --no-deps --build

backend_stop:
	cd $(CUR_DIR) && docker-compose down

swag:
	cd $(CUR_DIR) && swag init -g cmd/main.go