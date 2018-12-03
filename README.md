# goTcpRouter
using socket connection as simple as http requests.
Is used with string commands and splits the command on whitespace. We dont need the overhead over http use the socket directly.

```go

	tr := TCPRouter{}
	fmt.Println("Start Test")
	
	//we need a delimiter to find out when a command ends
	tr.StartServer("tcp", "localhost:3001", '\n')
	
	//function called when command from client matches. Gets socket connection and complete command string 
	f := func(conn net.Conn, msg string) {
		fmt.Println(msg)
	}

	//Add a route with command name is CONNECT. f is a function as defined above 
	tr.AddRoute("CONNECT", f)

	defer func() {
		//stops the server. Has to called manually
		tr.StopServer()
	}()

```
