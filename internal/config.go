package internal

import (
	"log"
	"os"
)

var LINUX_LOGO []byte
var LOGS_FILE *os.File

const LOGO_FILE_PATH = "./data/linux.logo"
const FULL_SERVER_ERROR = "Oops! The server is currently full.\nKindly try again later.\n"
const EMPTY_MESSAGE_ERROR = "Oops! It looks like you left the input field empty.\nplease enter some data.\n"
const TOO_LONG_INPUT = "Input exceeds the maxium length allowed.\n"
const LOGS_FILE_PATH = "./data/logs.txt"
const INVALID_NAME = "Invalid name, Please use only letters, spaces, hyphens, and apostrophes\n"
const TOO_LONG_NAME = "Name is too long.\nMaximum allowed characters: 25\n"
const DUPLICATE_NAME = "Duplicate name detected.\nThis client name already exists.\n"
const MAX_NAME_LEN = 25

func SetUp() string {
	var addr string

	if len(os.Args) == 1 {
		addr = "8989"
	} else if len(os.Args) == 2 {
		addr = os.Args[1]
	} else {
		log.Fatal("[USAGE]: ./TCPChat $port")
	}
	tmp, err := os.ReadFile(LOGO_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	LINUX_LOGO = tmp
	logs, err := os.Create(LOGS_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	LOGS_FILE = logs
	return addr
}
