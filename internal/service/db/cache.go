package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
)

type cacheInput struct {
	R7K1Neighbors  []string
	R8K1Neighbors  []string
	R8K2Neighbors  []string
	R9K1Neighbors  []string
	R9K2Neighbors  []string
	R10K1Neighbors []string
	R10K2Neighbors []string
}

func cacheTrip(ctx context.Context, redisClient *redis.Client, trip models.Trip) error {
	return cache(ctx, redisClient, "trip", cacheInput{
		R7K1Neighbors:  trip.R7K1Neighbors,
		R8K1Neighbors:  trip.R8K1Neighbors,
		R8K2Neighbors:  trip.R8K2Neighbors,
		R9K1Neighbors:  trip.R9K1Neighbors,
		R9K2Neighbors:  trip.R9K2Neighbors,
		R10K1Neighbors: trip.R10K1Neighbors,
		R10K2Neighbors: trip.R10K2Neighbors,
	})
}

func cache(ctx context.Context, redisClient *redis.Client, prefix string, in cacheInput) error {
	for _, cell := range in.R7K1Neighbors {
		pipe := redisClient.Pipeline()
		key := fmt.Sprintf("%s-%s", prefix, cell)

		// Get bytes (of string slice)
		b, err := pipe.Get(ctx, key).Bytes()
		if err != nil {
			return err
		}

		// Convert bytes to string slice

		pipe.Set()

		// TODO when does the trip expire? auto-expire the trip a day after its scheduled_for
		//pipe.Expire(ctx, key, time.Hour)

		cmds, err := pipe.Exec(ctx)
		if err != nil {
			panic(err)
		}
		s.redisClient.Set(ctx, fmt.Sprintf("trip-"), value)
	}
}
