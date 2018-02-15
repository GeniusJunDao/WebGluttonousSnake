package game

import (
	"fmt"
	"game/gs"
	"log"
	"time"
)

//Run 执行对战
func (m *Match) Run() {
	//初始化两条蛇
	var s1, s2 gs.Jerry
	s1.SetBlock(gs.Weight/2, 0, 1)
	s1.Head = [2]int{gs.Weight / 2, 0}
	s1.Grown(2)
	s2.SetBlock(gs.Weight/2, gs.Hight, 1)
	s2.Head = [2]int{gs.Weight / 2, gs.Hight}
	s2.Grown(4)
	SyncMatch := func(p *Player) error {
		err := p.conn.WriteJSON(map[string][gs.Weight][gs.Hight]int{"player1": s1.GetPlat(), "player2": s2.GetPlat()})
		return err
	}
	ReadD := func(p *Player) (int, error) {
		r := map[string]int{}
		err := p.conn.ReadJSON(&r)
		return r["d"], err
	}
	var d1, d2 = 4, 3
	var err1, err2 error
	go func() { //读取玩家1的操作
		for {
			d1, err1 = ReadD(m.p1)
			if err1 != nil {
				return
			}
		}
	}()
	go func() { //读取玩家2的操作
		for {
			d2, err2 = ReadD(m.p2)
			if err2 != nil {
				return
			}
		}
	}()
	ticker := time.NewTicker(1000 * time.Millisecond)
	for { //游戏进行时
		if err1 != nil || err2 != nil {
			log.Println("出错， 游戏退出", err1, err2)
			return
		}
		SyncMatch(m.p1) //与玩家1同步
		s1.Grown(d1)
		s1.Kick()
		SyncMatch(m.p2) //与玩家2同步
		s2.Grown(d2)
		s2.Kick()

		<-ticker.C
	}
}
func printSnake(sn gs.Jerry) {
	for i := 0; i < gs.Weight; i++ {
		s := ""
		for j := 0; j < gs.Hight; j++ {
			s += fmt.Sprintf("%d ", sn.GetBlock(i, j))
		}
		fmt.Println(s)
	}
	fmt.Println()
}
