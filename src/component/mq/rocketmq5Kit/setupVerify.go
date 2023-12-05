package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/protocol/v2"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/time/timeKit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

const (
	sendTimeout = time.Millisecond * 500

	// verifyTimeout 验证的最长timeout
	verifyTimeout = time.Second * 6
)

// verify 测试RocketMQ5服务是否启动正常.
/*
@param endpoint 用于测试的RocketMQ5服务的endpoint
@param topic 	用于测试的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）
@return 如果为nil，说明 RocketMQ5服务 正常启动
*/
func verify(topic string) error {
	if strKit.IsEmpty(topic) {
		// 不进行验证
		return nil
	}

	/* logger */
	//readWriter := ioKit.NewReadWriter(nil)
	readWriter := os.Stderr
	logger := logrusKit.NewLogger(logrusKit.WithOutput(readWriter),
		logrusKit.WithLevel(logrus.DebugLevel),
		logrusKit.WithDisableQuote(true),
	)
	logger.Infof("topic: [%s].", topic)

	/* texts */
	timeStr := timeKit.FormatCurrent(timeKit.FormatCommon)
	ulid := idKit.NewULID()
	texts := []string{
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$3"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$4"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$5"),
	}
	json, err := jsonKit.MarshalIndentToString(texts, "", "    ")
	if err != nil {
		return err
	}
	logger.Infof("texts:\n%s\n.", json)

	/* (1) producer */
	producer, err := NewProducer()
	if err != nil {
		return err
	}
	defer producer.GracefulStop()

	/* (2) consumer */
	//consumerGroup := idKit.NewULID()
	consumerGroup := "ccc"
	consumer, err := NewSimpleConsumer(consumerGroup, map[string]*rmq_client.FilterExpression{
		topic: rmq_client.SUB_ALL,
	})
	if err != nil {
		return err
	}
	defer consumer.GracefulStop()

	ctx, cancel := context.WithTimeout(context.TODO(), verifyTimeout)
	defer cancel()
	var producerCh chan error = make(chan error, 1)
	var consumerCh chan error = make(chan error, 1)

	/* (3) producer goroutine */
	go func() {
		defer func() {
			logger.Info("[PRODUCER] Goroutine ends.")
		}()

		for _, text := range texts {
			msg := &rmq_client.Message{
				Topic: topic,
				Body:  []byte(text),
			}
			ctx, _ := context.WithTimeout(context.TODO(), sendTimeout)
			_, err := producer.Send(ctx, msg)
			if err != nil {
				err = errorKit.Wrap(err, "Fail to send message(%s).", text)
				producerCh <- err
				return
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[PRODUCER] Manager to send a message.")
		}
		logger.Info("[PRODUCER] Manager to send all messages.")
	}()

	/* (4) consumer goroutine */
	textsCopy := sliceKit.Copy(texts)
	go func(texts []string) {
		defer func() {
			logger.Info("[CONSUMER] Goroutine ends.")
		}()

		for {
			select {
			case <-ctx.Done():
				consumerCh <- errorKit.New("Fail to receive all messages within timeout(%s).", verifyTimeout)
				return
			case <-time.After(time.Millisecond * 100):
				// do nothing
			}
			//time.Sleep(time.Millisecond * 100)

			mvs, err := consumer.Receive(context.TODO(), MaxMessageNum, InvisibleDuration)
			if err != nil {
				/* gRPC errors */
				if s, ok := status.FromError(err); ok {
					switch s.Code() {
					case codes.Canceled:
						consumerCh <- err
						return
						//case codes.Canceled:
						//	/* 提前结束（被取消） */
						//	break LOOP
						//case codes.DeadlineExceeded:
						//	/* 超时结束 */
						//	consumerErr = errorKit.New("consumer fails to receive all messages(count: %d) within timeout(%s), missing(%d)", len(texts), verifyTimeout.ToDsnString(), len(text1))
						//	break LOOP
					}
				}

				if errRpcStatus, ok := rmq_client.AsErrRpcStatus(err); ok {
					/* RocketMQ5 errors */
					switch errRpcStatus.Code {
					case int32(v2.Code_MESSAGE_NOT_FOUND):
						// 没有新消息
					default:
						logger.WithFields(logrus.Fields{
							"code":  errRpcStatus.Code,
							"error": err.Error(),
						}).Warn("[Consumer] Fail to receive.")
					}
				} else {
					/* other errors */
					logger.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Warn("[Consumer] Fail to receive.")
				}
				continue
			}

			// ack message
			for _, mv := range mvs {
				text := string(mv.GetBody())

				err := consumer.Ack(context.TODO(), mv)
				if err != nil {
					logger.WithFields(logrus.Fields{
						"text":  text,
						"error": err.Error(),
					}).Error("[Consumer] Fail to ack.")
					continue
				}

				var ok bool
				texts, ok = sliceKit.Remove(texts, text)
				left := len(texts)
				logger.WithFields(logrus.Fields{
					"valid": ok,
					"left":  left,
					"text":  text,
				}).Info("[CONSUMER] Receive and ack a message.")
				if left == 0 {
					// 成功收到所有预期消息
					logger.Info("[CONSUMER] Receive and ack all messages.")

					consumerCh <- nil
					return
				}
			}
		}
	}(textsCopy)

	select {
	case producerErr := <-producerCh:
		return producerErr
	case consumerCh := <-consumerCh:
		if consumerCh != nil {
			return consumerCh
		}
		// 通过验证
		return nil
	case <-ctx.Done():
		return errorKit.New("Fail to pass verification within timeout(%s).", verifyTimeout)
	}
}
