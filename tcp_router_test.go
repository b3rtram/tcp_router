package tcprouter

import (
	"fmt"
	"net"
	"testing"
)

func TestStartServer(t *testing.T) {
	tr := TCPRouter{}
	tr.StartServer("tcp", "localhost:3678", '\n')
	c := tr.AddRoute("CONNNECT")

	go func() {
		msg := <-c
		fmt.Println(msg)
	}()

	conn, err := net.Dial("tcp", ":3678")

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	conn.Write([]byte("SUBSCRIBE test\n"))

}
