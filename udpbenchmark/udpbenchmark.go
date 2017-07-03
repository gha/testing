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
	fmt.Println("Starting benchmark")

	go runServer()

	err := runBenchmark()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func runServer() {
	fmt.Println(" - starting server")

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
	i := 0
	t := time.Now().Unix()
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}

		msg := &UDPRequest{}
		err = msgpack.Unmarshal(buf[0:n], msg)
		if err != nil {
			fmt.Println(err)
		}

		nt := time.Now().Unix()
		if nt != t {
			t = nt
			fmt.Println(i, "packets/second")
			i = 0
		}

		i++
	}
}

func runBenchmark() error {
	time.Sleep(time.Second * 1)

	fmt.Println(" - starting benchmark")

	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8084")
	if err != nil {
		return err
	}

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		return err
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
	}
}
