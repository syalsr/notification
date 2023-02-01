package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/notification/internal/config"
)

// Client -
type Client struct {
	topicPerson string
	topicCommon string
	partition   int32
	consumer    sarama.Consumer
}

// NewConsumer -
func NewConsumer(cfg *config.App) (*Client, error) {
	saramsCfg := sarama.NewConfig()
	client, err := sarama.NewConsumer(cfg.Kafka.Dsn, saramsCfg)
	if err != nil {
		log.Err(err).Msgf("cant create new consumer: %w", err)
		return nil, fmt.Errorf("cant create new consumer: %w", err)
	}

	return &Client{
		topicPerson: cfg.Kafka.TopicPersonalized,
		topicCommon: cfg.Kafka.TopicCommon,
		partition: 0,
		consumer:    client,
	}, nil
}

// Run starts the listener and pushes task IDs to the provided strings channel
func (c *Client) Run(ctx context.Context, common, person chan<- string) {
	personPartition, err := c.consumer.ConsumePartition(c.topicPerson, c.partition, sarama.OffsetOldest)
	if err != nil {
		log.Err(err).Msg("error while consume person partition")
		return
	}

	commonPartition, err := c.consumer.ConsumePartition(c.topicCommon, c.partition, sarama.OffsetOldest)
	if err != nil {
		log.Err(err).Msg("error while consume common partition")
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-personPartition.Messages():
			log.Info().Msgf("Consumer get %s", string(msg.Value))
			person <- string(msg.Value)
		case msg := <-commonPartition.Messages():
			log.Info().Msgf("Consumer get %s", string(msg.Value))
			common <- string(msg.Value)
		}
	}
}

// Close - close connection to consumer
func (c *Client) Close() error {
	return c.consumer.Close()
}
