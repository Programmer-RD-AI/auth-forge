package config

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

type MongoConfig struct {
	Uri string
}

type Config struct {
	RedisConfig
	MongoConfig
}
