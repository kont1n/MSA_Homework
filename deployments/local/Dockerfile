FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

# Устанавливаем место назначения для COPY
WORKDIR /app

# Копируем файлы go.mod и go.sum в WORKDIR
COPY go.mod go.sum ./

# Скачиваем необходимые Go модули (зависимости нашего проетка)
RUN go mod download

# Копируем все исходные go файлы нашего проекта в образ
COPY *.go ./

# Собираем бинарный файл нашего приложения
RUN go build -o /bin/main


FROM alpine:latest AS runner
WORKDIR /
COPY --from=builder /bin/main /main
CMD ["/main"]