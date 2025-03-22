// 游戏基础数据结构

package game

import (
	"errors"
	"math/rand"
	"sync"
	"time"
	"werewolf-backend/config"

	"werewolf-backend/internal/player"
)

type Game struct {
	ID       string
	Players  []*player.Player
	Started  bool
	DayNight string // "day" or "night"
	Stage    string
	DayCount int
	Mutex    sync.Mutex
	Config   *config.Config
	Timeout  *time.Timer
	Msg      chan string
	Killed   string
}

func (g *Game) GetPlayer(playerName string) *player.Player {
	for _, player := range g.Players {
		if player.Name == playerName {
			return player
		}
	}
	return nil
}

func NewGame(config *config.Config, id string) *Game {
	return &Game{
		ID:       id,
		Players:  make([]*player.Player, 0, config.MaxPlayers),
		Started:  false,
		DayNight: "day",
		DayCount: 1,
		Config:   config,
	}
}

func (g *Game) AddHost(name string) (*player.Player, error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	player := &player.Player{
		Host:  true,
		ID:    generatePlayerID(),
		Name:  name,
		Role:  player.Villager,
		Alive: true,
		Voted: false,
	}

	g.Players = append(g.Players, player)
	return player, nil
}

// 给game添加玩家
func (g *Game) AddPlayer(name string) (*player.Player, error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if len(g.Players) >= g.Config.MaxPlayers {
		return nil, errors.New("游戏人数已满")
	}

	for _, player := range g.Players {
		if player == g.GetPlayer(name) {
			return nil, errors.New("重名玩家已加入！")
		}
	}

	player := &player.Player{
		Host:  false,
		ID:    generatePlayerID(),
		Name:  name,
		Role:  player.Villager,
		Alive: true,
		Voted: false,
	}

	g.Players = append(g.Players, player)
	return player, nil
}

func (g *Game) Start() error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if len(g.Players) < 5 {
		return errors.New("至少需要5名玩家才能开始游戏")
	}

	g.Started = true
	// g.Timeout = time.NewTimer(time.Duration(g.Config.GameTimeout) * time.Minute)
	err := g.assignRole()
	if err != nil {
		return err
	}
	go g.gameLoop()
	return nil
}

// 其他游戏逻辑方法...

func generatePlayerID() string {
	return "player_" + time.Now().Format("20060102150405")
}

func (g *Game) assignRole() error {
	// 初始化随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 创建角色列表
	roleList := make([]player.Role, 0)
	for j := 0; j < g.Config.Role.Witch; j++ {
		roleList = append(roleList, player.Witch)
	}
	for j := 0; j < g.Config.Role.Seer; j++ {
		roleList = append(roleList, player.Seer)
	}
	for j := 0; j < g.Config.Role.Hunter; j++ {
		roleList = append(roleList, player.Hunter)
	}
	for j := 0; j < g.Config.Role.Wolf; j++ {
		roleList = append(roleList, player.Werewolf)
	}
	for j := 0; j < g.Config.Role.Villager; j++ {
		roleList = append(roleList, player.Villager)
	}
	if len(g.Players) != len(roleList) {
		return errors.New("请设置好游戏角色数量！")
	}
	// 随机打乱角色顺序
	r.Shuffle(len(roleList), func(i, j int) {
		roleList[i], roleList[j] = roleList[j], roleList[i]
	})

	// 将角色分配给玩家
	for i, p := range g.Players {
		if i < len(roleList) {
			p.Role = roleList[i]
		} else {
			p.Role = player.Villager // 默认角色
		}
	}
	return nil
}
