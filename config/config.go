package config

type RedisConfig struct {
	Host string
	Port int
	Password string
	Database int
}

type Config struct {
	RedisConfig	
}
