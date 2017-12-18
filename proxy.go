package godpa

import (
	"github.com/hhxsv5/go-db-proxy-api/http"
	nhttp "net/http"
	"time"
	"io"
)

type Proxy struct {
	server *http.Server
}

func (p *Proxy) Run() {
	p.server.Run()
}

func NewProxy() *Proxy {
	var handlers = map[string]func(w nhttp.ResponseWriter, r *nhttp.Request){}
	handlers["/"] = func(w nhttp.ResponseWriter, r *nhttp.Request) {
		response := "H" +
			"ello http server by golang!\n"
		response += "Now: " + time.Now().String() + "\n"
		io.WriteString(w, response)
	}
	s := http.NewServer(handlers)
	return &Proxy{s}
}
