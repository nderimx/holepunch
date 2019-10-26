FROM golang:1.8

WORKDIR /go/src/app
COPY pinger.go .

RUN go build pinger.go

CMD ["./pinger"]

EXPOSE 8080:8001/udp