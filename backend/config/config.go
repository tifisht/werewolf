package config

import (
	"os"
)

type Config struct {
	ServerPort  string
	MaxPlayers  int
	GameTimeout int // 游戏超时时间（分钟）
	Role        *Role
}

type Role struct {
	Witch    int
	Seer     int
	Villager int
	Hunter   int
	Wolf     int
}

func LoadConfig() *Config {
	role := &Role{
		Witch:    0,
		Seer:     0,
		Villager: 0,
		Hunter:   0,
		Wolf:     0,
	}
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		MaxPlayers:  10,
		GameTimeout: 60,
		Role:        role,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
