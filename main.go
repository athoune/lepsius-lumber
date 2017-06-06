package main

import (
	"fmt"
	"github.com/elastic/go-lumber/server"
)

func main() {
	s, err := server.ListenAndServe(":5044", server.V2(true))
	if err != nil {
		panic(err)
	}
	for {
		batch := s.Receive()
		for _, event := range batch.Events {
			fmt.Println(event)
		}
		batch.ACK()
	}

}
