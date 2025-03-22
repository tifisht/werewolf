package game

import "werewolf-backend/internal/player"

func (g *Game) gameLoop() {

	for isOver(g) {

		// nightStart(g)

		// dayStart(g)

		// speeckStart(g)

		// voteStart(g)

	}
}

// 主要游戏阶段
func nightStart(g *Game) {
	//预言家 ——> 狼人 ——> 女巫 ——>if hunter死了hunter开枪——>天亮
	// seerAction(g)

	// wolfAction(g)

	// witchAction(g)

	// //内部条件判断条件
	// hunterAction(g)

}

func seerAction(g *Game) {

}

func ifGameOver(g *Game) {

}

// 判断游戏是否结束
func isOver(g *Game) bool {
	if getWolfNum(g) > getGoodNum(g) {
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
