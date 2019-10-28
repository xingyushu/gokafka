package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"xinyu/go_splitter/broker"
	//"xinyu/go_splitter/component"
)

var (
	DefaultBrokerConfig  = sarama.NewConfig()
	DefaultClusterConfig = sarama.NewConfig()
)

type brokerConfigKey struct{}
type clusterConfigKey struct{}

func BrokerConfig(c *sarama.Config) broker.Option {
	return setBrokerOption(brokerConfigKey{}, c)
}

func ClusterConfig(c *sarama.Config) broker.Option {
	return setBrokerOption(clusterConfigKey{}, c)
}

type subscribeContextKey struct{}

// SubscribeContext set the context for broker.SubscribeOption
func SubscribeContext(ctx context.Context) broker.SubscribeOption {
	return setSubscribeOption(subscribeContextKey{}, ctx)
}

// consumerGroupHandler is the implementation of sarama.ConsumerGroupHandler
type consumerGroupHandler struct {
	handler broker.Handler
	subopts broker.SubscribeOptions
	kopts   broker.Options
	cg      sarama.ConsumerGroup
	sess    sarama.ConsumerGroupSession
}

func (*consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		//fmt.Println(string(msg.Value))

		var m =broker.Message{
			Body:msg.Value,
		}



		err := h.handler(&publication{
			m:    &m,
			t:    msg.Topic,
			km:   msg,
			cg:   h.cg,
			sess: sess,
		})

		if err == nil && h.subopts.AutoAck {
			sess.MarkMessage(msg, "")
		}
		//sess.MarkMessage(msg, "")
	}
	return nil
}
