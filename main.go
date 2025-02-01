package main

import (
	"net-cat/internal"
	"os"
)

func main() {
	server := internal.NewServer(os.Args[1])
	server.StartServer()

}
