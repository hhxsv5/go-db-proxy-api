package main

import (
	"github.com/hhxsv5/go-db-proxy-api"
	"github.com/hhxsv5/go-db-proxy-api/examples/models/mydefault"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	var handlers = map[string]func(w http.ResponseWriter, r *http.Request){}
	handlers["/"] = func(w http.ResponseWriter, r *http.Request) {
		//rand.Seed(time.Now().UnixNano())
		//cellphone := "1878020" + strconv.Itoa(int(1000+rand.Int31n(9999-1000)))
		//mydefault.CreateUser(cellphone)

		users := mydefault.GetUsersByIds([]uint64{})

		var (
			rsp string
			it  time.Time
			mt  time.Time
		)

		for i, u := range users {
			it = time.Unix(u.InsertTime, 0)
			mt = time.Unix(u.ModifyTime, 0)
			rsp += fmt.Sprintf("%d: %d, %s, %s, %s\n", i, u.Id, u.Cellphone, it.Format(time.RFC3339), mt.Format("2006-01-02 15:04:05"))
		}

		io.WriteString(w, rsp)
	}

	p := godpa.NewProxy(handlers)
	p.Run()
}
