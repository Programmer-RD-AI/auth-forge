package repository

import "github.com/go-redis/redis/v9"

type RedisConfiguration struct {
	Host     string 
	Port     int
	Password string
	Database int
}

type RedisRepository struct {
	config RedisConfiguration
	client *redis.Client
}

func CreateRedisConfiguration(host *string, port *int, password *string, database *int){}

func (r *RedisReppository) Connect(){
	return redis.NewClient(createRedisOption(r.config));
}

func createRedisOption(redisConfig *RedisConfiguration) *redis.Options {
	return &redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
		Password: redisConfig.Password
		DB: redisConfig.Database
	}
}
