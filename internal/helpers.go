package internal

import (
	"fmt"
	"os"
)

func MessageValid(str []byte) error {
	for _, v := range str {
		if v <= 31 && v != '\n' && v != '\r' {
			return fmt.Errorf("you cannot enter control characters")
		}
	}
	return nil
}

func SetUp() string {
	var addr string

	if len(os.Args) == 1 {
		addr = "8989"
	} else if len(os.Args) == 2 {
		addr = os.Args[1]
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(0)
	}
	tmp, err := os.ReadFile(LOGO_FILE_PATH)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	LINUX_LOGO = tmp
	logs, err := os.Create(LOGS_FILE_PATH)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	LOGS_FILE = logs
	return addr
}
