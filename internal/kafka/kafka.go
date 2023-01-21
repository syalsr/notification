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
	topicCommin string
	partition   int32
	offset      int64
	consumer    sarama.Consumer
}

// NewConsumer -
func NewConsumer(cfg *config.App) (*Client, error) {
	saramaCfg := sarama.Config{}
	client, err := sarama.NewConsumer(cfg.KafkaURL, &saramaCfg)
	if err != nil {
		log.Err(err).Msgf("cant create new consumer: %w", err)
		return nil, fmt.Errorf("cant create new consumer: %w", err)
	}

	return &Client{
		topicPerson: cfg.KafkaTopicPersonalized,
		topicCommin: cfg.KafkaTopicCommon,
		partition:   cfg.KafkaParition,
		offset:      cfg.KafkaOffset,
		consumer:    client,
	}, nil
}

// Run starts the listener and pushes task IDs to the provided strings channel
func (c *Client) Run(ctx context.Context, common, person chan<- string) {
	personTopic, err := c.CreateTopic(c.topicPerson, c.partition, c.offset)
	if err != nil {
		log.Err(err).Msgf("cant create cosumer partition: %w", err)
		return
	}
	commonTopic, err := c.CreateTopic(c.topicCommin, c.partition, c.offset)
	if err != nil {
		log.Err(err).Msgf("cant create cosumer partition: %w", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-personTopic.Messages():
			log.Info().Msgf("Consumer get %s", string(msg.Value))
			person <- string(msg.Value)
		case msg := <-commonTopic.Messages():
			log.Info().Msgf("Consumer get %s", string(msg.Value))
			common <- string(msg.Value)
		}
	}
}

// CreateTopic - create topic for person and common email
func (c *Client) CreateTopic(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error) {
	pCons, err := c.consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Err(err).Msgf("cant create cosumer partition: %w", err)
		return nil, err
	}
	return pCons, nil
}

func (c *Client) Close() error {
	return c.consumer.Close()
}
