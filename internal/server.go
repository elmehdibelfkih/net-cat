package internal

import "net"
import "sync"

type message struct {
	Time    string
	Name    string
	Message string
}

var Message message

type Server struct {
	Addr string
	// ReadTimeout     time.Duration
	// MaxMessageBytes int // o
	// ErrorLog        *log.Logger
	// nextProtoOnce   sync.Once
	// nextProtoErr    error
	mu            sync.Mutex
	listeners     net.Listener
	listenerGroup sync.WaitGroup
	conn          net.Conn
}

func NewServer(addr string) *Server {
	server := Server{Addr: addr}
	return &server

}

func (s *Server) StartServer() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.listeners = ln
	s.acceptNewConnection()
	return nil
}

func (s *Server) acceptNewConnection() error {
	for {
		conn, err := s.listeners.Accept()
		s.conn = conn
		if err != nil {
			return err
		}
		s.listenerGroup.Add(1)
		go s.handleConnection()
	}
}

func (s *Server) handleConnection() {
	var buf = make([]byte, 2048)
	s.conn.Read(buf)
	println(string(buf))
	defer s.conn.Close()
}
