DIR="internal/migrations"
DB="postgres://postgres:postgres@localhost:5433/bsa?sslmode=disable"

migrate-create:
# 	migrate create (comando) -ext (driver .sql, .pgx) -dir (diretorio) -seq (sequencia)
	migrate create -ext sql -dir $(DIR) -seq $(n)

migrate-up: 
# 	migrate -path (de onde) -database (para onde) up N ou --all
	migrate -path $(DIR) -database $(DB) up $(n)

migrate-down: 
# 	migrate -path (de onde) -database (para onde) down N ou --all
	migrate -path $(DIR) -database $(DB) down $(n)

migrate-version:
	migrate -path $(DIR) -database $(DB) version

migrate-goto:
	migrate -path $(DIR) -database $(DB) goto $(v)

migrate-force:
	migrate -path $(DIR) -database $(DB) force $(v)
