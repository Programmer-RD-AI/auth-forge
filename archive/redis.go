// package store

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/Programmer-RD-AI/auth-forge/config"
// 	"github.com/Programmer-RD-AI/auth-forge/pkg/authforge"
// 	"github.com/redis/go-redis/v9"
// )

// const (
// 	defaultExpire time.Duration = 0 * time.Second
// 	providerName                = "redis"
// )

// type RedisStore struct {
// 	name   string
// 	config config.RedisConfig
// 	client *redis.Client
// 	ctx    context.Context
// }

// func NewRedisStore(config config.RedisConfig) *RedisStore {
// 	return &RedisStore{
// 		name:   providerName,
// 		config: config,
// 		client: nil,
// 		ctx:    nil,
// 	}
// }

// func (r *RedisStore) Connect() {
// 	redisClient := redis.NewClient(createRedisOption(&r.config))
// 	initCtx := context.Background()
// 	r.client = redisClient
// 	r.ctx = initCtx
// }

// func (r *RedisStore) HealthCheck() (bool, error) {
// 	if err := r.client.Ping(r.ctx).Err(); err != nil {
// 		return false, authforge.NewDbHealthCheckFail(r.name, fmt.Sprintf("Failed to connect to Redis: %v", err))
// 	}
// 	return true, nil
// }

// func (r *RedisStore) Read(key string) (string, error) {
// 	val, err := r.client.Get(r.ctx, key).Result()
// 	if err == redis.Nil {
// 		return "", authforge.NewKeyDoesNotExistError(key)
// 	} else if err != nil {
// 		return "", err
// 	} else {
// 		return val, nil
// 	}
// }

// func (r *RedisStore) Close() bool {
// 	if err := r.client.Close(); err != nil {
// 		return false
// 	}
// 	return true
// }

// func (r *RedisStore) Create(key string, value string, TTL *int) bool {
// 	var expire time.Duration
// 	if TTL != nil {
// 		expire = time.Duration(*TTL) * time.Second
// 	} else {
// 		expire = defaultExpire
// 	}
// 	if err := r.client.Set(r.ctx, key, &value, expire).Err(); err != nil {
// 		fmt.Println("Error creating key:", err)
// 		return true
// 	}
// 	return false
// }

// func (r *RedisStore) Delete(key string) (int64, error) {
// 	deleted, err := r.client.Del(r.ctx, key).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return deleted, nil
// }

// func (r *RedisStore) Update(key string, value *any) bool {
// 	if err := r.client.Set(r.ctx, key, value, 0).Err(); err != nil {
// 		return false
// 	}
// 	return true
// }

//	func createRedisOption(redisConfig *config.RedisConfig) *redis.Options {
//		return &redis.Options{
//			Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
//			Password: redisConfig.Password,
//			DB:       redisConfig.Database,
//		}
//	}
