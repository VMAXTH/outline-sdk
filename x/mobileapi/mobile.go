package mobileapi

type Proxy struct {
	running bool
}

func NewProxy() *Proxy { return &Proxy{} }

func (p *Proxy) Start(endpoint string) (bool, string) {
	p.running = true
	return p.running, "started:" + endpoint
}

func (p *Proxy) Stop() (bool, string) {
	p.running = false
	return !p.running, "stopped"
}

func (p *Proxy) IsRunning() bool { return p.running }

func Version() string { return "0.1.0-mobileapi" }
