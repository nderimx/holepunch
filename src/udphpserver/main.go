package main

import "net"

type client struct {
	ID       string
	Endpoint string
}

type session struct {
	Client     client
	Connection net.UDPConn
}

var (
	conns   []session
	clients [][]client
)

func listenOnRegistrations() {

}

func main() {
}
