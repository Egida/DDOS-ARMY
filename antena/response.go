package antena

// RESPONSE TYPE
const (
	HELLO = "HELLO"
	ORDER = "ORDER"
)

// ORDER TYPE
const (
	START   = "START"
	STOP    = "STOP"
	NOTHING = "NOTHING"
)

type Response struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func (p *Response) setPong() {
	p.Text = "pong"
	p.Type = HELLO
}

func (p *Response) setOrder(order string) {
	p.Text = order
	p.Type = ORDER
}
