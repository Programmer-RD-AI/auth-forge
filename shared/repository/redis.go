package repository

import (
	"context"
	"fmt"
	"github.com/Programmer-RD-AI/auth-forge/internal/errors"
	"github.com/Programmer-RD-AI/auth-forge/pkg/util"
	"github.com/redis/go-redis/v9"
)

type RedisConfiguration struct {
	Host     string
	Port     int
	Password string
	Database int
}

type RedisRepository struct {
	BaseRepository	
	config RedisConfiguration
	client *redis.Client
}

func CreateRedisConfiguration(host *string, port *int, password *string, database *int) RedisConfiguration {
	if host == nil {
		h, _ := util.GetEnv[string]("REDIS_HOST", "localhost")
		host = &h
	}
	if port == nil {
		p, _ := util.GetEnv[int]("REDIS_PORT", 6379)
		port = &p
	}
	if password == nil {
		pass, _ := util.GetEnv[string]("REDIS_PASSWORD", "")
		password = &pass
	}
	if database == nil {
		db, _ := util.GetEnv[int]("REDIS_DATABASE", 0)
		database = &db
	}
	return RedisConfiguration{
		Host:     *host,
		Port:     *port,
		Password: *password,
		Database: *database,
	}
}

func (r *RedisRepository) Connect() (*redis.Client, context.Context) {
	redisClient := redis.NewClient(createRedisOption(&r.config))
	initCtx := context.Background()
	return redisClient, initCtx
}

func (r *RedisRepository) HealthCheck(ctx context.Context, client *redis.Client) (bool, error) {
	if err := client.Ping(ctx).Err(); err != nil {
		return false, errors.dbHealthCheckFail{message: err.Error(), provider: r.BaseRepository.providerName}
	}	
	return true, nil
}

func (r *RedisRepository) Read(client *redis.Client, ctx context.Context, key *string){
	val, err := client.Get(ctx, *key).Result()
}
func createRedisOption(redisConfig *RedisConfiguration) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	}
}
