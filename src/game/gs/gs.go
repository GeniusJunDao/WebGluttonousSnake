package gs

const (
	Weight = 16
	Hight  = 9
)

//Direction 表示了单个方向d值如何分解至笛卡尔坐标系，其中Direction[0]并不应该被使用
var Direction = [5][2]int{
	[2]int{0, 0},
	[2]int{0, -1},
	[2]int{0, 1},
	[2]int{-1, 0},
	[2]int{1, 0},
}

//Jerry 表示一条蛇(不要问为什么叫Jerry，Jerry是只老鼠)
type Jerry struct {
	plat [Weight][Hight]int //地图
	Head [2]int             //蛇头所在位置
	d    int                //蛇目前行进的方向
}

//GetPlat 返回整个地图
func (j *Jerry) GetPlat() [Weight][Hight]int {
	return j.plat
}

//GetBlock 获取地图上一个块所储存的数据，0表示空，>0表示是蛇， <0表示是食物，参数x, y是坐标返回值是数据
func (j *Jerry) GetBlock(x, y int) int {
	x, y = FormattingCoordinates(x, y)
	return j.plat[x][y]
}

//SetBlock 写入地图上一个块的数据，0表示空，>0表示是蛇， <0表示是食物，参数x, y是坐标，b是要写入的数据
func (j *Jerry) SetBlock(x, y, b int) {
	x, y = FormattingCoordinates(x, y)
	j.plat[x][y] = b
}

//FormattingCoordinates 函数对坐标进行格式化，使x值落在[0, Weight)内，使y值落在[0, Hight)内
func FormattingCoordinates(x, y int) (int, int) {
	for x < 0 {
		x += Weight
	}
	for y < 0 {
		y += Hight
	}
	x %= Weight
	y %= Hight
	return x, y
}

//Grown 使蛇向d所指示的方向伸长，d的取值为1up, 2down, 3left, 4right。返回值是蛇吃到的块的数据值，若没吃到东西，返回0
func (j *Jerry) Grown(d int) (b int) {
	if d > 0 && d < 5 {
		j.d = d
	}
	v := j.GetBlock(j.Head[0], j.Head[1]) //蛇头值
	x, y := FormattingCoordinates(j.Head[0]+Direction[d][0],
		j.Head[1]+Direction[d][1])
	b = j.GetBlock(x, y)
	j.SetBlock(x, y, v+1)
	return
}

//Kick 使蛇尾去掉一格
func (j *Jerry) Kick() {
	for i := 0; i < Weight; i++ {
		for ii := 0; ii < Hight; ii++ {
			b := j.GetBlock(i, ii)
			if b > 0 {
				j.SetBlock(i, ii, b-1)
			}
		}
	}
}
