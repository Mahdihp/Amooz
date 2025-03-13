package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	AppServer AppServer
	Postgres  Postgres
	CronJob   CronJob
}
type AppServer struct {
	Host string
	Port int
}
type Postgres struct {
	Host     string
	DbName   string
	User     string
	Pass     string
	Port     int
	SSLMode  string
	MaxConns int
	MaxIdle  int
}
type CronJob struct {
	CronJobIntervalCoinInfo         int
	CronJobIntervalFearGreed        int
	CronJobIntervalCoinMarketCapMap int
	CronJobIntervalListingLatestCMC int
	CronJobIntervalUpdatePairSpot   int
	CronJobIntervalUpdateBaseCoin   int
}
type MongodbConfig struct {
	Host        string
	DbName      string
	BybitDbName string
	User        string
	Pass        string
	Port        int
	MaxPoolSize int
}
type RedisConfig struct {
	Host   string
	DbName string
	User   string
	Pass   string
	Port   int
}

func LoadConfig() Config {

	var c Config

	c.AppServer.Host = os.Getenv("APP_HOST")
	numberCfg, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	c.AppServer.Port = numberCfg

	return c
}

func GetConfig(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
