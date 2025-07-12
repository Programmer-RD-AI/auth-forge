package repository

import (
	"context"
	"fmt"

	"github.com/Programmer-RD-AI/auth-forge/config"
	errors "github.com/Programmer-RD-AI/auth-forge/pkg/authforge/"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	name   string
	config config.RedisConfig
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository(config config.RedisConfig) *RedisRepository {
	return &RedisRepository{
		name:   "Redis",
		config: config,
		client: nil,
		ctx:    nil,
	}
}

func (r *RedisRepository) Connect() {
	redisClient := redis.NewClient(createRedisOption(&r.config))
	initCtx := context.Background()
	r.client = redisClient
	r.ctx = initCtx
}

func (r *RedisRepository) HealthCheck() (bool, error) {
	if err := r.client.Ping(r.ctx).Err(); err != nil {
		return false, errors.NewDbHealthCheckFail(r.name, fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	return true, nil
}

func (r *RedisRepository) Read(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", &errors.KeyDoesNotExistError{key: key}
	} else if err != nil {
		return "", err
	} else {
		return val, nil
	}
}

func (r *RedisRepository) Close() bool {
	if err := r.client.Close(); err != nil {
		return false
	}
	return true
}

func (r *RedisRepository) Create(key string, value string) bool {
	if err := r.client.Set(r.ctx, key, &value, 0).Err(); err != nil {
		fmt.Println("Error creating key:", err)
		return true
	}
	return false
}

func (r *RedisRepository) Delete(key string) (int64, error) {
	deleted, err := r.client.Del(r.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (r *RedisRepository) Update(key string, value *any) bool {
	if err := r.client.Set(r.ctx, key, value, 0).Err(); err != nil {
		return false
	}
	return true
}

func createRedisOption(redisConfig *config.RedisConfig) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	}
}
