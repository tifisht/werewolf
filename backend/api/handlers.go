package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"werewolf-backend/internal/game"
	"werewolf-backend/internal/player"
)

var gm *game.GameManager

// 接口概览（（
func SetupRoutes(r *gin.Engine, gameManager *game.GameManager) {
	gm = gameManager
	// 游戏相关路由
	game := r.Group("/game")
	{
		game.POST("/create", createGame)
		game.POST("/join", joinGame)
		game.POST("/start", startGame)
		game.POST("/action", playerAction)
		game.POST("/setRole", setRole)
	}

	// 健康检查(服务器就绪)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	//每秒查询接口
	r.GET("/status", func(c *gin.Context) {
		var req basicGameRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		game, err := gm.GetGame(req.GameID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"game":   game,
			"you":    game.GetPlayer(req.PlayerName),
		})
	})
}

// http创建游戏接口
func createGame(c *gin.Context) {
	var req basicGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g, err := gm.CreateGame(req.GameID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := g.AddHost(req.PlayerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game_id":   g.ID,
		"player_id": player.ID,
	})
}

// 加入游戏
func joinGame(c *gin.Context) {
	var req basicGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g, err := gm.GetGame(req.GameID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	player, err := g.AddPlayer(req.PlayerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player_id": player.ID,
	})
}

// 开始游戏
func startGame(c *gin.Context) {
	gameID := c.Query("game_id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少game_id参数"})
		return
	}

	game, err := gm.GetGame(gameID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := game.Start(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "游戏开始",
	})
}

// 玩家行动
func playerAction(c *gin.Context) {
	// TODO: 实现玩家行动逻辑
	var req basicGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g, err := gm.GetGame(req.GameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	p := g.GetPlayer(req.PlayerName)

	switch p.Role {
	case player.Seer:

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "行动成功",
	})
}

func setRole(c *gin.Context) {
	var req setRoleRequest
	//解析json
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//检查是否已经创建
	game, err := gm.GetGame(req.GameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	game.Config.Role.Villager = req.Villager
	game.Config.Role.Seer = req.Seer
	game.Config.Role.Witch = req.Witch
	game.Config.Role.Hunter = req.Hunter
	game.Config.Role.Wolf = req.Wolf
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
