package main 

import (
"log"
"net/http"
"flag"
"text/template"
)


var srv string;

func init(){
    flag.StringVar(&srv, "s", "0.0.0.0:9997", "srv addr"); 
}

func reg_handle(){
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("public/"))));
	http.Handle("/data/", http.FileServer(http.Dir("./")));
	http.HandleFunc("/", view);
}

func main(){
	reg_handle();
	log.Printf("srv: %s", srv);
	log.Fatal(http.ListenAndServe(srv, nil));
}


type Val map[string]interface{};

func get_pic(param string)string{
	log.Printf("in get_pic()\n");
	return "this s ok";
}

func view(wr http.ResponseWriter, req *http.Request){
	var err error;
	tpl := template.Must(template.ParseGlob("tpl/*.html"));

	tp := tpl.Lookup("index.html");
	tp = tp.Funcs(template.FuncMap{"get_pic":get_pic});
	tp, err = tp.Parse(`
{{ define "代码" }}
<p> ID {{ .ID }} </p>
{{ end }}
`);
	if err != nil{
		log.Fatal(err);
	}


//{{ with $pic := ("ok" | get_pic)}}{{$pic}} {{ end}}
	tp, err = tp.Parse(`
{{ define "图表" }}
{{ "none" | get_pic}}
{{ end }}
`);
	if err != nil{
		log.Fatal(err);
	}

	if err := tp.Execute(wr, map[string]string{"ID":"One"}); err != nil{
		log.Fatal(err);
	}

	// if err := tpl.ExecuteTemplate(wr, "index.html", Val{"Title": "Xbed"}); err != nil{
	// 	log.Fatal(err);
	// }

}

