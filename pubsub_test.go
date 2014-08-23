package pubsub_test

import (
    . "github.com/puffinframework/nats.pubsub"
	"github.com/apcera/nats"
    . "github.com/puffinframework/pubsub/impltests"
	"testing"
)

func TestRemote(t *testing.T) {
	pb := NewPubSub(nats.DefaultURL)
	TestSubscribe(t, pb)
	pb.Close()

	pb = NewPubSub(nats.DefaultURL)
	TestUnsubscribe(t, pb)
	pb.Close()

	pb = NewPubSub(nats.DefaultURL)
	TestSubscribeSync(t, pb)
	pb.Close()
}
