package main

import (
	"fmt"

	"github.com/go-url-shortener/model"
	"github.com/go-url-shortener/server"
)

func main() {
	fmt.Println("Go-Url-Shortener")
	model.SetUp()
	server.ServeAndListen()
}
