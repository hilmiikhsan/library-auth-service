FROM golang:1.22.8-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o library-auth-service

RUN chmod +x library-auth-service

EXPOSE 9090

EXPOSE 6000

CMD ["./library-auth-service"]
