package tcp_router

import (
	"fmt"
	"net"
	"testing"
)

func TestStartServer(t *testing.T) {
	tr := TCPRouter{}
	fmt.Println("Start Test")
	tr.StartServer("tcp", "localhost:3001", '\n')

	c := make(chan bool)
	f := func(conn net.Conn, msg string) {
		fmt.Println(msg)
		c <- true
	}

	tr.AddRoute("CONNECT", f)

	conn, err := net.Dial("tcp", "localhost:3001")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Dialed")
	conn.Write([]byte("CONNECT test\n"))

	defer func() {
		conn.Close()
		tr.StopServer()
	}()

	b := <-c

	fmt.Println(b)

}
