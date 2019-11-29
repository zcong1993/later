package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zcong1993/later/queue"
)

var (
	redisURLEnv = "REDIS_URL"
	addressEnv  = "ADDRESS"
	redisURL    = flag.String("redis", "redis://127.0.0.1:6379/0", "redis address")
	address     = flag.String("address", ":8080", "serve listen address")
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	flag.Parse()
}

func loadEnvOrDefault(envName string, defaultValue string) string {
	envVal := os.Getenv(envName)
	if envVal == "" {
		return defaultValue
	}
	return envVal
}

func main() {
	redisURLVal := loadEnvOrDefault(redisURLEnv, *redisURL)
	addressVal := loadEnvOrDefault(addressEnv, *address)
	err := queue.InitRedis(redisURLVal)
	if err != nil {
		log.Fatal(err)
	}
	queue.RunWorker()
	log.Infof("server listen on :%v", *address)
	err = queue.ListenAndServe(addressVal)
	if err != nil {
		log.Fatal(err)
	}
}
