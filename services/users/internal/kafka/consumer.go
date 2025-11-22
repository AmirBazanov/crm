package kafka

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	logger *slog.Logger
	reader *kafka.Reader
	router *EventRouter
}

func New(
	brokers []string,
	topic string,
	groupID string,
	logger *slog.Logger,
	router *EventRouter,
) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})

	return &Consumer{
		logger: logger,
		reader: reader,
		router: router,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	c.logger.Info("Kafka consumer started " + c.reader.Config().Brokers[0])

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			c.logger.Error("kafka read", "error", err)
			continue
		}
		eventType := ""
		for _, h := range msg.Headers {
			if h.Key == "event-type" {
				eventType = string(h.Value)
			}
		}
		//TODO: Looping when cant create
		err = c.router.Route(ctx, eventType, msg.Value)
		if err != nil {
			c.logger.Error("kafka route", "error", err)
			return err
		}

	}
}

func (c *Consumer) Stop() error {
	c.logger.Info("Kafka consumer stopped")
	return c.reader.Close()
}
