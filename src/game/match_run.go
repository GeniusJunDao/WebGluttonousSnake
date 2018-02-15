package game

import (
	"fmt"
	"game/gs"
	"log"
	"math/rand"
	"time"
)

//Run 执行对战
func (m *Match) Run() {
	//初始化两条蛇
	var s1, s2 gs.Jerry
	s1.SetBlock(gs.Weight/2, 0, 1)
	s1.Head = [2]int{gs.Weight / 2, 0}
	s1.Grown(2)
	s2.SetBlock(gs.Weight/2, gs.Hight-1, 1)
	s2.Head = [2]int{gs.Weight / 2, gs.Hight - 1}
	s2.Grown(1)
	SyncMatch := func(p *Player, food [2]int) error {
		err := p.conn.WriteJSON(map[string]interface{}{"player1": s1.GetPlat(), "player2": s2.GetPlat(), "food": food})
		return err
	}
	ReadD := func(p *Player) (int, error) {
		r := map[string]int{}
		err := p.conn.ReadJSON(&r)
		return r["d"], err
	}
	newFood := func() [2]int {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var x, y int = r.Intn(gs.Weight), r.Intn(gs.Hight)
		for s1.GetBlock(x, y) != 0 || s2.GetBlock(x, y) != 0 {
			x, y = r.Intn(gs.Weight), r.Intn(gs.Hight)
		}
		return [2]int{x, y}
	}
	var d1, d2 = 0, 0
	var err1, err2 error
	go func() { //读取玩家1的操作
		for {
			d1, err1 = ReadD(m.p1)
			if err1 != nil {
				fmt.Println("d1 out")
				return
			}
		}
	}()
	go func() { //读取玩家2的操作
		for {
			d2, err2 = ReadD(m.p2)
			if err2 != nil {
				fmt.Println("d2 out")
				return
			}
		}
	}()
	ticker := time.NewTicker(200 * time.Millisecond)
	live1, live2 := true, true
	food := newFood()

	for { //游戏进行时
		if err1 != nil || err2 != nil {
			log.Println("出错， 游戏退出", err1, err2)
			return
		}
		if live1 {
			SyncMatch(m.p1, food) //与玩家1同步
			SyncMatch(m.p2, food) //与玩家2同步
			s1.Grown(d1)
			if s1.Head != food {
				s1.Kick()
			} else {
				food = newFood()
			}
			if s1.GetBlock(s1.Head[0], s1.Head[1]) > 1 && s2.GetBlock(s1.Head[0], s1.Head[1]) > 0 {
				live1 = false
			}
			//printSnake(s1)
		}
		<-ticker.C
		if live2 {
			SyncMatch(m.p1, food) //与玩家1同步
			SyncMatch(m.p2, food) //与玩家2同步
			s2.Grown(d2)
			if s2.Head != food {
				s2.Kick()
			} else {
				food = newFood()
			}
			if s1.GetBlock(s2.Head[0], s2.Head[1]) > 0 && s2.GetBlock(s2.Head[0], s2.Head[1]) > 1 {
				live2 = false
			}
			//printSnake(s2)
		}
		<-ticker.C

		if !live1 && !live2 {
			m.p1.conn.WriteJSON(map[string]string{"msg": "GAME OVER"})
			m.p2.conn.WriteJSON(map[string]string{"msg": "GAME OVER"})
			score1, score2 := s1.Score(), s2.Score()
			if score1 == score2 {
				m.p1.conn.WriteJSON(map[string]string{"final": "DRAW"})
				m.p2.conn.WriteJSON(map[string]string{"final": "DRAW"})
			} else if score1 > score2 {
				m.p1.conn.WriteJSON(map[string]string{"final": "WIN"})
				m.p2.conn.WriteJSON(map[string]string{"final": "LOSS"})
			} else if score2 > score1 {
				m.p1.conn.WriteJSON(map[string]string{"final": "LOSS"})
				m.p2.conn.WriteJSON(map[string]string{"final": "WIN"})
			}
		}
	}
}

//调试用绘制函数
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
