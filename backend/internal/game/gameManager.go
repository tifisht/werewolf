package game

import (
	"errors"
	"sync"
	"werewolf-backend/config"
)

type GameManager struct {
	Games  map[string]*Game
	Mutex  sync.Mutex
	Config *config.Config
}

// 主函数调用产生新的gamemanager
func NewGameManager(config *config.Config) *GameManager {
	return &GameManager{
		//用map管理游戏
		Games:  make(map[string]*Game),
		Config: config,
	}
}

func (gm *GameManager) CreateGame(id string) (*Game, error) {
	//此步骤会上锁，必须放在上锁前
	_, err := gm.GetGame(id)

	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	if err == nil {
		return nil, errors.New("已有同名房间！")
	}

	game := NewGame(gm.Config, id)
	gm.Games[game.ID] = game
	return game, nil
}

func (gm *GameManager) GetGame(gameID string) (*Game, error) {
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	game, exists := gm.Games[gameID]
	if !exists {
		return nil, errors.New("游戏不存在")
	}
	return game, nil
}
