// this method return a function that any time is called (with http.Request and ResponseWriter)
// will create a new instance with the same type as "sample" parameter
// and invoke the "method"
func RestDispatchMeth(sample interface{}, method string) http.HandlerFunc {
       t := reflect.ValueOf(sample).Type()
       return func(w http.ResponseWriter, req *http.Request) {
               v := reflect.New(t)
               log.Printf("Controller: %v", t.Name())

               callControllerMethod(method, v, w, req)
       }
}

// this function receiveis the controler "cont" and the "methodName" parameters
// together with the ResponseWriter and Request
// If it is able to find the method it call's the method.
func callControllerMethod(methodName string, cont reflect.Value, w
http.ResponseWriter, req *http.Request) {
       // check if the method is valid
       if method := cont.MethodByName(methodName); method.IsValid() {
               vw := reflect.ValueOf(w)
               vreq := reflect.ValueOf(req)
               log.Printf("Method: %v", methodName)
               // call it with the rigth parameters.
               method.Call([]reflect.Value{0: vw, 1: vreq})

               // Calls the after request handler
               // If one is defined
               if AfterRequest != nil {
                       AfterRequest.ServeHTTP(w, req)
               }

       } else {
               log.Printf("Unable to find method: %v", methodName)


       }
}

/*
    params := make([]reflect.Value, 3);
    params[0] = reflect.ValueOf(&args);
    params[1] = reflect.ValueOf(&reply);
    params[2] = reflect.ValueOf(&ret);

    defer func(){
      if ex := recover(); ex!=nil{
        logger.Error("no this event: ", args.Event);
      }
    }();

    (func(){
       reflect.ValueOf(&st{}).MethodByName(args.Event).Call(params);    
    })();
*/


var PMS_Type = make(map[string]reflect.Type);

func init(){
  PMS_Type["OpenLock_REQ"] = reflect.TypeOf(OpenLock_REQ{});
}

func makeInstance(name string)(interface{}){
   return reflect.New(PMS_Type[name]).Elem().Interface();
}