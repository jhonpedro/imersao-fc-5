FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

RUN go mod tidy

CMD ["go", "run", "cmd/main.go"]
