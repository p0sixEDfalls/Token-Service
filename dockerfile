FROM golang:1.12.4 AS builder

COPY . /go/src/github.com/p0sixEDfalls/ico/
WORKDIR /go/src/github.com/p0sixEDfalls/ico/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/ethereum/go-ethereum
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/lib/pq

RUN go build -o service ./src/main/main.go

FROM ubuntu:18.04
WORKDIR /root/
COPY --from=builder /go/src/github.com/p0sixEDfalls/ico/service .

ENTRYPOINT ["./service"]