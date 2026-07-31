package osc

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	oscLogger = log.New(os.Stdout, "osc-client", log.Ltime)
)

type Client struct {
	Ip      string
	Port    int
	Running bool

	address    *net.UDPAddr
	connection *net.UDPConn
}

func (c *Client) connectTo(ip string, port int) error {
	c.Ip = ip
	c.Port = port

	addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", c.Ip, c.Port))
	if err != nil {
		return err
	}
	c.address = addr

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return err
	}
	c.connection = conn
	return nil
}

func (c *Client) ConnectTo(ip string, port int) error {
	if c.connection != nil {
		c.connection.Close()
	}

	err := c.connectTo(ip, port)
	if err != nil {
		oscLogger.Printf("Error connecting to mixer: %v\n", err)
		oscLogger.Println("Falling back to address 192.168.1.1:10023")
		err = c.connectTo("192.168.1.1", 10023)
		if err != nil {
			oscLogger.Fatalf("Error connecting to mixer: %v", err)
		}
	}
	return nil
}

func (c *Client) Start(out chan Message, signal chan bool) {
	<-signal
	oscLogger.Println("Client started")
	c.Running = true
	buffer := make([]byte, 256)
	for {
		n, _, err := c.connection.ReadFromUDP(buffer)
		if err != nil {
			oscLogger.Println("Client stopped")
			c.Running = false
			out <- Message{
				Address: "stopped",
			}
			<-signal
			oscLogger.Println("Client resumed")
			c.Running = true
			continue
		}
		message := Decode(buffer[:n])
		out <- message
		oscLogger.Printf("Received message: %v\n", message)
	}
}

func (c *Client) Send(msg Message) error {
	buffer := make([]byte, 256)
	size := msg.Encode(buffer)
	_, err := c.connection.Write(buffer[:size])
	oscLogger.Printf("Sent message: %v\n", msg)
	return err
}

func (c *Client) Close() {
	c.connection.Close()
	oscLogger.Println("Client shutdown")
}
