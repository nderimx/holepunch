package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

var (
	clients []client
)

type client struct {
	ID       string
	Endpoint string
}

func serveCommunicate() {
	lAddr := net.UDPAddr{
		Port: 9090,
		IP:   nil,
	}
	conn, err := net.ListenUDP("udp", &lAddr) // code does not block here
	if err != nil {
		fmt.Printf("Could not listen on port 9090: %s\n", err)
	}
	var buf [1024]byte
	_, rAddr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Printf("Could not Read from UDP: %s\n", err)
	}
	fmt.Printf("\nrec: %s\tfrom: %s\n", buf, rAddr)
	conn.Close()
	conn, err = net.DialUDP("udp", &lAddr, rAddr)
	if err != nil {
		fmt.Printf("Could not Connect to remote Address: %s\n", err)
	}
	defer conn.Close()
	go listen(conn)
	send(conn)
}

func listen(conn *net.UDPConn) {
	for {
		var buf [1024]byte
		_, remoteAddr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("Could not Read from UDP: %s\n", err)
		}
		fmt.Printf("\nrec: %s\tfrom: %s\n", buf, remoteAddr)
	}
}

func send(conn *net.UDPConn) {
	for {
		fmt.Printf("send: ")
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
	}
}

func main() {
	serveCommunicate()
}
