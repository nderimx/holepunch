package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func listen(port string) {
	intlPort, err := strconv.Atoi(port)
	if err != nil {
		fmt.Printf("Could not convert imput port to integer: %s\n", err)
	}
	addr := net.UDPAddr{
		Port: intlPort,
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
func communicate(remoteAddress, rPort, lPort string) {
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
	var conn *net.UDPConn
	var buf [1024]byte
	for {
		conn, err = net.DialUDP("udp", &lAddr, &rAddr)
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

		// listen

		_, _, err = conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("Could not Read from UDP: %s\n", err)
		}
		fmt.Printf("%s\n", buf)
		conn.Close()
	}
}
func main() {
	address := string(os.Args[1])
	localPort := string(os.Args[2])
	remotePort := string(os.Args[3])
	go communicate(address, remotePort, localPort)
	for {
	}
}
