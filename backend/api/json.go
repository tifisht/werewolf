package api

// api接受的json模板
type basicGameRequest struct {
	GameID     string `json:"game_id" binding:"required"`
	PlayerName string `json:"player_name" binding:"required"`
}

type setRoleRequest struct {
	GameID     string `json:"game_id" binding:"required"`
	PlayerName string `json:"player_name" binding:"required"`
	Seer       int    `json:"seer" binding:"required"`
	Witch      int    `json:"witch" binding:"required"`
	Wolf       int    `json:"wolf" binding:"required"`
	Hunter     int    `json:"hunter" binding:"required"`
	Villager   int    `json:"villager" binding:"required"`
}

type actionGameRequest struct {
	GameID     string `json:"game_id" binding:"required"`
	PlayerName string `json:"player_name" binding:"required"`
	Target     string `json:"targe" binding:"required"`
}
