help: .PHONY
	echo "Read the Makefile, im too lazy"

dependencies:
	npm install -g pnpm
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

local-setup:
	cp backend/.env.example backend/.env
	cd local-setup && docker-compose up -d
	make migrate
	cd local-setup && docker-compose down
	cd frontend/admin && pnpm install

admin-frontend:
	cd ./frontend/admin/ && pnpm start
	
admin-backend:
	cd ./backend/admin/ && go run .

m: 
	@echo migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./backend/db/migrations

migrate:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./backend/db/migrations up

mock:
	mock-dev custom -f ./backend/db/mock/mock_skeleton.yaml --database postgres -v --password password

test-build:
	docker build -t brunstadtv-admin-test:latest -f ./backend/Dockerfile.test ./backend --progress=plain

test:
	make test-build
	cd backend && go test ./... -v