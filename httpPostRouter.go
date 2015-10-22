package main
import ("fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type Response struct{
    Greeting string `json:"greeting"`
}

func PostHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    var data map[string]interface{}
    var response Response
    err:=json.NewDecoder(req.Body).Decode(&data)
    name:=data["name"]
    if(err==nil){
        response.Greeting ="Hello, "+ name.(string)+"!"
        rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
        result,err:=json.Marshal(response)
        if(err!=nil){
            http.Error(rw, err.Error(), 500)
        }
        fmt.Fprintln(rw,string(result))
    }else{
         http.Error(rw, err.Error(), 500)
    }
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func main() {
    mux := httprouter.New()
    mux.POST("/hello",PostHello)
    mux.GET("/hello/:name", hello)
     server := http.Server{
            Addr:        "0.0.0.0:8088",
            Handler: mux,
    }
    server.ListenAndServe()
}