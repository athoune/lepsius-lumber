package main

import (
	"fmt"
	"github.com/athoune/lepsius-lumber/sql"
	"github.com/elastic/go-lumber/server"
	"github.com/percona/go-mysql/query"
	"reflect"
)

func readPacket(raw []byte) (interface{}, error) {
	pb := sql.Packetbeat{}
	err := pb.UnmarshalJSON(raw)
	if err != nil {
		return nil, err
	}
	switch {
	case pb.Type == "mysql":
		m := sql.Mysql{}
		err = m.UnmarshalJSON(raw)
		if err != nil {
			return nil, err
		}
		return m, nil
	default:
		return pb, nil
	}
}

func main() {
	s, err := server.ListenAndServe(":5044", server.V2(true), server.JSONDecoder(readPacket))
	if err != nil {
		panic(err)
	}
	for {
		batch := s.Receive()
		for _, event := range batch.Events {
			switch v := event.(type) {
			case sql.Mysql:
				fmt.Printf("Mysql method:%s response time: %v\n%s\n%#v\n",
					v.Method, v.ResponseTime, query.Fingerprint(v.Query), v.Mysql)
			default:
				fmt.Println("Event: ", reflect.TypeOf(event))
			}
		}
		batch.ACK()
	}
}
