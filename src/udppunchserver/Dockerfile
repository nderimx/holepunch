FROM golang:1.8

WORKDIR /go/src/app
COPY tmpsrv.go .

RUN go build tmpsrv.go

CMD ["./tmpsrv"]

EXPOSE 9090:9090/udp