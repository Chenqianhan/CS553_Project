package main

import (
	"log"
	"net/http"
	//"time"
	"html/template"
	"strings"
        "reflect"
	//"flag"
	//"io"
	//"os"
)

func main() {
/*
	server := &http.Server{
		Addr: ":4000",
		WriteTimeout: 2 * time.Second,
	}

	//Costomize Mux
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})
	//mux.Hadnle("/", indexHandler)
	//mux.Handle("/css/", http.FileServer(http.Dir("www")))
    //mux.Handle("/js/", http.FileServer(http.Dir("www")))
	server.Handler = mux

	log.Println("Server starts")
	log.Println("=============")
	log.Fatal(server.ListenAndServe())
*/

	log.Print("Server start...")
	http.Handle("/css/", http.FileServer(http.Dir("www")))
	http.Handle("/js/", http.FileServer(http.Dir("www")))
	//http.Handle("/download", &downloadHandler{})
	http.Handle("/download/", http.FileServer(http.Dir("www")))
	
	http.HandleFunc("/index", indexHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("indexHandler")
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    index := &indexController{}
    controller := reflect.ValueOf(index)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

//type indexHandler struct{}
type indexController struct{}


/*func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a Webpage from Go TinyServer"))
}
*/
/*
func (*indexHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("www/html/login/index.html")
	if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}
*/


func (*indexController)IndexAction(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("www/html/index.html")
	if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}

type downloadHandler struct{}



