package main

import (
	"fmt"
	"net"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type UDPRequest struct {
	One   string
	Two   int
	Three string
}

func main() {
	fmt.Println("Starting server")

	addr, err := net.ResolveUDPAddr("udp", ":8084")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}

		msg := &UDPRequest{}
		err = msgpack.Unmarshal(buf[0:n], msg)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Received packet from", remoteAddr)
		fmt.Printf("One: %s, Two: %d, Three: %s\n", msg.One, msg.Two, msg.Three)
	}
}
