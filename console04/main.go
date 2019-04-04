// vi:nu:et:sts=4 ts=4 sw=4


package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "sync"
    "github.com/gomodule/redigo/redis"
)


var hitmeTpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
    <p>OUCH - You have hit me {{.}} times!</p>
</body>
</html>`


//***************************************************************
//                      Redis Interface
//***************************************************************

func redisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	return c
}


func updateCount(cnt int) {

	c := redisConnect()
	defer c.Close()

	// set the value on redis for the key viewedcount
	reply, err := c.Do("SET", "count", cnt)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("GET ", reply)
}

func getCount() int {

	c := redisConnect()
	defer c.Close()
	// get the value from redis for the key viewed count
	reply, err := c.Do("GET", "count")
	if err != nil {
		log.Fatal(err)
	}
	if reply != nil {
		s := string(reply.([]byte))
		log.Println("GET ", s)
		i, _ := strconv.Atoi(s)
		return i
	}

	return 0
}



//***************************************************************
//                  HTTP Request Handlers
//***************************************************************

// Request Handlers run as independent goroutines so any shared
// data must be protected.

type httpHandler struct{
    mu      sync.Mutex
    count   int
}


func (h *httpHandler) base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached - %s\n", r.URL.Path)
}


func (h *httpHandler) hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! - %s\n", r.URL.Path)
}


func (h *httpHandler) hitme(w http.ResponseWriter, r *http.Request) {
    h.mu.Lock()
    defer h.mu.Unlock()
    h.count = getCount()
    h.count++
    updateCount(h.count)
    tpl, err := template.New("HitMe").Parse(hitmeTpl)
    if err != nil {
        log.Fatalln("Error while parsing template:", err)
    }
    err = tpl.ExecuteTemplate(w, "HitMe", h.count)
    if err != nil {
        log.Fatalln("Error while executing template:", err)
    }

}





//***************************************************************
//                          m a i n 
//***************************************************************

func main() {

    h := new(httpHandler)

    http.HandleFunc("/", h.base)
    http.HandleFunc("/hi", h.hi)
    http.HandleFunc("/hitme", h.hitme)

    log.Fatal(http.ListenAndServe(":8080", nil))

}
