package main

import (
	"fmt"
	"strings"

	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/net/tls"
)

type HttpClient struct {
	server string
	conn   net.Conn
	buf    [256]byte
}

func (client *HttpClient) request(path string) (result []byte, err error) {

	if client.conn != nil {
		client.conn.Close()
	}

	println("\r\n---------------\r\nDialing TCP connection")
	client.conn, err = tls.Dial("tcp", client.server, nil)
	for ; err != nil; client.conn, err = tls.Dial("tcp", client.server, nil) {
		return nil, err
	}
	println("Connected!\r")

	print("Sending HTTPS request...")
	fmt.Fprintln(client.conn, "GET ", path, " HTTP/1.1")
	fmt.Fprintln(client.conn, "Host:", strings.Split(client.server, ":")[0])
	fmt.Fprintln(client.conn, "User-Agent: TinyGo")
	fmt.Fprintln(client.conn, "Connection: close")
	fmt.Fprintln(client.conn)
	println("Sent!\r\n\r")

	response := []byte{}
	for n, err := client.conn.Read(client.buf[:]); n > 0; n, err = client.conn.Read(client.buf[:]) {
		if err != nil {
			return nil, err
		} else {
			response = append(response, client.buf[0:n]...)
		}
	}

	return response, nil

}
