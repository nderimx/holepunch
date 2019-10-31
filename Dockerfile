FROM golang:1.8

WORKDIR /go/src/app
RUN mkdir client
RUN mkdir server
RUN printf "package main\nfunc main(){\nfor{}\n}" > empty.go
COPY src/udphpclient/main.go client
COPY src/udps/main.go server

RUN apt update -y
RUN apt install -y net-tools

RUN go build empty.go
RUN rm empty.go
CMD ["./empty"]