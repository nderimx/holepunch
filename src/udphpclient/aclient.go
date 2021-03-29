package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
)

var (
	serverAddress = "localhost"
	serverPort    = "9090"
	localPort         = "7171"
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

func communicate(remoteAddress, rPort string) {
	intrPort, err := strconv.Atoi(rPort)
	if err != nil {
		fmt.Printf("Could not convert input port to integer: %s\n", err)
	}
	intlPort, err := strconv.Atoi(localPort)
	if err != nil {
		fmt.Printf("Could not convert input port to integer: %s\n", err)
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

	// punch: could send a secure identification token istead of the work punch
	jtext, err := json.Marshal("punch")
	if err != nil {
		fmt.Printf("Could not parse text to json: %s\n", err)
	}
	conn.Write(jtext)

	defer conn.Close()
	go send(conn)
	listen(conn)
}

func register(ID string) {
	remoteAddress := serverAddress
	rPort := serverPort
	intrPort, err := strconv.Atoi(rPort)
	if err != nil {
		fmt.Printf("Could not convert imput port to integer: %s\n", err)
	}
	intlPort, err := strconv.Atoi(localPort)
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

	fmt.Print("Requesting Registration")
	text := "r" + ID
	jtext, err := json.Marshal(text)
	if err != nil {
		fmt.Printf("Could not parse text to json: %s\n", err)
	}
	conn.Write(jtext)

	var buf [1024]byte
	_, _, err = conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Printf("Could not Read from UDP: %s\n", err)
	}
	fmt.Printf("\nrec: %s\nsend: ", buf)
}

func updateRegistration(ID string) {
	remoteAddress := serverAddress
	rPort := serverPort
	intrPort, err := strconv.Atoi(rPort)
	if err != nil {
		fmt.Printf("Could not convert imput port to integer: %s\n", err)
	}
	intlPort, err := strconv.Atoi(localPort)
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

	fmt.Print("Requesting Registration Update")
	text := "u" + ID
	jtext, err := json.Marshal(text)
	if err != nil {
		fmt.Printf("Could not parse text to json: %s\n", err)
	}
	conn.Write(jtext)

	var buf [1024]byte
	_, _, err = conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Printf("Could not Read from UDP: %s\n", err)
	}
	fmt.Printf("\nrec: %s\nsend: ", buf)
}

func getConnection(ID string) (string, string, error) {
	remoteAddress := serverAddress
	rPort := serverPort
	intrPort, err := strconv.Atoi(rPort)
	if err != nil {
		fmt.Printf("Could not convert input port to integer: %s\n", err)
	}
	intlPort, err := strconv.Atoi(localPort)
	if err != nil {
		fmt.Printf("Could not convert input port to integer: %s\n", err)
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

	fmt.Print("Requesting Peer ID")
	text := "c" + ID
	jtext, err := json.Marshal(text)
	if err != nil {
		fmt.Printf("Could not parse text to json: %s\n", err)
	}
	conn.Write(jtext)

	var buf [1024]byte
	_, _, err = conn.ReadFromUDP(buf[:])
	if err != nil {
		fmt.Printf("Could not Read from UDP: %s\n", err)
	}
	fmt.Printf("\nrec: %s\nsend: ", buf)

	response := string(bytes.Trim(buf[:], "\x00"))
	if strings.Contains(response, "fail") {
		return "", "", fmt.Errorf("Could not get Connection to peer: %s", response)
	}
	peerEndpoint := response
	tempSlice := strings.Split(peerEndpoint, ":")
	fmt.Println(response)
	return tempSlice[0], tempSlice[1], nil
}

func main() {

	config, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Could not read configuration file: %s\n", err)
		return
	}
	var ipInfo map[string]string
	err = json.Unmarshal(config, &ipInfo)
	if err != nil {
		fmt.Printf("Could not parse the json configuration file: %s\n", err)
		return
	}
	serverAddress = ipInfo["serverAddress"]
	serverPort = ipInfo["serverPort"]
	localPort = ipInfo["localPort"]

	register(os.Args[1])
	var peerAddress string
	var peerPort string
	// var err error
	for {
		peerAddress, peerPort, err = getConnection(os.Args[2])
		if err == nil {
			break
		}
	}
	fmt.Println(peerAddress + " " + peerPort)
	communicate(peerAddress, peerPort)
}
