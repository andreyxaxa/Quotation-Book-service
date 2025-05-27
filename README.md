# Мини-сервис "Цитатник"

## REST API сервис для хранения и управления цитатами.

### Функционал:
1. Добавление новой цитаты (POST /quotes)
2. Получение всех цитат (GET /quotes)
3. Получение случайной цитаты (GET /quotes/random)
4. Фильтрация по автору (GET /quotes?author=Confucius)
5. Удаление цитаты по ID (DELETE /quotes/{id})

### Проверочные команды (curl):
```
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

```
curl http://localhost:8080/quotes
```

```
curl http://localhost:8080/quotes/random
```

```
curl http://localhost:8080/quotes?author=Confucius
```

```
curl -X DELETE http://localhost:8080/quotes/1
```

## Детали 

- У сервиса есть конфиг - [config/config.go](https://github.com/andreyxaxa/Quotation-Book-service/blob/main/config/config.go); Читается из `.env` файла. В рамках тестового задания .env прямо в репозитории, очевидно в проде он должен быть заигнорен.
- В слое хэндлеров применяется версионирование - [internal/controller/http/v1](https://github.com/andreyxaxa/Quotation-Book-service/tree/main/internal/controller/http/v1).
  Для версии v2 нужно будет просто добавить папку `http/v2` с таким же содержимым, в файле [internal/controller/http/router.go](https://github.com/andreyxaxa/Quotation-Book-service/blob/main/internal/controller/http/router.go) добавить строку:
```go
{
    v1.NewQuotesRoutes(r, q)
}

{
    v2.NewQuotesRoutes(r, q)
}
```
Я не стал добавлять префикс `/v1`, чтобы приведённые проверочные команды (URL) никак не изменялись.
- Используется dependency injection - [internal/controller/http/v1/controller.go](https://github.com/andreyxaxa/Quotation-Book-service/blob/main/internal/controller/http/v1/controller.go).
- Реализован graceful shutdown - [internal/app/app.go](https://github.com/andreyxaxa/Quotation-Book-service/blob/main/internal/app/app.go).
- Удобная и гибкая конфигурация HTTP сервера - [pkg/httpserver/options.go](https://github.com/andreyxaxa/Quotation-Book-service/blob/main/pkg/httpserver/options.go).
  Позволяет конфигурировать сервер в конструкторе таким образом:
```go
httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
```

## Запуск

### Local:
Клонируем репозиторий, выполняем:
```
make run
```

### Docker:
Клонируем репозиторий, выполняем:
```
make compose-up
```

## Тесты

```
make test
```

## Прочие `make` команды
Зависимости:
```
make deps
```
docker compose down:
```
make compose-down
```
