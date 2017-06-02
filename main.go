package main

import (
  "os"
  "github.com/fatih/color"
  "bufio"
  "net"
  "fmt"
  "github.com/alepacheco/ServerTCP"
)
var server *ServerTCP.Server

func main() {
	server = ServerTCP.NewServer()
	go server.Listen(":6666", listenFunc)
	go server.Read(readFun)
	for {
		server.Broadcast(getInput())
	}

}
func readFun(message ServerTCP.ClientMessage) {
	fmt.Println("'", message.Data, "'")
}
func listenFunc(conn net.Conn) {
	fmt.Println("New connection: ", conn.RemoteAddr().String())
}
/*
func listen() {
	newUser := color.New(color.Bold, color.FgGreen).FprintlnFunc()
	
		fmt.Println("")
		newUser(os.Stdout, "New connection: ", conn.RemoteAddr().String())
		fmt.Println("")

	for {

	}
}
*/
func getInput() string {
  color.Set(color.FgYellow)
  defer color.Unset()
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Command --> ")
  text, _ := reader.ReadString('\n')
  return text
}
