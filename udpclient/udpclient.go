package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type UDPRequest struct {
	One   string
	Two   int
	Three string
}

func main() {
	fmt.Println("Starting client")

	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8084")
	if err != nil {
		fmt.Println(err)
		return
	}

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	i := 0
	for {
		msg := &UDPRequest{}
		msg.One = strconv.Itoa(i)
		msg.Two = i
		msg.Three = "some string"

		i++

		b, err := msgpack.Marshal(msg)
		if err != nil {
			fmt.Println(err)
		}

		_, err = conn.Write(b)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second * 1)
	}
}
