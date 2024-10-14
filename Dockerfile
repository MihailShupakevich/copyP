# Укажите базовый образ. Здесь используется образ Golang.
FROM golang:1.23 AS builder

# Установите рабочую директорию
WORKDIR /copy

# Скопируйте необходимые файлы
COPY go.mod go.sum ./
RUN go mod download

COPY cmd .
COPY config .
COPY internal .
COPY . .


# Соберите ваше Go-приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./main.go

# Создайте финальный образ с минимальным размером
FROM alpine:latest
WORKDIR /root/

# Скопируйте бинарник из образа сборки
COPY --from=builder /copy/myapp .

# Укажите команду для запуска приложения
CMD ["./myapp"]

# Откройте порт, если ваше приложение использует его
EXPOSE 8080
