package main

import (
	"fmt"

	"github.com/go-url-shortener/model"
	"github.com/go-url-shortener/server"
)

func main() {
	fmt.Println("Go-Url-Shorter")
	model.SetUp()
	server.ServeAndListen()
}
