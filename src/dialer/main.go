package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	go send(os.Args[1], os.Args[2], os.Args[3])
	for {
	}
}

func send(remoteAddress, rPort, lPort string) {
	intrPort, err := strconv.Atoi(rPort)
	if err != nil {
		fmt.Printf("Could not convert imput port to integer: %s\n", err)
	}
	intlPort, err := strconv.Atoi(lPort)
	if err != nil {
		fmt.Printf("Could not convert imput port to integer: %s\n", err)
	}

	rAddr := net.UDPAddr{
		Port: intrPort,
		IP:   net.ParseIP(remoteAddress),
	}
	lAddr := net.UDPAddr{
		Port: intlPort,
		IP:   nil,
	}
	for {
		conn, err := net.DialUDP("udp", &lAddr, &rAddr)
		if err != nil {
			fmt.Printf("Could not Connect to remote Address: %s\n", err)
		}
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
