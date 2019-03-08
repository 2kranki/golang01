package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)




//***************************************************************
//                  HTTP Request Handlers
//***************************************************************

type httpHandler struct{
    count   int
}


func (h *httpHandler) base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached - %s\n", r.URL.Path)
}


func (h *httpHandler) hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! - %s\n", r.URL.Path)
}


func (h *httpHandler) hitme(w http.ResponseWriter, r *http.Request) {
    h.count++
	fmt.Fprintf(w, "Ouch! %d - %s\n", h.count, r.URL.Path)
}


func (h *httpHandler) quit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye!\n")
    os.Exit(0)
}





//***************************************************************
//                          m a i n 
//***************************************************************

func main() {

    h := new(httpHandler)

    http.HandleFunc("/", h.base)
    http.HandleFunc("/hi", h.hi)
    http.HandleFunc("/hitme", h.hitme)
    http.HandleFunc("/quit", h.quit)

    log.Fatal(http.ListenAndServe("localhost:9000", nil))

}
