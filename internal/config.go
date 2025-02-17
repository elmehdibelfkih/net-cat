package internal

import (
	"fmt"
	"os"
	"strconv"
)

var LINUX_LOGO []byte
var LOGS_FILE *os.File

const MAX_CLIENT = 10
const MAX_NAME_LEN = 25
const MAX_LENGTH_MESSGE = 256
const LOGO_FILE_PATH = "./data/linux.logo"
const FULL_SERVER_ERROR = "Oops! The server is currently full.\nKindly try again later.\n"
const EMPTY_MESSAGE_ERROR = "Oops! It looks like you left the input field empty.\nplease enter some data.\n"
const LOGS_FILE_PATH = "./data/logs.txt"
const INVALID_NAME = "Invalid name, Please use only letters, spaces, hyphens, and apostrophes\n"
const TOO_LONG_NAME = "Name is too long.\nMaximum allowed characters: 25\n"
const DUPLICATE_NAME = "Duplicate name detected.\nThis client name already exists.\n"

var TOO_LONG_INPUT = fmt.Sprintf("Error: The input message exceeds the maximum length of %d characters. Please shorten your message.\n", MAX_LENGTH_MESSGE)

func SetUp() string {
	var addr string

	if len(os.Args) == 1 {
		addr = "8989"
	} else if len(os.Args) == 2 {
		port, err := strconv.Atoi(os.Args[1])
		if port == 0 || err != nil {
			fmt.Println("you cannot use that port is either invalid or reserved by IANA!") // internet assigned numbers authourity
			os.Exit(0)
		}
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
