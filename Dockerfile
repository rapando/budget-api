FROM golang:1.21.4-alpine

WORKDIR /app
RUN mkdir -p /app/logs

COPY . .
RUN go build -o budget-api -ldflags="-s -w" cmd/api/main.go
CMD ["./budget-api"]