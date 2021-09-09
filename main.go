package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	listen = flag.Bool("l", false, "Start server")
	send   = flag.Bool("s", false, "Send")
	wg     = &sync.WaitGroup{}
)

func Send(msg string) {
	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 10001, Zone: ""})
	defer Conn.Close()

	Conn.Write([]byte(msg))
}

func Listen() {
	defer wg.Done()

	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 10001, Zone: ""})
	defer ServerConn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, _ := ServerConn.ReadFromUDP(buf)

		payload := string(buf[0:n])
		fmt.Println("Received from ", addr)
		fmt.Println(payload)
	}
}

func main() {
	flag.Parse()

	if *listen {
		wg.Add(1)
		go Listen()
	}

	if *send {
		Send(strings.Join(flag.Args(), " "))
	}

	wg.Wait()
}
