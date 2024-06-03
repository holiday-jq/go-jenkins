FROM golang:1.21.3 AS builder

MAINTAINER holiday

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

COPY go.mod go.sum ./
RUN go mod download
COPY ../ ./

RUN go build -o main ./main.go




FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/main


EXPOSE 11111

CMD ["./main"]

