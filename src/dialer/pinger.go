package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	send("ts.cld", "9090", "8001")
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
	i := 0
	for {
		conn, err := net.DialUDP("udp", &lAddr, &rAddr)
		if err != nil {
			fmt.Printf("Could not Connect to remote Address: %s\n", err)
		}
		fmt.Print("send: ")
		time.Sleep(1 * time.Second)
		text := "no trouble " + strconv.Itoa(i)
		i++
		fmt.Println(text)
		jtext, err := json.Marshal(text)
		if err != nil {
			fmt.Printf("Could not parse text to json: %s\n", err)
		}
		conn.Write(jtext)
		conn.Close()
	}
}
