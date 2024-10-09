include .env
export $(shell sed 's/=.*//' .env)

migrateup:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

gotodb:
	docker exec -it test_db psql -U $(DB_USER) -d $(DB_NAME)

.PHONY: migrateup migratedown gotodb
