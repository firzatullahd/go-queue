package queue

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Client         redis.Client
	Queue          string
	TimeoutSeconds int
}

type Queue struct {
	Config Config
}

func New(config Config) *Queue {
	return &Queue{Config: config}
}

func (q *Queue) Publish(ctx context.Context, message any) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return q.Config.Client.LPush(ctx, q.Config.Queue, jsonMessage).Err()
}
