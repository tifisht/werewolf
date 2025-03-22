# http api
```api
/game
	/create
	/join
	/start
	/action
	/setRole

get:
/health
/status
```

/health 
返回json {"status":"ok"}

/status 
发送：json：{
	game_id,
	player_name
}

返回：json:{
	{
			"status":   "ok",
			"seer":     game.Seer,
			"witch":    game.Witch,
			"villager": game.Villager,
			"wolf":     game.Wolf,
			"hunter":   game.Hunter,
			"you":自己的信息
		}
}

/game:
	/createGame:json:{
		player_name,
		id
	}
	返回json:{
		game_id
		player_id
	}