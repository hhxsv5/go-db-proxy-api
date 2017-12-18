package main

import (
	"github.com/hhxsv5/go-db-proxy-api"
	"time"
	"io"
	"net/http"
)

func main() {
	//user := mydefault.CreateUser("18780207350")
	//fmt.Println(user)
	//fmt.Printf("%p", user)

	var handlers = map[string]func(w http.ResponseWriter, r *http.Request){}
	handlers["/"] = func(w http.ResponseWriter, r *http.Request) {
		response := "Hello http server by golang!\n"
		response += "Now: " + time.Now().String() + "\n"
		io.WriteString(w, response)
	}

	p := godpa.NewProxy(handlers)
	p.Run()
}
