function InitSharders(gl, vertexSharderId, fragmentSharderId){
	var vertShdr;
	var fragShdr;
	
	var vertElem = document.getElementById(vertexSharderId);
	if (!vertElem) {
		alert("不能加载顶点着色器	" + vertexSharderId);
	}
	else {
		vertShdr = gl.createShader(gl.VERTEX_SHADER);
		gl.shaderSource(vertShdr, vertElem.text);
		gl.compileShader(vertShdr);
		if (!gl.getShaderParameter(vertShdr, gl.COMPILE_STATUS)) {
			var msg = "顶点着色器编译失败！"
				+ "错误日志： "
				+ "<pre>" + gl.getProgramInfoLog(vertShdr) + "</pre>";
			alert(msg);
			return -1;
		}
	}
}
