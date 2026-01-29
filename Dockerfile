FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o app

EXPOSE 8090

CMD ["./app"]