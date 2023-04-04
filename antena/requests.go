package antena

type Request struct {
	Text string `json:"text"`
}

func (p *Request) SetPing() {
	p.Text = "ping"
}
