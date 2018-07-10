package main

import (
	"fmt"

	"github.com/zerosoul13/postoffice"
	"github.com/zerosoul13/postoffice/workers"
)

const importantMessage = "This is a test"

func main() {
	conf := map[string]string{"type": "graphite", "HOSTNAME": "127.0.0.1", "PORT": "2005"}
	bmsg := []byte(importantMessage)
	graphite, err := postoffice.NewWorker(conf)
	if err != nil {
		fmt.Println(err)
	}

	rbtconf := map[string]string{"type": "rabbitmq", "HOSTNAME": "127.0.0.1", "PORT": "5672"}
	rabbitmq, err := postoffice.NewWorker(rbtconf)
	if err != nil {
		fmt.Println(err)
	}

	var ws []workers.FactoryWorker
	ws = append(ws, graphite, rabbitmq)

	postoffice.Broadcast(bmsg, ws)
}
