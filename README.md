# go-example-app

## Запуск тестов

```bash
go test ./...
```

## Архитектура

```mermaid
graph TD;
    A[HTTP Server] -->|Обрабатывает запросы от пользователей| B[Web Controller<br>(HTTP Handlers)];
    B -->|Принимает запросы| C[Service Layer<br>(Business Logic)];
    C -->|Взаимодействует с репозиториями| D[Repository Layer<br>(Data Access)];
    D -->|Использует хранилище данных| E[Data Source<br>(e.g., In-memory, DB)];

    F[Пользователь] -->|Отправляет запросы| A
```
