FROM golang:1.22

RUN mkdir app
WORKDIR ./app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o bin/server ./cmd/server

CMD ["./bin/server"]