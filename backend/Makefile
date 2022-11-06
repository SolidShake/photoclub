BACKEND_DIR = backend
DSN = "postgres://photoclub:photoclub_pass@localhost:5432/photoclub_db?sslmode=disable"

backend_start:
	cd $(BACKEND_DIR) && docker-compose up -d --build

backend_start_force:
	cd $(BACKEND_DIR) && docker-compose up -d --no-deps --build

backend_stop:
	cd $(BACKEND_DIR) && docker-compose down

migrate_up:
	cd $(BACKEND_DIR) && migrate -database $(DSN) -path db/migrations up

migrate_down:
	cd $(BACKEND_DIR) && migrate -database $(DSN) -path db/migrations down

migrate_drop:
	cd $(BACKEND_DIR) && migrate -database $(DSN) -path db/migrations drop

swag:
	cd $(BACKEND_DIR) && swag init -g cmd/main.go