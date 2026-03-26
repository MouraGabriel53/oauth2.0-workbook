DIR="internal/migrations"
DB="postgres://postgres:postgres@postgres:5432/BSA?sslmode=disable"

migrate-create:
# 	migrate create (comando) -ext (driver .sql, .pgx) -dir (diretorio) -seq (sequencia)
	migrate create -ext sql -dir $(DIR) -seq $(name)

migrate-up: 
# 	migrate -path (de onde) -database (para onde) up N ou --all
	migrate -path $(DIR) -database $(DB) up $(number)

migrate-down: 
# 	migrate -path (de onde) -database (para onde) down N ou --all
	migrate -path $(DIR) -database $(DB) down $(number)