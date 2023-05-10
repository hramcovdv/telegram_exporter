BINARY_NAME=telegram_exporter

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean