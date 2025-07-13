package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	godotenv.Load()
	config := &Config{
		RedisConfig: loadRedisConfig(),
		MongoConfig: loadMongoConfig(),
	}
	return config
}

func loadRedisConfig() RedisConfig {
	h, _ := GetEnv("REDIS_HOST", "localhost")
	p, _ := GetEnv("REDIS_PORT", 6379)
	pass, _ := GetEnv("REDIS_PASSWORD", "")
	db, _ := GetEnv("REDIS_DATABASE", 0)
	return RedisConfig{
		Host:     h,
		Port:     p,
		Password: pass,
		Database: db,
	}
}

func loadMongoConfig() MongoConfig {
	uri, _ := GetEnv("MONGODB_URI", "mongodb://localhost:27017")
	return MongoConfig{
		Uri: uri,
	}
}

func GetEnv[T string | int](key string, defaultVal T) (T, error) {
	rawEnv, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal, nil
	}
	var ret any
	switch any(defaultVal).(type) {
	case string:
		ret = rawEnv
	case int:
		intVal, err := strconv.Atoi(rawEnv)
		if err != nil {
			return defaultVal, err
		}
		ret = intVal

	}
	return ret.(T), nil
}
