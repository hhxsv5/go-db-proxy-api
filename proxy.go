package godpa

import (
	"github.com/hhxsv5/go-db-proxy-api/http"
	nhttp "net/http"
)

type Proxy struct {
	server *http.Server
}

func (p *Proxy) Run() {
	p.server.Run()
}

func NewProxy(handlers map[string]func(w nhttp.ResponseWriter, r *nhttp.Request)) *Proxy {
	s := http.NewServer(handlers)
	return &Proxy{s}
}
