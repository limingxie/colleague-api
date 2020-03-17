FROM pangpanglabs/golang:builder AS builder
WORKDIR /go/src/hublabs/colleague-api
COPY ./ /go/src/hublabs/colleague-api
# disable cgo 
ENV CGO_ENABLED=0
# build steps
ENV GOPROXY=https://goproxy.cn
ENV GO111MODULE=on
RUN echo ">>> 1: go version" && go version \
    && echo ">>> 2: go get" && go get -v -d \
    && echo ">>> 3: go install" && go install

# make application docker image use alpine
FROM pangpanglabs/alpine-ssl

WORKDIR /go/bin/
# copy execute file to image
COPY --from=builder /go/bin/colleague-api ./
EXPOSE 8001

CMD ["./colleague-api", "api-server"]