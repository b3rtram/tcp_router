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
	routes map[string](chan Msg)
	conns  []net.Conn
}

//Msg sends over the generated channels
type Msg struct {
	conn net.Conn
	msg  string
}

//StartServer starts the server
func (t *TCPRouter) StartServer(typ string, host string, delimiter byte) {

	var err error
	t.routes = make(map[string](chan Msg))
	t.listen, err = net.Listen(typ, host)

	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			fmt.Println("start listening")
			conn, err := t.listen.Accept()
			fmt.Println("client connected")
			if err != nil {
				fmt.Println(err)
				continue
			}
			go t.startRead(conn, delimiter)
		}
	}()

}

func (t *TCPRouter) startRead(c net.Conn, delimiter byte) {

	fmt.Println("start reading")
	reader := bufio.NewReader(c)

	for {
		msg, _ := reader.ReadString(delimiter)
		fmt.Printf("MSG: %s", msg)
		sa := strings.SplitN(msg, " ", 2)
		val, ok := t.routes[sa[0]]
		if ok {
			val <- Msg{conn: c, msg: msg}
		}
	}
}

//AddRoute adds a route to the server
func (t *TCPRouter) AddRoute(route string) chan Msg {
	c := make(chan Msg)
	t.routes[route] = c

	return c
}

//StopServer to manually close the Server
func (t *TCPRouter) StopServer() {
	t.listen.Close()
}
