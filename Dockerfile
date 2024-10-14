#FROM ubuntu:latest
#LABEL authors="dunice"
#
#ENTRYPOINT ["top", "-b"]

# Укажите базовый образ. Здесь используется образ Golang.
FROM golang:1.20 AS builder

RUN apk --no-cache add bash git make gcc gettext
# Установите рабочую директорию
WORKDIR /copy

# Скопируйте необходимые файлы
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Соберите ваше Go-приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Создайте финальный образ с минимальным размером
FROM alpine:latest
WORKDIR /root/

# Скопируйте бинарник из образа сборки
COPY --from=builder /app/myapp .

# Укажите команду для запуска приложения
CMD ["./myapp"]

# Откройте порт, если ваше приложение использует его
EXPOSE 8080