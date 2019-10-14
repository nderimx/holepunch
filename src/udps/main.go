package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

var (
	conn    net.UDPConn
	clients []client
)

type client struct {
	ID       string
	Endpoint string
}

func listen() {
	addr := net.UDPAddr{
		Port: 9090,
		IP:   nil,
	}
	conn, err := net.ListenUDP("udp", &addr) // code does not block here
	if err != nil {
		fmt.Printf("Could not listen on port 9090: %s\n", err)
		return
	}
	defer conn.Close()

	var buf [1024]byte
	for {
		_, remoteAddr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("Could not Read from UDP: %s\n", err)
		}

		fmt.Printf("%s\n%s\n", remoteAddr, buf)
	}
}

func send() {
	for {
		fmt.Print("send: ")
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("Could not read line: %s\n", err)
		}
		text := string(line)
		jtext, err := json.Marshal(text)
		if err != nil {
			fmt.Printf("Could not parse text to json: %s\n", err)
		}
		conn.Write(jtext)
		conn.Close()
	}
}

func main() {
	go listen()
	send()
}
