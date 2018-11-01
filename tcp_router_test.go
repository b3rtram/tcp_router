package tcprouter

import (
	"fmt"
	"net"
	"testing"
)

func TestStartServer(t *testing.T) {
	tr := TCPRouter{}
	fmt.Println("Start Test")
	tr.StartServer("tcp", "localhost:3001", '\n')
	c := tr.AddRoute("CONNECT")

	conn, err := net.Dial("tcp", "localhost:3001")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Dialed")
	conn.Write([]byte("CONNECT test\n"))

	defer conn.Close()

	msg := <-c
	fmt.Println(msg)

	defer func() {
		tr.StopServer()
	}()

}
