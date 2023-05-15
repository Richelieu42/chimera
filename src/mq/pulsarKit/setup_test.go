package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu42/chimera/v2/src/idKit"
	"github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

func TestSetUp(t *testing.T) {
	pulsarConfig := &Config{
		Addresses: []string{"192.168.80.27:6650", "192.168.80.42:6650", "192.168.80.43:6650"},
		//Addresses: []string{"172.18.21.50:6650"},
		VerifyConfig: &VerifyConfig{
			Topic: "test",
			Print: true,
		},
	}
	MustSetUp(pulsarConfig)

	topic := "test"
	consumer, err := NewConsumer(context.TODO(), pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: idKit.NewULID(),
		Type:             pulsar.Exclusive,
	}, "test-consumer.log")
	if err != nil {
		logrus.Fatal(err)
	}
	defer consumer.Close()

	producer, err := NewProducer(context.TODO(), pulsar.ProducerOptions{
		Topic:       topic,
		SendTimeout: sendTimeout,
	}, "test-producer.log")
	if err != nil {
		logrus.Fatal(err)
	}
	defer producer.Close()

	/* 发消息 */
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)

			text := strconv.Itoa(i)
			id, err := producer.Send(context.TODO(), &pulsar.ProducerMessage{
				Payload: []byte(text),
			})
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"text": text,
				}).Error("[PRODUCER] fail to send message")
			} else {
				logrus.WithFields(logrus.Fields{
					"id":   id.String(),
					"text": text,
				}).Info("[PRODUCER] succeed to send message")
			}
		}
	}()

	/* 收消息 */
	p, err := ants.NewPoolWithFunc(256, func(i interface{}) {
		msg, ok := i.(pulsar.Message)
		if !ok {
			return
		}
		logrus.WithFields(logrus.Fields{
			"id":   msg.ID().String(),
			"text": string(msg.Payload()),
		}).Info("[CONSUMER] receive a message")
	})
	if err != nil {
		logrus.Fatal(err)
	}
	defer p.Release()
	for {
		msg, err := consumer.Receive(context.TODO())
		if err != nil {
			logrus.Error("[CONSUMER] fail to receive")
			break
		}
		if err := consumer.Ack(msg); err != nil {
			logrus.WithFields(logrus.Fields{
				"id":   msg.ID().String(),
				"text": string(msg.Payload()),
			}).Error("[CONSUMER] fail to ack")
			continue
		}

		if err := p.Invoke(msg); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("[CONSUMER] fail to invoke")
		}
	}
}
