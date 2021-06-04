FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=off \
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