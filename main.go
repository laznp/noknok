package main

import (
	"fmt"
	"flag"
	"net"
)

func main() {
	host := flag.String("h","localhost","Hostname")
	port := flag.String("p","80","Port")
	flag.Parse()

	conn, err := net.Dial("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Printf("Can't connect to %s port %s\n", *host, *port)
	} else {
		fmt.Printf("Connected to %s port %s\n", *host, *port)
		defer conn.Close()
	}
}
