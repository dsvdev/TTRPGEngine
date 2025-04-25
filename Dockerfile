# Указываем базовый образ с Go
FROM golang:1.22.5 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum (для кэширования зависимостей)
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарный файл (предположим, что у вас main.go в корне)
RUN go build -o app ./cmd/bot

# Используем минимальный образ для продакшн (без Go)
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/app .

# Команда по умолчанию
CMD ["./app"]