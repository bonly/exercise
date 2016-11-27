package main

import (
"fmt"
"net/http"
"syscall"
"os"
"os/signal"
"log"
)

func main(){
   sigs := make(chan os.Signal, 1);
   signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);
   
   go srv();
   
   <-sigs;
   fmt.Println("\n recv end sigs\n");
}

func srv(){
    http.Handle("/",http.FileServer(http.Dir(".")));
    http.HandleFunc("/index.html", index);
    err := http.ListenAndServe(":8888", nil);
    if err != nil{
        log.Fatal(err);
    }
}

func index(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, home);
}

var home = `
<!doctype html>
<html ng-app="FilterDemo">
<head>
  <title>Google phone gallery</title>
  <script src="angular.min.js"></script>

  <script>
  var App = angular.module('FilterDemo', []);
  
  App.controller('PhoneCtl', function($scope){
      $scope.phones = [
          {'name':'Nexus S',
              'snippet': 'Fast just got faster with Nexus S.'},
          {'name':'Motorola XOOM with wi-fi',
              'snippet': 'The Next, Next Generation tablet.'},
          {'name':'Motorola XOOM',
              'snippet':'The Next, Next Generation tablet.'}
      ];
  });
  </script>
</head>


<body ng-controller="PhoneCtl">

<div>
  <div class="row">
     
     <div>
        Search: <input ng-model="query">
     </div>
     
     <div>
       <ul class="phones">
         <li ng-repeat="phone in phones | filter:query">
           <span>{{phone.name}}</span>
           <p> {{phone.snippet}} </p>
         </li>
       </ul>
    </div>
  </div>
</div>

</body>
  

</html>

`;