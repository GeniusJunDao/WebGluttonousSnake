package game

import (
	"game/gs"
)

//Run 执行对战
func (m *Match) Run() {
	var s1, s2 gs.Jerry
	s1.SetBlock(gs.Weight/2, 0, 1)
	s1.Head = [2]int{gs.Weight / 2, 0}
	s1.Grown(2)
	s2.SetBlock(0, gs.Hight/2, 1)
	s2.Head = [2]int{0, gs.Hight / 2}
	s2.Grown(4)
	for {
		//未完待续
	}
}
