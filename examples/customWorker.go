package main

import (
	"fmt"

	"github.com/zerosoul13/postoffice"
	"github.com/zerosoul13/postoffice/workers"
)

const importantMessage = "This is a test"
type customWorker struct {}

func (cw customWorker) Name() string {
	return "Custom"
}

func (cw customWorker) Deliver(msg []byte) error {
	fmt.Println(string(msg))
	return nil
}

func NewCustomWorker(conf map[string]string) (workers.FactoryWorker, error) {
	c := customWorker{}
	return c, nil
}

func main() {
	conf := map[string]string{"type": "custom"}
	workers.Register("custom", NewCustomWorker)
	c, err := postoffice.NewWorker(conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.Name())
	m := []byte(importantMessage)
	c.Deliver(m)
}
