var ID //玩家ID
var d = 1

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

function paintSnake(snake) {
	console.log(snake.player1)
	//console.log(snake.player2)
	var c = document.getElementById("gs-canvas");
	var cxt = c.getContext("2d");
	cxt.fillStyle = "#DDDDDD"

	cxt.fill()
}