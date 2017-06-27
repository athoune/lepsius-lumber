package main

import (
	"fmt"
	"github.com/athoune/lepsius-lumber/sql"
	"github.com/elastic/go-lumber/server"
	"github.com/percona/go-mysql/query"
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
			beat, ok := event.(sql.Packetbeat)
			if !ok {
				fmt.Println("Not a beat: ", event)
				continue
			}
			if beat.Type == "mysql" {
				m, ok := event.(sql.Mysql)
				if !ok {
					panic("Cast drama")
				}
				fmt.Printf("Mysql method:%s response time: %v\n%s\n%#v\n", m.Method, m.ResponseTime, query.Fingerprint(m.Query), m.Mysql)
			}
		}
		batch.ACK()
	}
}
