FROM golang:1.22 as builder

RUN mkdir app
WORKDIR ./app

COPY . .

RUN go mod download
RUN go mod verify

ARG APP_VERSION
RUN go build -o bin/server -ldflags "-X main.Version=$APP_VERSION" ./cmd/server


FROM busybox

RUN mkdir /app

COPY --from=builder /go/app/bin/* /app

# copy any static files if needed

CMD ["/app/server"]