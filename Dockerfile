FROM golang:1.13 AS builder

RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/github.com/hublabs/colleague-api
ADD go.mod go.sum ./
RUN go mod download
ADD . /go/src/github.com/hublabs/colleague-api
ENV CGO_ENABLED=0
RUN go build -o colleague-api

FROM pangpanglabs/alpine-ssl
WORKDIR /go/src/github.com/hublabs/colleague-api
COPY --from=builder /go/src/github.com/hublabs/colleague-api ./

EXPOSE 8001

CMD ["./colleague-api", "api-server"]