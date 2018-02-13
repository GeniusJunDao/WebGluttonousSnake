var ID//玩家ID
function game(){
	if ("WebSocket" in window)
	{
	var Socket = new WebSocket("ws://"+window.location.host+"/game/websocket")
	Socket.onopen = function()
	{
		
		Socket.onmessage=function(evt){
			ID=evt
			Socket.send(JSON.stringify({"ready":"OK"}))
		}
	}
	}
	else
	{
		alert("浏览器不支持WebSocket！")
	}
}
