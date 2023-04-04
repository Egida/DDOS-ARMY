package antena

type Request struct {
	Text string `json:"text"`
}

func (p *Request) setPing() {
	p.Text = "ping"
}

type Response struct {
	Text string `json:"text"`
}

func (p *Response) setPong() {
	p.Text = "pong"
}
