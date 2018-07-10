package workers

import "fmt"

// GRAPHITE CONFIGURATION
// HOSTNAME = 127.0.0.1
// PORT     = 2003
type GraphiteWorker struct {}

func NewGraphiteWorker(conf map[string]string) (FactoryWorker, error) {
	var g GraphiteWorker
	if conf == nil {
		return g, ErrInvalidWorkerConfig
	}
	g = GraphiteWorker{}
	return g, nil
}

func (gw GraphiteWorker) Name() string {
	return "Graphite"
}

func (gw GraphiteWorker) Deliver(msg []byte) error {
	fmt.Println(string(msg))
	return nil
}
