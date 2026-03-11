package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Db MySQL
}

func Load() *Config {
	LoadEnv()

	return &Config{
		Db: loadMySQL(),
	}
}

func loadMySQL() MySQL {
	return MySQL{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Name:     os.Getenv("MYSQL_NAME"),
		MaxOpenConn: envInt("MYSQL_MAX_OPEN_CONNECT", 20),
		MaxIdleConn: envInt("MYSQL_MAX_IDLE_CONNECT", 10),
	}
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	i, err:=strconv.Atoi(v)
	if err != nil {
		log.Printf("invalid env %s=%v", key, v)
		return def
	}
	return i
}
