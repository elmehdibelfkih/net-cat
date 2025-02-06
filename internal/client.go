package internal

import (
	"io"
	"net"
	"os"
	"strings"
	"time"
)

type Client struct {
	noPrefix       bool
	Authentication bool
	Name           string
	Message        string
	conn           net.Conn
	s              *Server
	isReady        bool
}

func (cl *Client) handleInput() {
	if cl.Name == "" {
		if cl.Message == "\n" {
			cl.sendMessage(EMPTY_MESSAGE_ERROR)
			cl.conn.Write([]byte("[ENTER YOUR NAME]: "))
		} else {
			cl.Name = cl.Message[:len(cl.Message)-1]
			if !cl.isValideName(cl.Name) {
				cl.conn.Write([]byte("[ENTER YOUR NAME]: "))
				return
			}
			cl.Message = ""
			cl.Authentication = true
			cl.noPrefix = true
			cl.Message = cl.Name + " has joinned our chat...\n"
			cl.s.broadcastMessage(cl.conn)
			cl.noPrefix = false
			data, err := os.ReadFile(LOGS_FILE_PATH)
			if err == nil {
				cl.sendMessage(string(data))
			}
		}
	} else if cl.Message != "\n" {
		cl.s.broadcastMessage(cl.conn)

	} else {
		cl.sendMessage(EMPTY_MESSAGE_ERROR)
	}
}

func (cl *Client) sendMessage(message string) {
	cl.conn.Write([]byte(message))
}

func (cl *Client) endConnection() {
	cl.s.clientCounterMutex.Lock()
	cl.s.clientCounter--
	cl.s.clientCounterMutex.Unlock()
	if cl.s.nextProtoErr == io.EOF {
		cl.Message = cl.Name + " has left our chat...\n"
	} else {
		cl.Message = cl.Name + "error connection\n"
	}
	cl.noPrefix = true
	cl.s.broadcastMessage(cl.conn)
	cl.noPrefix = false
	delete(cl.s.clients, cl.conn)
}

func (cl *Client) getPrefix() string {
	currentTime := time.Now()
	return "[" + currentTime.Format("2006-01-02 15:04:05") + "][" + cl.Name + "]:"
}

func (cl *Client) FormatedMessage() string {
	if cl.Name == "" {
		return cl.Message
	}
	return cl.getPrefix() + cl.Message
}

func (cl *Client) isValideName(name string) bool {
	if len(name) > MAX_NAME_LEN {
		cl.sendMessage(TOO_LONG_NAME)
		cl.Name = ""
		return false
	}

	for _, c := range name {
		if !(strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", c) ||
			strings.ContainsRune(" -'", c)) {
			cl.sendMessage(INVALID_NAME)
			cl.Name = ""
			return false
		}
	}
	return cl.isUniqueName()
}

func (cl *Client) isUniqueName() bool {
	for cn, clt := range cl.s.clients {
		if cn != cl.conn && clt.Name == cl.Name {
			cl.sendMessage(DUPLICATE_NAME)
			cl.Name = ""
			return false
		}
	}
	return true
}
