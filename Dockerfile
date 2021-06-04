FROM golang:latest
RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get https://github.com/harrycoa/apiMatriz/main

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/app
COPY . .
RUN go get "github.com/labstack/echo"
RUN go build main.go

FROM scratch
COPY --from=builder /go/src .
ENTRYPOINT ["./main"]