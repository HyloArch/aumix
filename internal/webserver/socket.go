package webserver

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/coder/websocket"
)

type Socket struct {
	server     *Server
	connection *websocket.Conn
	context    context.Context
	cancel     context.CancelFunc
}

func NewSocket(w http.ResponseWriter, req *http.Request) (*Socket, error) {
	socket := &Socket{}

	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{})
	if err != nil {
		return socket, err
	}

	ctx, cancel := context.WithCancel(req.Context())

	socket.connection = conn
	socket.context = ctx
	socket.cancel = cancel

	server, ok := ctx.Value(ServerContextKey).(*Server)
	if !ok {
		return socket, errors.New("Unable to get server for websocket")
	}
	server.addSocket(socket)
	socket.server = server

	return socket, nil
}

func (s *Socket) Reader() (websocket.MessageType, io.Reader, error) {
	return s.connection.Reader(s.context)
}

func (s *Socket) Send(msg Message) error {
	writer, err := s.connection.Writer(s.context, websocket.MessageText)
	if err != nil {
		return err
	}
	defer writer.Close()

	json.NewEncoder(writer).Encode(msg)
	// webLogger.Printf("Sent message: %v", msg)
	return nil
}

func (s *Socket) Write(messageType websocket.MessageType, data []byte) error {
	return s.connection.Write(s.context, messageType, data)
}

func (s *Socket) Listen() error {
	for {
		_, reader, err := s.Reader()
		if err != nil {
			return err
		}

		var msg Message
		json.NewDecoder(reader).Decode(&msg)
		msg.Sender = s
		s.server.output <- msg
		// webLogger.Printf("Received message: %v\n", msg)
	}
}

func (s *Socket) removeFromServer() {
	if s.server == nil {
		return
	}
	s.server.removeSocket(s)
	s.server = nil
}

func (s *Socket) Close(code websocket.StatusCode, reason string) {
	s.connection.Close(code, reason)
	s.removeFromServer()
}

func (s *Socket) CloseNow() {
	s.connection.CloseNow()
	s.removeFromServer()
}

func (s *Socket) Shutdown() {
	s.connection.Close(websocket.StatusGoingAway, "Server shutting down")
	s.cancel()
}

func ws(w http.ResponseWriter, req *http.Request) {
	socket, err := NewSocket(w, req)
	if err != nil {
		webLogger.Printf("Handshake failed: %v\n", err)
		return
	}
	defer socket.CloseNow()

	webLogger.Printf("Client connected: %s\n", req.RemoteAddr)

	err = socket.Listen()

	webLogger.Printf("Socket error: %v\n", err)

	webLogger.Printf("Client disconnected: %s\n", req.RemoteAddr)
}
