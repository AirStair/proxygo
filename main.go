package main

import (
	"io"
	"log"
	"net"
	"fmt"
	"bufio"
	"net/http"
	"strings"
)

func readRequest(raw, scheme string) (*http.Request, error) {
	r, err := http.ReadRequest(bufio.NewReader(strings.NewReader(raw)))
	if err != nil { return nil, err }
	r.RequestURI, r.URL.Scheme, r.URL.Host = "", scheme, r.Host
	return r, nil
 }

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			buf, err := io.ReadAll(c)
			if err != nil {
				log.Fatal(err)
			}
			myString := string(buf[:])
			//log.Print(myString)
			req, err := readRequest(myString, "http")
			if err != nil {
				log.Fatal(err)
			}
			conn, err := net.Dial("tcp", req.URL.Host)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(conn, myString)
			respBuf, err := io.ReadAll(conn)
			if err != nil {
				log.Fatal(err)
			}
			myRespString := string(respBuf[:])

			//c.Write(status)
			log.Print(myRespString)

			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
