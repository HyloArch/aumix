package webserver

type messageOp string

var (
	MessageOpGET     messageOp = "GET"
	MessageOpSET     messageOp = "SET"
	MessageOpGET_OSC messageOp = "GET_OSC"
	MessageOpSET_OSC messageOp = "SET_OSC"
)

type Message struct {
	Sender *Socket   `json:"-"`
	Op     messageOp `json:"op"`
	Key    string    `json:"key"`
	Value  any       `json:"value"`
}
