# go-example-app

## Test run

```bash
go test ./...
```


## Architecture

+---------------------------+
|       Web Controller      |
| (HTTP Handlers)           |
| - Receives requests       |
| - Interacts with services |
+---------------------------+
            |
            V
+---------------------------+
|       Service Layer       |
| - Business Logic          |
| - Interacts with repos    |
+---------------------------+
            |
            V
+---------------------------+
|     Repository Layer      |
|   (Data Access)          |
| - In-memory storage       |
| - Database access (if any)|
+---------------------------+
            |
            V
+---------------------------+
|        Data Source        |
| (e.g., In-memory, DB)    |
+---------------------------+
