package main

import (
	"fmt"
	"net-cat/internal"
	"os"
)

func main() {
	server := internal.NewServer(internal.SetUp())
	err := server.StartServer()
	internal.LOGS_FILE.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

}
