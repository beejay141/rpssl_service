FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o rpssl_service

EXPOSE 5000

CMD ["./rpssl_service"]