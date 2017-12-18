package main

import (
	"fmt"
	"github.com/hhxsv5/go-db-proxy-api/models/mdefault"
)

func main() {
	user := mdefault.CreateUser("18780207350")
	fmt.Println(user)
	fmt.Printf("%p", user)
}
