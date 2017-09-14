package main 

import(
"fmt"
"net/http"
"log"
)

import _ "net/http/pprof" 

type tracker struct{
	r *http.Request
};

func (t *tracker) Tracker(){
	fmt.Println("init");
}

func (t *tracker) Close(){
	fmt.Println("exit");
}

func main(){
	reqs := make(chan *http.Request, 1000);
	http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request){
		reqs <- r;
		fmt.Fprintf(w, "ok\n");
	});

	go trackerRequests(reqs);
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil));
}

func trackerRequests(reqs chan *http.Request){
	for {
		select {
		case r := <- reqs:
			tracker := &tracker{r};
			defer tracker.Close();

			tracker.Tracker();
		}
	}
}