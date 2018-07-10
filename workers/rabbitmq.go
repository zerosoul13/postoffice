package workers

import "fmt"

type RabbitMQWorker struct {}

func NewRabbitMQWorker(conf map[string]string) (FactoryWorker, error) {
	var r RabbitMQWorker
	if conf == nil {
		return r, ErrInvalidWorkerConfig
	}
	r = RabbitMQWorker{}
	return r, nil
}
func (rw RabbitMQWorker) Name() string {
	return "RabbitMQ"
}

func (rw RabbitMQWorker) Deliver(msg []byte) error {
	fmt.Println(string(msg))
	return nil
}
