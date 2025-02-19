package internal

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type Server struct {
	Addr               string
	clientCounter      int
	clientCounterMutex sync.Mutex
	listener           net.Listener
	clients            map[net.Conn]*Client
	nextProtoErr       error
}

func NewServer(addr string) *Server {
	server := Server{Addr: addr, clients: make(map[net.Conn]*Client, MAX_CLIENT)}
	return &server
}

func (s *Server) StartServer() error {
	ln, err := net.Listen("tcp", ":"+s.Addr)
	if err != nil {
		return err
	}
	fmt.Println("Listening on the port :", s.Addr)
	defer ln.Close()
	s.listener = ln
	return s.acceptNewConnection()
}

func (s *Server) acceptNewConnection() error {
	for {
		conn, err := s.listener.Accept()
		if s.clientCounter < MAX_CLIENT {
			s.clientCounterMutex.Lock()
			s.clientCounter++
			s.clientCounterMutex.Unlock()
			s.clients[conn] = &Client{s: s, conn: conn}
			if err != nil {
				return err
			}
			go s.handleConnection(conn)
		} else if err == nil {
			conn.Write([]byte(FULL_SERVER_ERROR))
			conn.Close()
		}
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	// TODO: add a go routine that listen in the stdin for the exit command then execute Shutdown() function
	defer conn.Close()
	buf := make([]byte, 2048)
	conn.Write(LINUX_LOGO)
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	for {
		byteReaded, err := conn.Read(buf)
		s.nextProtoErr = err
		if err != nil {
			s.clients[conn].endConnection()
			return
		} else {
			err := MessageValid(buf[:byteReaded])
			if err != nil {
				s.clients[conn].sendMessage(fmt.Sprintf("%v\n", err))
			}
			s.clients[conn].Message += string(buf[:byteReaded])
			if strings.Contains(s.clients[conn].Message, "\n") {
				s.clients[conn].isReady = true
			}
		}
		if s.clients[conn].isReady {
			s.clients[conn].handleInput()
			if s.clients[conn].Name != "" {
				s.clients[conn].sendMessage(s.clients[conn].getPrefix()) // display to the client his input
			}
			s.clients[conn].Message = ""    //free the message field
			s.clients[conn].isReady = false // wait the next \n to be send from the client
		}
	}
}

func (s *Server) broadcastMessage(conn net.Conn) {
	if len(s.clients[conn].Message) > MAX_LENGTH_MESSGE {
		s.clients[conn].sendMessage(TOO_LONG_INPUT)
		return
	}
	if !s.clients[conn].noPrefix {
		message := s.clients[conn].FormatedMessage()
		fmt.Fprint(LOGS_FILE, message)
	}
	for cn, cl := range s.clients {
		if cn != conn {
			if cl.Authentication {
				if s.clients[conn].noPrefix {
					cl.sendMessage("\n" + s.clients[conn].Message)
					cl.sendMessage(cl.getPrefix())
				} else {
					cl.sendMessage("\n" + s.clients[conn].FormatedMessage())
					cl.sendMessage(cl.getPrefix())
				}
			}
		}
	}
}

func (s *Server) Shutdown() {
	for conn := range s.clients {
		conn.Close()
	}
	s.listener.Close()
}
