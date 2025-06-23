package example

import (
	"context"
	"time"

	goqueue "github.com/firzatullahd/go-queue"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	queue := goqueue.New(*redisClient, "my-queue", 5*time.Second)
	queue.Publish(context.Background(), "test")
}
