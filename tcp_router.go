package tcprouter

import (
	"bufio"
	"fmt"
	"net"
)

//TCPRouter router is the struct
type TCPRouter struct {
	listen net.Listener
	routes []string
	conns  []net.Conn
}

//StartServer starts the server
func (t *TCPRouter) StartServer(typ string, host string) {

	var err error
	t.listen, err = net.Listen(typ, host)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		t.listen.Close()
		fmt.Println("Listener closed")
	}()

	for {

		conn, err := t.listen.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Println(msg)

	}

}

//AddRoute adds a route to the server
func (t *TCPRouter) AddRoute(route string, f func()) {
	t.routes = append(t.routes, route)
}
