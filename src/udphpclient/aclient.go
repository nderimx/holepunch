package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func listen(conn *net.UDPConn) {
	for {
		var buf [1024]byte
		_, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("Could not Read from UDP: %s\n", err)
		}
		fmt.Printf("\nrec: %s\nsend: ", buf)
	}
}
func send(conn *net.UDPConn) {
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
	conn, err := net.DialUDP("udp", &lAddr, &rAddr)
	if err != nil {
		fmt.Printf("Could not Connect to remote Address: %s\n", err)
	}
	defer conn.Close()
	go send(conn)
	listen(conn)
}

func register() {

}

func updateRegistration() {

}

func getConnection() {

}

func connectp2p() {

}

func main() {
	communicate(os.Args[1], os.Args[2], os.Args[3])
}
