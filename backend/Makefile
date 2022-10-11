CUR_DIR = backend

backend_start:
	cd $(CUR_DIR) && docker-compose up -d --build

backend_start_force:
	cd $(CUR_DIR) && docker-compose up -d --no-deps --build

backend_stop:
	cd $(CUR_DIR) && docker-compose down

migrate_up_dev:
	cd $(CUR_DIR) && migrate -database "postgres://photoclub:photoclub_pass@localhost:5432/photoclub_db?sslmode=disable" -path db/migrations up

migrate_down_dev:
	cd $(CUR_DIR) && migrate -database "postgres://photoclub:photoclub_pass@localhost:5432/photoclub_db?sslmode=disable" -path db/migrations down

swag:
	cd $(CUR_DIR) && swag init -g cmd/main.go