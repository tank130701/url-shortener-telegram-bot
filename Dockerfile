# Используем образ с Go для сборки приложения
FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# Сборка Go приложения
RUN go mod download
RUN go build -o telegram-bot ./cmd/main.go

CMD ["./telegram-bot"]