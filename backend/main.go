package main

import (
	"werewolf-backend/api"
	"werewolf-backend/config"
	"werewolf-backend/internal/game"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化游戏管理器
	GameManager := game.NewGameManager(cfg)

	// 初始化Gin引擎
	r := gin.Default()

	// 设置API路由
	api.SetupRoutes(r, GameManager)

	// 启动服务器
	r.Run(":" + cfg.ServerPort)
}
