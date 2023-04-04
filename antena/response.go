package antena

type Response struct {
	Text string `json:"text"`
}

func (p *Response) setPong() {
	p.Text = "pong"
}
