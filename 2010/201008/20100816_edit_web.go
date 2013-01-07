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

const lenPath = len("/Doc/");


func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[lenPath:];
    p, _ := loadPage(title);
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body);
}

func main() {
    http.HandleFunc("/Doc/", viewHandler);
    http.HandleFunc("/Debug/", editHandler);
    http.ListenAndServe(":8080", nil);
}

func editHandler(w http.ResponseWriter, r *http.Request){
   title := r.URL.Path[len("/Debug/"):];
   p, err := loadPage(title);
   if err != nil{
      p = &Page{Title: title};
   }
   fmt.Fprintf(w, "<h1>Editing %s</h1>" +
               "<form action=\"/save/%s\" method=\"POST\">" +
               "<textarea name=\"body\">%s</textarea><br>" +
               "<input type=\"submit\" value=\"Save\">" +
               "</form>",
               p.Title, p.Title, p.Body);
}
