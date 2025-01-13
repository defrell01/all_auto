# Используем базовый образ для Go
FROM golang:1.23-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git gcc musl-dev

# Копируем файлы для управления зависимостями
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарника
RUN go build -o market_service ./cmd/main.go

# Используем минимальный образ для исполнения
FROM alpine:latest

# Добавляем PostgreSQL клиент для работы с базой
RUN apk add --no-cache postgresql-client

# Копируем скомпилированный бинарник
COPY --from=builder /app/market_service /usr/local/bin/market_service

# Копируем конфигурации
COPY ./config/config.yaml /app/config/config.yaml

# Устанавливаем рабочую директорию
WORKDIR /app

# Экспортируем порт
EXPOSE 8080

# Команда для запуска приложения
CMD ["market_service"]
