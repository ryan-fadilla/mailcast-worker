# mailcast-worker

## Development

Please install or update your golang version to  1.21 or latest

Please follow this command step by step :

```
go mod init mailcast-worker
go mod tidy
go run cmd/main.go
```

## Docker

Please follow this command step by step :

```
docker build -t mailcast-worker .
docker run mailcast-worker:latest
```

## Git Dependency

- [Get started with asynqmon](https://github.com/hibiken/asynqmon)
- [Get started with asynq](https://github.com/hibiken/asynq)
