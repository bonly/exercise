package main 

import (
"log"
"net/http"
"flag"
)


var head string =`
<!DOCTYPE html>
<html>
<head>
	<title>vue test</title>
	<script src="/lib/vue.min.js"></script>

</head>
`;

var foot string = `
</html>
`;

var body string = `
<body>

<div id="app">
{{message}}
</div>

</body>
`;


// var app string = `
// <script type="text/javascript">
// 	new Vue({
// 		el : '#app',
// 		data : {
// 			message : 'hello for vue test'
// 		}
// 	});

// </script>
// `;

var app string = `
<script src="/lib/msg.js"></script>
`;

var srv string;

func init(){
  flag.StringVar(&srv, "s", "0.0.0.0:9997", "srv addr"); 
}

func reg_handle(){
	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir("./"))));
	http.HandleFunc("/hm", complete_task_func);
	// templates, err := template.ParseGlob("*.html");
	// if err != nil{
	// 	log.Printf(err);
	// }
}

func main(){
	reg_handle();
	log.Printf("srv: %s", srv);
	log.Fatal(http.ListenAndServe(srv, nil));
}

func complete_task_func(wr http.ResponseWriter, req *http.Request){
	wr.Write([]byte(head + body + app + foot));
}