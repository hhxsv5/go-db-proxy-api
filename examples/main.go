package main

import (
	"github.com/hhxsv5/go-db-proxy-api"
)

func main() {
	//user := mydefault.CreateUser("18780207350")
	//fmt.Println(user)
	//fmt.Printf("%p", user)

	p := godpa.NewProxy()
	p.Run()
}
