# ---------- build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# копируем сервис
COPY go-service ./go-service

# переходим в модуль
WORKDIR /app/go-service

# скачиваем зависимости
RUN go mod tidy

# собираем бинарник
RUN go build -o server ./cmd/server

# ---------- runtime stage ----------
FROM alpine:3.18

WORKDIR /app

# копируем бинарник
COPY --from=builder /app/go-service/server .

EXPOSE 8080

CMD ["./server"]
