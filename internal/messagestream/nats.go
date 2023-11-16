package messagestream

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go-streams/internal/messagestream/handlers"
	"go.uber.org/zap"
	"time"
)

type NatsConfig struct {
	Url string
}

func GetConfig() *NatsConfig {
	return &NatsConfig{
		Url: "",
	}
}

func CreateNatsConnection(config *NatsConfig) *nats.Conn {
	nc, err := nats.Connect(config.Url)
	if err != nil {
		panic(err)
	}

	return nc
}
func CreateJetStream(conn *nats.Conn) jetstream.JetStream {
	js, _ := jetstream.New(conn)

	return js
}

const ConversionsStreamName = "CONVERSION"

func CreateConversionJetStream(ctx context.Context, js jetstream.JetStream) jetstream.Stream {
	stream, _ := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     ConversionsStreamName,
		Subjects: []string{Conversion.String()},
		MaxAge:   360 * time.Second,
		Discard:  jetstream.DiscardOld,
	})

	return stream
}

func CreateConversionConsumer(ctx context.Context, js jetstream.JetStream) jetstream.Consumer {
	consumer, _ := js.CreateOrUpdateConsumer(ctx, ConversionsStreamName, jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverNewPolicy,
	})

	return consumer
}

func ListenToConsumer(consumer jetstream.Consumer, logger *zap.Logger, handler handlers.IBaseHandler) {
	consContext, _ := consumer.Consume(func(msg jetstream.Msg) {
		err := handler.Handle(msg)

		if err != nil {
			logger.Error(err.Error())
			err := msg.Nak()
			if err != nil {
				logger.Error(err.Error())
			}
		} else {
			err := msg.Ack()
			if err != nil {
				logger.Error(err.Error())
			}
		}
	})
	defer consContext.Stop()
}
