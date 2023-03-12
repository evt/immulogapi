FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o logapi cmd/server/main.go

EXPOSE 9090

CMD ["./logapi"]
