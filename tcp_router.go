package tcprouter

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//TCPRouter router is the struct
type TCPRouter struct {
	listen net.Listener
	routes map[string]func(net.Conn, string)
	conns  []net.Conn
}

//StartServer starts the server
func (t *TCPRouter) StartServer(typ string, host string, delimiter byte) {

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

	go func() {
		for {

			conn, err := t.listen.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go t.startRead(conn, delimiter)
		}
	}()

}

func (t *TCPRouter) startRead(c net.Conn, delimiter byte) {

	reader := bufio.NewReader(c)

	for {
		msg, _ := reader.ReadString(delimiter)

		sa := strings.SplitN(msg, " ", 2)

		val, ok := t.routes[sa[0]]

		if ok {
			val(c, msg)
		}

	}
}

//AddRoute adds a route to the server
func (t *TCPRouter) AddRoute(route string, f func(net.Conn, string)) {
	t.routes[route] = f
}
