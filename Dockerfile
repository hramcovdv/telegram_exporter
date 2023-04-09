FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /app ./...

EXPOSE 2112

ENTRYPOINT ["/app/telegram_exporter"]
