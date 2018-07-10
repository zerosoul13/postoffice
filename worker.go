package postoffice

import (
	"fmt"
	"errors"
	"strings"

	"github.com/zerosoul13/postoffice/workers"
)

var deliveryOptions map[string]workers.WorkerFactory
var deliveryOptionNotRegistered = errors.New("Delivery option is not registered")

func init() {
	// By doing early registration we ensure these defaults
	// are available
	workers.Register("graphite", workers.NewGraphiteWorker)
	workers.Register("rabbitmq", workers.NewRabbitMQWorker)
	deliveryOptions = workers.GetFactories() // We get our workers using this factories
}

func NewWorker(conf map[string]string) (workers.FactoryWorker, error) {
	wName := conf["type"]
	wFactory, ok := deliveryOptions[wName]

	if !ok {
		// Worker has not been registered.
		// Make a list of all available worker factories for logging.
		availableDeliveryOpts := make([]string, len(deliveryOptions))

		for k, _ := range deliveryOptions {
			availableDeliveryOpts = append(availableDeliveryOpts, k)
		}

		fmt.Printf("Invalid delivery option name. Must be one of: %s", strings.Join(availableDeliveryOpts, ", "))
		return nil, deliveryOptionNotRegistered
	}

	// Run the factory with the configuration
	return wFactory(conf)
}

//Send is used to send out message to chosen deliveryOpt 
func Send(msg []byte, deliveryOpt workers.FactoryWorker) {
	fmt.Println(deliveryOpt.Name())
	fmt.Println(deliveryOpt.Deliver(msg))

}

// Broadcast is used to publish a single message to X number of delivery
// options
func Broadcast(msg []byte, deliveryOpts []workers.FactoryWorker) {
	for _, d := range deliveryOpts {
		fmt.Println(d.Name())
		fmt.Println(d.Deliver(msg))
	}
}
