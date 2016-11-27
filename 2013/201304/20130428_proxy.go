package main

import(
	"net/http"
	"log"
	"os"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request){
	resp, err := http.DefaultClient.Do(r);
	defer resp.Body.Close();
	if err != nil{
		panic(err);
	}
	for k, v := range resp.Header{
		for _, vv := range v{
			w.Header().Add(k, vv);
		}
	}
	for _, c := range resp.SetCookie{
		w.Header().Add("Set-Cookie", c.Raw);
	}
	w.WriteHeader(resp.StatusCode);
	result, err := ioutil.ReadAll(resp.Body);
	if err != nil && err != os.EOF{
		panic(err);
	}
	w.Write(result);
}

func main(){
	http.HandleFunc("/", handler);
	log.Println("Start serving on port 8888");
	http.ListenAndServ(":8888", nil);
	os.Exit(0);
}