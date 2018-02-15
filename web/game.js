var ID //玩家ID
var d = 0

function game() {
	window.addEventListener('keydown', function(e) {
		var x = e.key;
		if(x == "Up") {
			d = 1
		} else if(x == "Down") {
			d = 2
		} else if(x == "Left") {
			d = 3
		} else if(x == "right") {
			d = 4
		}
	})

	if("WebSocket" in window) {
		var Socket = new WebSocket("ws://" + window.location.host + "/game/websocket")
		Socket.onopen = function() {
			Socket.onmessage = function(evt) {
				var ID = JSON.parse(evt.data);
				console.log("你的ID： " + ID.yourID);
				Socket.onmessage = function(evt) {
					paintSnake(JSON.parse(evt.data));
					Socket.send(JSON.stringify({
						"d": d
					}));
				}
				Socket.send(JSON.stringify({
					"ready": "OK"
				}));
			}
		}
	} else {
		alert("浏览器不支持WebSocket！");
	}
}

function paintSnake(a) {
	//console.log(snake.player1)
	//console.log(snake.player2)
	var c = document.getElementById("gs-canvas");
	drawSnake(a.player1, c, "#FF0000")
	drawSnake(a.player2, c, "#0000FF")
}

function drawSnake(snake, c, color) {
	console.log(snake)
	var cxt = c.getContext("2d");
	cxt.beginPath()
	
	var wight = c.width;
	var height = c.height;
	var wSnake = snake.length;
	var hSnake = snake[0].length;
	var stepW = wight/wSnake;
	var stepH = height/hSnake;
	console.log(wight, height, wSnake, hSnake, stepW, stepH)
	for (i=0; i<wSnake;i++) {
		for (j=0; j<hSnake;j++) {
			if (snake[i][j]==0){
				cxt.fillStyle = "#DDDDDD"
			} else if (snake[i][j]>0) {
				cxt.fillStyle=color
			} else {
				cxt.fillStyle="#00FF00"
			}
			cxt.beginPath()
			cxt.arc(stepW/2+i*stepW,stepH/2+j*stepH,7,0,Math.PI*2,true)
			cxt.stroke()
			cxt.closePath()
			cxt.fill()
		}
	}
}
