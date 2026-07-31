package x32

import (
	"aumix/internal/osc"
	"aumix/internal/util"
	"aumix/internal/webserver"
	"aumix/internal/x32/state"
	"log"
	"regexp"
	"sync"
	"time"
)

var xremote = osc.Message{Address: "/xremote"}

type queueElement struct {
	sender *webserver.Socket
	key    string
}

type oscHandler func(manager *Manager, message osc.Message, replyFunc func(webserver.Message) error)
type oscHandlerElement struct {
	matcher *regexp.Regexp
	handler oscHandler
}

type webHandler func(manager *Manager, message webserver.Message)

type mixerAddressMessage struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
type x32StatusMessage struct {
	State      any `json:"state"`
	IpAddress  any `json:"ipAddress"`
	ServerName any `json:"serverName"`
}

type Manager struct {
	Threads   int
	waitGroup sync.WaitGroup

	ConfigWrapper *state.ConfigWrapper

	OscClient         *osc.Client
	oscOutputChannel  chan osc.Message
	defaultOscHandler oscHandler
	oscHanders        []oscHandlerElement
	oscSignal         chan bool
	oscConnectStatus  int

	Webserver   *webserver.Server
	webChannel  chan webserver.Message
	webHandlers map[string]webHandler

	xremoteTicker   time.Ticker
	shutdownChannel chan struct{}

	ReceiveQueue *util.Queue[queueElement]
}

func NewManager(threads int, configWrapper *state.ConfigWrapper, oscClient *osc.Client, webserver *webserver.Server) *Manager {
	return &Manager{
		Threads:          threads,
		ConfigWrapper:    configWrapper,
		OscClient:        oscClient,
		oscHanders:       make([]oscHandlerElement, 0),
		oscConnectStatus: 0,
		Webserver:        webserver,
		webHandlers:      make(map[string]webHandler),
		ReceiveQueue:     util.NewQueueWithCapacity[queueElement](4),
	}
}

func (m *Manager) StartServices() {
	m.oscOutputChannel = make(chan osc.Message, 1)
	m.oscSignal = make(chan bool)
	go m.OscClient.Start(m.oscOutputChannel, m.oscSignal)

	m.webChannel = make(chan webserver.Message, 10)
	go m.Webserver.ListenAndServe(m.webChannel)
}

func (m *Manager) SetDefaultOscHandler(handler oscHandler) {
	m.defaultOscHandler = handler
}

func (m *Manager) RegisterOscHandler(matcher *regexp.Regexp, handler oscHandler) {
	m.oscHanders = append(m.oscHanders, oscHandlerElement{
		matcher: matcher,
		handler: handler,
	})
}

func (m *Manager) RegisterWebHandler(key string, handler webHandler) {
	m.webHandlers[key] = handler
}

func (m *Manager) handleOscMessage(msg osc.Message) {
	var replyFunc func(webserver.Message) error
	head, ok := m.ReceiveQueue.Peek()
	if ok && msg.Address == head.key {
		m.ReceiveQueue.Dequeue()
		replyFunc = head.sender.Send
	}

	for _, element := range m.oscHanders {
		found := element.matcher.MatchString(msg.Address)
		if found {
			element.handler(m, msg, replyFunc)
			return
		}
	}
	m.defaultOscHandler(m, msg, replyFunc)
}

func (m *Manager) handleWebMessage(msg webserver.Message) {
	handler, ok := m.webHandlers[msg.Key]
	if ok {
		handler(m, msg)
	}
}

func (m *Manager) runLoop() {
	for {
		select {
		case msg := <-m.oscOutputChannel:
			m.handleOscMessage(msg)
		case msg := <-m.webChannel:
			m.handleWebMessage(msg)
		case <-m.xremoteTicker.C:
			m.refreshX32Connection()
		case <-m.shutdownChannel:
			return
		}
	}
}

func (m *Manager) Run() {
	m.xremoteTicker = *time.NewTicker(9 * time.Second)
	m.shutdownChannel = make(chan struct{})

	for range m.Threads {
		m.waitGroup.Go(m.runLoop)
	}

	m.refreshX32Connection()
}

func (m *Manager) refreshX32Connection() {
	switch m.oscConnectStatus {
	case 1:
		m.Webserver.Broadcast(webserver.Message{
			Op:    webserver.MessageOpSET,
			Key:   "status",
			Value: false,
		})
		m.oscConnectStatus--
		if !m.OscClient.Running {
			m.oscSignal <- true
		}
	case 0:
		m.ReceiveQueue.Clear()
		if !m.OscClient.Running {
			m.oscSignal <- true
		}
	default:
		m.oscConnectStatus--
	}

	err := m.OscClient.Send(xremote)
	if err != nil {
		log.Printf("Failed to send data: %v\n", err)
	}
	err = m.OscClient.Send(osc.Message{
		Address: "/status",
	})
	if err != nil {
		log.Printf("Failed to send data: %v\n", err)
	}
}

func (m *Manager) IsOSCConnected() bool {
	return m.oscConnectStatus > 0
}

func (m *Manager) Shutdown() {
	close(m.shutdownChannel)
	m.waitGroup.Wait()
}
