FROM go:latest

WORKDIR /app
RUN mkdir -p /app/logs

COPY . .
RUN go build -o budget-api -ldflags="-s -w" cmd/api/main.go
CMD ["./budget-api"]