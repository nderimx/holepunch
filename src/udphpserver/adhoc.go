package main

import (
	"errors"
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

func main() {

	udpAddr := net.UDPAddr{
		Port: 9090,
		IP:   nil,
	}

	conn, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not listen for UDP datagrams: %s", err.Error())
		os.Exit(1)
	}

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {

	var buf [1024]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Printf("Could not read UDP datagrams: %s", err)
		return
	}
	endpoint := fmt.Sprintf("%s", addr)
	response, err := handleRequest(buf, endpoint)
	if err != nil {
		fmt.Printf("Could not handle request: %s", err)
	}

	conn.WriteToUDP([]byte(response), addr)
}

func handleRequest(buf [1024]byte, Endpoint string) (string, error) {
	reqType := string(buf[1])
	switch reqType {
	case "r":
		register(string(buf[2:]), Endpoint)
		return "successfuly registered.", nil
	case "u":
		err := updateRegistration(string(buf[2:]), Endpoint)
		if err != nil {
			return "updating ID failed", err
		}
		return "successfuly updated.", nil
	case "c":
		peerEndpoint, err := getConnection(string(buf[2:]))
		if err != nil {
			return "finding peer failed", err
		}
		return peerEndpoint, nil
	default:
		return "invalid character", fmt.Errorf("invalid code: %s", reqType)
	}
}

func register(ID, Endpoint string) {
	clients = append(clients, client{ID, Endpoint})
}

func updateRegistration(ID, Endpoint string) error {
	for _, client := range clients {
		if ID == client.ID {
			client.Endpoint = Endpoint
			return nil
		}
	}
	return errors.New("Could not find the given ID in the database")
}

func getConnection(ID string) (string, error) {
	for _, client := range clients {
		if ID == client.ID {
			return client.Endpoint, nil
		}
	}
	return "", errors.New("Could not find the given ID in the database")
}
