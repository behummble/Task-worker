package config

import (
	"fmt"
	"flag"
	"encoding/json"
)

type Config struct {
	Log LogConfig
	Server ServerConfig
	Worker WorkerConfig
	Jobs JobsConfig
	Redis RedisConfig
	Kafka KafkaConfig
}

type ServerConfig struct {
	Host string
	Port int
	ReadTimeout int
}

type LogConfig struct {
	OuptutPath string
	LogLevel int
}

type WorkerConfig struct {
	NumeberOfWorkers int
}

type JobsConfig struct {
	JobsType []string
	JobsParameters map[string]JobConfig
}

type JobConfig struct {
	Timeout int
	Retry int
	RateLimit int
	Priority int
}

type RedisConfig struct {
	Host string
	Port int
	Username string
	Password string
	ReadTimeout int
	WriteTimeout int
	DialTimeout int
	Retries int
}

type KafkaConfig struct {

}

func NewConfig(params string) Config {

}

func parseConfigFile(path string) (Config, error) {

}

func defaultConfig() Config {

}