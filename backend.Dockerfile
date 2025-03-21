FROM golang:1.24-alpine3.20

WORKDIR /app

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go mod download
RUN go build -o main . && chmod +x main

EXPOSE 8080

CMD ["./main"]
