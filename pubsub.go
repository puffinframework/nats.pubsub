package pubsub

import (
    . "github.com/puffinframework/pubsub"
	"time"
	"github.com/apcera/nats"
)

type natsPubSub struct {
	conn *nats.Conn
}

func NewPubSub(natsURL string) PubSub {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		panic(err)
	}
	return &natsPubSub{conn: conn}
}

func (self *natsPubSub) Close() {
	self.conn.Close()
}

func (self *natsPubSub) Subscribe(topic string, callback Callback) (Subscription, error) {
	sub, err := self.conn.Subscribe(topic, func(msg *nats.Msg) {
		callback(msg.Data)
	})
	return &natsSubscription{sub: sub}, err
}

func (self *natsPubSub) SubscribeSync(topic string, callback CallbackSync) (Subscription, error) {
	sub, err := self.conn.Subscribe(topic, func(msg *nats.Msg) {
		result, err := callback(msg.Data)
		if err != nil {
			// TODO
		} else {
			self.conn.Publish(msg.Reply, result)
		}
	})
	return &natsSubscription{sub: sub}, err
}

func (self *natsPubSub) Publish(topic string, data []byte) error {
	return self.conn.Publish(topic, data)
}

func (self *natsPubSub) PublishSync(topic string, data []byte, timeout time.Duration) ([]byte, error) {
	msg, err := self.conn.Request(topic, data, timeout)
	if err != nil {
		return nil, err
	}
	return msg.Data, err
}

type natsSubscription struct {
	sub *nats.Subscription
}

func (self *natsSubscription) Unsubscribe() error {
	return self.sub.Unsubscribe()
}
