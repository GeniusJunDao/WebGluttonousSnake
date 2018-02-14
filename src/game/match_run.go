package game

import (
	"game/gs"
)

//Run 执行对战
func (m *Match) Run() {
	//初始化两条蛇
	var s1, s2 gs.Jerry
	s1.SetBlock(gs.Weight/2, 0, 1)
	s1.Head = [2]int{gs.Weight / 2, 0}
	s1.Grown(2)
	s2.SetBlock(0, gs.Hight/2, 1)
	s2.Head = [2]int{0, gs.Hight / 2}
	s2.Grown(4)
	SyncMatch := func() error {
		err := m.p1.conn.WriteJSON(map[string][gs.Weight][gs.Hight]int{"player1": s1.GetPlat(), "player2": s2.GetPlat()})
		if err != nil {
			return err
		}
		err = m.p2.conn.WriteJSON(map[string][gs.Weight][gs.Hight]int{"player1": s1.GetPlat(), "player2": s2.GetPlat()})
		if err != nil {
			return err
		}
		return nil
	}
	ReadD := func(p *Player) (int, error) {
		r := map[string]int{}
		err := p.conn.ReadJSON(r)
		return r["d"], err
	}
	for {
		//未完待续
		SyncMatch()
		d1, err := ReadD(m.p1)
		if err != nil {
			return
		}
		s1.Grown(d1)
		s1.Kick()
		SyncMatch()
		d2, err := ReadD(m.p1)
		if err != nil {
			return
		}
		s2.Grown(d2)
		s2.Kick()
	}
}
