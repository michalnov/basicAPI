package main

import (
	"fmt"

	"github.com/michalnov/basicAPI/server"
)

func helloServer() {
	fmt.Println("Hello server")
}

func main() {
	var exit = make(chan int)
	//var serv server.Server
	serv, err := server.NewServer(":1201", exit)
	if err != nil {
		fmt.Println("Error while creating server")
		return
	} else {
		go runServer(serv)
		ex := <-exit
		fmt.Println(ex)
	}
}

func runServer(s server.Server) {
	s.Start()
}
