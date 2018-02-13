package game

import (
	ws "github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"time"
)

var match chan *Player

//Player 表示一个访问的玩家
type Player struct {
	ID   uuid.UUID //玩家的ID
	conn *ws.Conn  //与玩家的连接
}

//Close 可以关闭Player的连接等
func (p *Player) Close() (err error) {
	err = p.conn.Close()
	if err != nil {
		return
	}
	return nil
}

//Match 表示一场对战
type Match struct {
	p1, p2 *Player
}

//初始化match管道
func init() {
	match = make(chan *Player)
}

//CreatePlayer 返回一个玩家UID
func CreatePlayer(conn *ws.Conn) Player {
	return Player{
		ID:   uuid.Must(uuid.NewV4()),
		conn: conn,
	}
}

//ServePlayer 为玩家WebSocket提供服务
func ServePlayer(player Player) {

	player.conn.WriteJSON(map[string]string{"yourID": player.ID.String()}) //告诉客户端玩家ID
	respond := map[string]string{}
	//读取客户端发送的"OK"
	isTimeOut, err := TimeOut(func() error {
		return player.conn.ReadJSON(respond)
	}, 30*time.Second)
	if isTimeOut {
		player.Close()
		return
	}
	if err != nil {
		player.Close()
		return
	}
	if respond["ready"] != "OK" {
		player.conn.WriteJSON(map[string]string{"msg": "close"})
		player.Close()
		return
	}
	player.searchRival()
}

//TimeOut 可以在d时间内执行f函数，如果超时则isTimeOut返回true，不过并不会结束执行f
func TimeOut(f func() error, d time.Duration) (isTimeOut bool, err error) {
	chin := make(chan error, 1)
	timer := time.NewTimer(d)

	go func() {
		err := f()
		chin <- err
	}()
	select {
	case err = <-chin:
		return false, err
	case <-timer.C:
		return true, nil
	}
}

//匹配队友并开始对战。该函数可能返回，也可能不返回
func (p *Player) searchRival() {
	select {
	case match <- p:
	case rival := <-match:
		m := Match{p1: p, p2: rival}
		m.Run()
	}
}
