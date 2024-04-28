DB_DSN ?= mysql://root:root@tcp(localhost:13306)/test_db

migration_create:
	migrate create -ext sql -dir db/migrations/ -seq new


migration_up:
	migrate -path db/migrations/ -database "$(DB_DSN)" -verbose up
