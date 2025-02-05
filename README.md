# go-example-app

## Запуск тестов

```bash
go test ./...
```

## Архитектура

```mermaid
graph TD;
    A[HTTP Server] -->|Обрабатывает запросы от пользователей| B[Web Controller\n(HTTP Handlers)];
    B -->|Принимает запросы| C[Service Layer\n(Business Logic)];
    C -->|Взаимодействует с репозиториями| D[Repository Layer\n(Data Access)];
    D -->|Использует хранилище данных| E[Data Source\n(e.g., In-memory, DB)];

    F[Пользователь] -->|Отправляет запросы| A
```
