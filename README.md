# STORE-PROJECT
## API сервис, эмулятор сервиса транзакций
### Запуск
Чтобы собрать проект, необходимо создать `.env` файл в корне директории со следующим содержанием:

```
export GIN_MODE=release
POSTGRES_PASSWORD=qwerty
POSTGRES_USER=postgres
POSTGRES_DB=postgres
``` 

Для запуска необходимо в папке проекта запустить сборку контейнеров:
в первый раз `docker-compose up --build`, в последующие можно использовать `docker-compose up`. Если нужно запустить сервис в фоне, использовать `docker-compose up -d`

При первом запуске необходимо инициализировать базу данных при помощи пакета `golang-migrate`:
1. `go install github.com/golang-migrate/migrate`
2. `migrate -path internal/handler/infrastructure/repository/schema -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/postgres?sslmode=disable' up` , где `localhost:5432` - адрес postgres

### Swagger
Swagger доступен по адресу `localhost:8000/swagger/index.html`

### Схема изменения состояний
![alt text]()