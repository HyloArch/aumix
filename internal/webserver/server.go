package webserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"slices"
	"sync"
	"time"
)

type contextKey struct {
	key string
}

var (
	ServerContextKey = &contextKey{"web-server"}

	webLogger = log.New(os.Stdout, "web-server", log.Ltime)
)

type Server struct {
	output chan Message

	httpServer *http.Server
	serveMux   *http.ServeMux
	mu         sync.Mutex
	sockets    []*Socket
}

func NewServer(port int) *Server {
	server := &Server{}

	serveMux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: serveMux,
		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(context.Background(), ServerContextKey, server)
		},
	}

	server.sockets = make([]*Socket, 0, 4)

	server.httpServer = httpServer
	server.serveMux = serveMux
	return server
}

func (s *Server) InitRoutes() {
	s.serveMux.HandleFunc("/{$}", index)
	s.serveMux.HandleFunc("/static/", static)
	s.serveMux.HandleFunc("/ws", ws)
}

func (s *Server) ListenAndServe(output chan Message) {
	webLogger.Println("Server started")
	s.output = output
	err := s.httpServer.ListenAndServe()
	log.Printf("Server stopped: %v\n", err)
}

func (s *Server) addSocket(socket *Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sockets = append(s.sockets, socket)
}

func (s *Server) removeSocket(socket *Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sockets = slices.DeleteFunc(s.sockets, func(soc *Socket) bool { return soc == socket })
}

func (s *Server) Broadcast(message Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, socket := range s.sockets {
		err := socket.Send(message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) BroadcastExcept(message Message, omit *Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, socket := range s.sockets {
		if socket != omit {
			socket.Send(message)
		}
	}
}

func (s *Server) Shutdown() {
	func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		for _, socket := range s.sockets {
			socket.Shutdown()
		}
		clear(s.sockets)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.httpServer.Shutdown(ctx)
	webLogger.Println("Server shutdown")
}
