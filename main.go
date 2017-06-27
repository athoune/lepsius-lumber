package main

import (
	"fmt"
	"github.com/athoune/lepsius-lumber/sql"
	"github.com/elastic/go-lumber/server"
	"github.com/percona/go-mysql/query"
)

func main() {
	s, err := server.ListenAndServe(":5044", server.V2(true), server.JSONDecoder(func(raw []byte, v interface{}) error {
		pb := sql.Packetbeat{}
		err := pb.UnmarshalJSON(raw)
		if err != nil {
			return err
		}
		switch {
		case pb.Type == "mysql":
			m := sql.Mysql{}
			err = m.UnmarshalJSON(raw)
			if err != nil {
				return err
			}
			v = m
			fmt.Println("Mysql ", m.Method, m.ResponseTime, m.Query, query.Fingerprint(m.Query))
		default:
			v = pb
			fmt.Println("JSON : ", pb.Type, string(raw))
		}
		return nil
	}))
	if err != nil {
		panic(err)
	}
	for {
		batch := s.Receive()
		for _, event := range batch.Events {
			fmt.Printf("Event : %#v\n", event)
			beat, ok := event.(map[string]interface{})
			if ok {
				fmt.Printf("Type: %v\n", beat["type"])
			}
		}
		batch.ACK()
	}

}

type Packetbeat struct {
	Type string `json:"type"`
}
