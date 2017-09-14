
	//全局变量
	var wsUri = "ws://127.0.0.1:9998/Main";
	var output; //显示用控件 

	//初始化
	function init(){
		output = document.getElementById("output");
		setup_WebSocket();
	}

	//设置连接的响应函数
	function setup_WebSocket(){
		websocket = new WebSocket(wsUri);
		websocket.onopen = function(evt) { onOpen(evt) };
		websocket.onclose = function(evt) { onClose(evt) };
		websocket.onmessage = function(evt) { onMessage(evt) };
		websocket.onerror = function(evt) { onError(evt) };
	}

	//连接打开处理
	function onOpen(evt) {
		writeToScreen('<span class="label label-success">' + "CONNECTED" +'</span>');
	}
	//连接断开处理
	function onClose(evt) { writeToScreen("DISCONNECTED"); }
	//连接错误处理
	function onError(evt) {
	  writeToScreen('<span style="color: red;">ERROR:</span> '+ evt.data); 
	}

	//发送信息
	function doSend(message) { 
	    writeToScreen("SENT: " + message);  
	    websocket.send(message); 
	}  

	//显示控件
	function writeToScreen(message) { 
	    var pre = document.createElement("p"); 
	    pre.style.wordWrap = "break-word"; 
	    pre.innerHTML = message; 
	    output.appendChild(pre); 
	}

	function writeToTable(msg){
		var tab = document.getElementById("table");
		var tr = document.createElement("tr");
		var td = document.createElement("td");

		td.appendChild(document.createTextNode("abc"));
		tr.appendChild(td);
		
		td.appendChild(document.createTextNode("adfd"));
		tr.appendChild(td);
		// tab.appendChild(td);
		// var td = document.createElement("td").Value = "ddfe";
		// tab.appendChild(td);

		tab.appendChild(tr);
	}

	//接收信息处理
	function onMessage(evt) { 
	    writeToScreen('<span style="color: blue;">RESPONSE: '+ evt.data +'</span>'); 

	    cmd = JSON.parse(evt.data);
	    switch (cmd.Cmd){
	      case "RQuery":
	        delete cmd.Cmd; //删除命令码
	        if (cmd.Ret=="0"){
        		writeToScreen('<span style="color:green;">INFO: '+ '查询成功' + ':' + cmd.Msg + '</span>'); 
        		writeToTable(cmd);
        	}else{
        		writeToScreen('<span style="color:red;">INFO: '+ '查询失败' + ':' + cmd.Msg + '</span>'); 
        	}
	        break;	                   
	      default:
	        writeToScreen('<span style="color:green;">INFO: 返回 '+ cmd.Cmd + '</span>'); 
	        break;
	    }
	}  

	window.addEventListener("load", init, false); 

//=================== 处理函数

	function query(){
		var oms_order_id = document.getElementById("oms_order_id_text").value;

		var head={"Cmd":"Query"};
		doSend(JSON.stringify(head));

		var call_query={"Oms_order_id":oms_order_id};
		doSend(JSON.stringify(call_query));
	}

	function prc_query(cmd){
		for (j = 0; j < cmd.Ver.length; j++){
    		var new_option = new Option(cmd.Ver[j].Tag, cmd.Ver[j].Tag); //定义新选择
			version.options.add(new_option); //加入选择
		}
	}

