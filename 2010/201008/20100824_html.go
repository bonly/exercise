package main;
import _ "fmt";
import "net/http";
import "html/template";
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

const lenPath = len("/");

func main() {
    http.HandleFunc("/", viewHandler);
    http.ListenAndServe(":8090", nil);
}

func viewHandler(w http.ResponseWriter, r *http.Request){
   title := r.URL.Path[len("/"):];
   p, err := loadPage(title);
   if err != nil{
      p = &Page{Title: title};
   }
   t, _ := template.ParseFiles("20100825_html.html");
   err = t.Execute(w, p);
   if err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError);
   }
}
