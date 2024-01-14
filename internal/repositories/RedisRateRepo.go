package repositories

import (
	"TestTask/internal/configs"
	"TestTask/internal/models"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisRateRepo struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepo(cfg configs.Config) RateRepo {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("db:%s", cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	return &RedisRateRepo{client: client, ctx: ctx}
}

func (r *RedisRateRepo) SetRate(rate models.Rate) error {
	err := r.client.Set(r.ctx, rate.Symbol, rate.Price, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRateRepo) GetRate(symbol string) (models.Rate, bool) {
	val, err := r.client.Get(r.ctx, symbol).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return models.Rate{}, false
	case err != nil:
		fmt.Println("Get value redis error", err)
		return models.Rate{}, false
	case val == "":
		fmt.Println("value is empty")
		return models.Rate{}, false
	}
	return models.Rate{Symbol: symbol, Price: val}, true
}
