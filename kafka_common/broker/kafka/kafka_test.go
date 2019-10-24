package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"testing"
	"xinyu/go_splitter/broker"
)

func Test_kBroker_Subscribe(t *testing.T) {
	config := DefaultClusterConfig
	config.Version = sarama.V0_10_2_1

	client := NewBroker(ClusterConfig(config))

	err := client.Connect()
	if err != nil {
		return
	}

	//err = client.Publish("btc", &broker.Message{
	//	Body: []byte{},
	//})
	//
	//err = client.Publish("btc", &broker.Message{
	//	Body: []byte{},
	//})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go client.Subscribe("btc", func(event broker.Event) error {
		fmt.Println(event.Message())
		_ = event.Ack()
		//wg.Done()
		return nil
	}, func(options *broker.SubscribeOptions) {
		options.Queue = "example1"
	})

	//err = client.Publish("btc", &broker.Message{
	//	Body: []byte{},
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	wg.Wait()
}
