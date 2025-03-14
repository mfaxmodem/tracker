package config

import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    APIPort    string
    RedisHost  string
    RedisPort  string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        APIPort:    os.Getenv("API_PORT"),
        RedisHost:  os.Getenv("REDIS_HOST"),
        RedisPort:  os.Getenv("REDIS_PORT"),
    }, nil
}

func (c *Config) GetDBConnString() string {
    return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}