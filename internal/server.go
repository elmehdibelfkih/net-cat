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
	server := Server{Addr: addr, clients: make(map[net.Conn]*Client)}
	return &server
}

func (s *Server) StartServer() error {
	ln, err := net.Listen("tcp", ":"+s.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.listener = ln
	return s.acceptNewConnection()
}

func (s *Server) acceptNewConnection() error {
	for {
		conn, err := s.listener.Accept()
		if s.clientCounter < 10 {
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
	defer conn.Close()
	var buf = make([]byte, 2048)
	conn.Write(LINUX_LOGO)
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	for {
		byteReaded, err := conn.Read(buf)
		s.nextProtoErr = err
		if err != nil {
			s.clients[conn].endConnection()
			return
		} else { // FIXME: handel long input
			s.clients[conn].Message += string(buf[:byteReaded])
			if strings.Contains(s.clients[conn].Message, "\n") {
				s.clients[conn].isReady = true
			}
		}
		if s.clients[conn].isReady {
			s.clients[conn].handleInput()
			s.clients[conn].Message = ""
			s.clients[conn].isReady = false
		}
	}
}

func (s *Server) broadcastMessage(conn net.Conn) {
	if !s.clients[conn].noPrefix {
		fmt.Fprint(LOGS_FILE, s.clients[conn].FormatedMessage())
	}
	for cn, cl := range s.clients {
		if cn != conn {
			if cl.Authentication {
				if s.clients[conn].noPrefix {
					cl.sendMessage(s.clients[conn].Message)
				} else {
					cl.sendMessage(s.clients[conn].FormatedMessage())
				}
			}
		}
	}
}
