package main

import (
	"github.com/PhillipJeffries/goPractise"
	"Log"
)

func main(){
	srv := new(todo.Server)
	if err := srv.Run(port: "8000"); err != nil {
		log.Fatalf(format: "error: %s", err.Error())
	}
}