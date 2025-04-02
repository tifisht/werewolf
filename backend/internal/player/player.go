package player

type Role string

const (
	Werewolf Role = "werewolf"
	Villager Role = "villager"
	Seer     Role = "seer"
	Witch    Role = "witch"
	Hunter   Role = "hunter"
)

type Player struct {
	Host   bool //是否是房主
	ID     string
	Name   string //接口获取
	Role   Role   //职业
	Alive  bool   //存活
	Voted  bool   //是否投票
	VoteTo string //投票对象
}
