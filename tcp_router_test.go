package tcprouter

import (
	"fmt"
	"net"
	"testing"
)

func TestStartServer(t *testing.T) {
	tr := TCPRouter{}
	fmt.Println("Start Test")
	tr.StartServer("tcp", "localhost:3678", '\n')
	c := tr.AddRoute("CONNNECT")

	conn, err := net.Dial("tcp", "localhost:3678")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Dialed")

	defer conn.Close()

	msg := <-c
	fmt.Println(msg)

}
