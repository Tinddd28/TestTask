# Запуск приложения

### 1) В файле 'internal/service/song.go' в функции 'CreateSong' заменить 'example.com' на url, по которому совершается запрос к openapi
### 2) Установить нужные зависимости командой 'go mod tidy'
### 3) Запустить docker контейнер с БД (PostgreSQL) командой 'docker-compose up' (-d, чтобы запустить в фоне)
### 4) Сгенерировать swagger командой 'swag init -g cmd/main.go'
### 5) Создать таблицу в БД командой 'make createdb'
### 6) Применить миграции к БД командой 'make migrateup'
### 7) Запустить приложение командой 'go run cmd/main.go'