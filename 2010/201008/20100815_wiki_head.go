package main;
import "fmt";
import "net/http";
import "io/ioutil";


type Page struct{
  Title string;
  Body  []byte;
};

func  (p *Page) save() error{
   filename := p.Title + ".txt";
   return ioutil.WriteFile(filename, p.Body, 0600);
}

func loadPage(title string) (*Page, error){
   filename := title + ".txt";
   body, err := ioutil.ReadFile(filename);
   if err != nil {
      return nil, err;
   }
   return &Page{Title: title, Body: body},nil;
}

const lenPath = len("/mydoc/");


func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[lenPath:];
    p, _ := loadPage(title);
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body);
}

func main() {
    http.HandleFunc("/mydoc/", viewHandler);
    http.ListenAndServe(":8090", nil);
}
