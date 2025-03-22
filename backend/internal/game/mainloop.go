package game

import (
	"time"
	"werewolf-backend/internal/player"
)

func (g *Game) gameLoop() {

	for isOver(g) {

		nightStart(g)

		dayStart(g)

		speechStart(g)

		voteStart(g)

	}
}

// 主要游戏阶段
func nightStart(g *Game) {
	g.DayNight = "night"
	//预言家 ——> 狼人 ——> 女巫 ——>if hunter死了hunter开枪——>天亮
	seerAction(g)

	wolfAction(g)

	witchAction(g)

	for _, p := range g.Players {
		if p.Role == player.Hunter {
			if !p.Alive {
				hunterAction(g)
			}
		}
	}
}

func seerAction(g *Game) {
	g.Stage = "seerAction"
	select {
	case msg := <-g.Msg: // 从 ch 读取数据
		p := g.GetPlayer(msg)
		if p.Role == player.Werewolf {
			g.Msg <- "bad"
		} else {
			g.Msg <- "good"
		}
	case <-time.After(15 * time.Second): // 设置超时时间为 15秒
	}
}

func wolfAction(g *Game) {
	g.Stage = "wolfAction"
	select {
	case msg := <-g.Msg: // 从 ch 读取数据
		p := g.GetPlayer(msg)
		p.Alive = false
		g.Killed = p.Name
	case <-time.After(90 * time.Second):
	}
}

func witchAction(g *Game) {
	g.Stage = "witchAction"
	action := <-g.Msg
	select {
	case msg := <-g.Msg: // 从 ch 读取数据
		if action == "save" {
			p := g.GetPlayer(g.Killed)
			p.Alive = true
			g.Killed = ""
		} else if action == "poison" {
			p := g.GetPlayer(msg)
			p.Alive = false
		}
	case <-time.After(30 * time.Second):
	}
}

func hunterAction(g *Game) {
	g.Stage = "hunterAction"
	select {
	case msg := <-g.Msg: // 从 ch 读取数据
		p := g.GetPlayer(msg)
		p.Alive = false
		g.Killed = p.Name
	case <-time.After(30 * time.Second):
	}
}

// 判断游戏是否结束
func isOver(g *Game) bool {
	if getWolfNum(g) > getGoodNum(g) || getWolfNum(g) == 0 {
		//游戏结束
		return false
	} else {
		return true
	}
}

// 其他实现函数
func getWolfNum(g *Game) int {
	var count int = 0
	for _, p := range g.Players {
		if p.Alive && p.Role == player.Werewolf {
			count++
		}
	}
	return count
}

func getGoodNum(g *Game) int {
	var count int = 0
	for _, p := range g.Players {
		if p.Alive && p.Role != player.Werewolf {
			count++
		}
	}
	return count
}

func dayStart(g *Game) {
	g.DayNight = "day"
	g.DayCount++
}

func speechStart(g *Game) {
	g.Stage = "speech"
	time.Sleep(5 * time.Minute)
}

func voteStart(g *Game) {

}
