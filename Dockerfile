# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server
RUN go build -o app

# Финальный образ
FROM alpine:latest

RUN apk add --no-cache postgresql-client

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/cmd/server/app /app/app

# Копируем миграции из корня проекта
COPY --from=builder /app/migrations /app/migrations

# Копируем sh скрипт ожидания подключения к БД и выставляем права на файл
COPY wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh

CMD ["/wait-for-postgres.sh", "db", "/app/app"]
