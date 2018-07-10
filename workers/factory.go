package workers

import (
	"fmt"
	"log"
	"errors"
)

var ErrInvalidWorkerConfig = errors.New("Check worker configuration")

// FactoryWorker is the interface we expect for out workers
type FactoryWorker interface {
	Name() string
	Deliver(msg []byte) error
}

// WorkerFactory provides an easy interface for Worker configurations
// All workers must have their own factory implementation
type WorkerFactory func(conf map[string]string) (FactoryWorker, error)

// Keeps track of registered workers
var workerFactories = make(map[string]WorkerFactory)

func GetFactories() map[string]WorkerFactory {
	return workerFactories
}

// Register receives a new registration candidate with name and a Worker
// if worker is already registered the candidate will be ignored
// else registration will continue
func Register(name string, w WorkerFactory) {
	if w == nil {
		log.Panicf("Worker %s does not exist.", name)
	}
	_, registered := workerFactories[name]
	if registered {
		fmt.Printf("Worker %s already registered. Ignoring.", name)
	}
	workerFactories[name] = w
}
