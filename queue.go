package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Queue struct {
	Client  redis.Client
	Queue   string
	Timeout time.Duration
}

func New(client redis.Client, queue string, timeout time.Duration) *Queue {
	return &Queue{Client: client, Queue: queue, Timeout: timeout}
}

func (q *Queue) Publish(ctx context.Context, message any) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return q.Client.LPush(ctx, q.Queue, jsonMessage).Err()
}

func (q *Queue) Consume(ctx context.Context, handler func(message string)) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			result, err := q.Client.BRPop(ctx, q.Timeout, q.Queue).Result()
			if err != nil {
				if err == redis.Nil {
					continue
				}
				return err
			}

			// result[0] is key, result[1] is value
			if len(result) == 2 {
				handler(result[1])
			}
		}
	}
}
