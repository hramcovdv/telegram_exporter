FROM golang:1.20-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o telegram_exporter main.go

FROM alpine:3.17

COPY --from=build /src/telegram_exporter /bin/telegram_exporter

EXPOSE 2112

USER nobody

ENTRYPOINT ["/bin/telegram_exporter"]
