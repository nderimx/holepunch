package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"errors"
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
	
}

func spawnNewConnection() {
	conn, err = net.DialUDP("udp", &lAddr, rAddr)
	if err != nil {
		fmt.Printf("Could not Connect to remote Address: %s\n", err)
	}
	defer conn.Close()
	go listen(conn)
	send(conn)
}

func listenForConnections() {
	for {
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
		err=conn.Close()
		if err!=nil {
			fmt.Printf("Could not Close first contact connection: %s\n", err)
		}
	}
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

func parseRequest(buf [1024]byte) (string, string, error) {


	return "", "", errors.New("Could not parse received input")
}

func register(ID, Endpoint string) {
		clients=append(clients, client{ID, Endpoint})
}

func updateRegistration(ID, Endpoint string) error {
	for _, client:=range clients {
		if ID==client.ID {
			client.Endpoint=Endpoint
			return nil
		}
	}
	return errors.New("Could not find the given ID in the database")
}

func getConnection(ID string) (string, error) {
	for _, client:=range clients {
		if ID==client.ID {
			return client.Endpoint, nil
		}
	}
	return "", errors.New("Could not find the given ID in the database")
}

func main() {
	serveCommunicate()
}
