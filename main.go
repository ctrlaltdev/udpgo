package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	dst  = flag.String("d", "::1", "Destination")
	host = flag.String("h", "::", "Host")
	port = flag.Int("p", 1337, "Port")

	listen = flag.Bool("l", false, "Start server")
	send   = flag.Bool("s", false, "Send")

	wg = &sync.WaitGroup{}
)

func Send(host *string, msg string) {
	Conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(*host), Port: *port, Zone: ""})
	defer Conn.Close()

	Conn.Write([]byte(msg))
}

func Listen(host *string, port *int) {
	defer wg.Done()

	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(*host), Port: *port, Zone: ""})
	defer ServerConn.Close()
	fmt.Println("Listening on ", *host, " on port ", *port)

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
		go Listen(host, port)
	}

	if *send {
		Send(dst, strings.Join(flag.Args(), " "))
	}

	wg.Wait()
}
