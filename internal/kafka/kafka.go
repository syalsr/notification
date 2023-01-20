package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/notification/internal/config"
)

type Client struct {
	topic     string
	partition int32
	offset    int64
	consumer  sarama.Consumer
}

func NewConsumer(cfg *config.App) (*Client, error) {
	saramaCfg := sarama.Config{}
	client, err := sarama.NewConsumer(cfg.KafkaURL, &saramaCfg)
	if err != nil {
		log.Err(err).Msgf("cant create new consumer: %w", err)
		return nil, fmt.Errorf("cant create new consumer: %w", err)
	}

	return &Client{topic: cfg.KafkaTopic, partition: cfg.KafkaParition, offset: cfg.KafkaOffset, consumer: client}, nil
}

// Run starts the listener and pushes task IDs to the provided strings channel
func (c *Client) Run(ctx context.Context) {
	pCons, err := c.consumer.ConsumePartition(c.topic, c.partition, c.offset)
	if err != nil {
		log.Err(err).Msgf("cant create cosumer partition: %w", err)
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-pCons.Messages():
			log.Info().Msg(string(msg.Value))
		default:
			continue
		}
	}
}
