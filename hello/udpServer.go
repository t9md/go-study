package main

import (
	"fmt"
	"net"
	"time"
)

const port = "127.0.0.1:1200"

func main() {

	udpAddress, err := net.ResolveUDPAddr("udp4", port)

	if err != nil {
		fmt.Println("error resolving UDP address on ", port)
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddress)

	if err != nil {
		fmt.Println("error listening on UDP port ", port)
		fmt.Println(err)
		return
	}

	defer conn.Close()

	var buf []byte = make([]byte, 1500)

	for {
		time.Sleep(100 * time.Millisecond)
		n, address, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("error reading data from connection")
			fmt.Println(err)
			return
		}

		if address != nil {

			fmt.Println("got message from ", address, " with n = ", n)

			if n > 0 {
				fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
			}
		}
	}

}
