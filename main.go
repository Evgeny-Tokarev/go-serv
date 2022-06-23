package main

import (
	"fmt"

	"github.com/Evgeny-Tokarev/go-serv/client"
	"github.com/Evgeny-Tokarev/go-serv/server"
)

func main() {
	fmt.Println("package main")
	server.InitServer()
	client.InitClient()
}
