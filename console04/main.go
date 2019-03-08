// vi:nu:et:sts=4 ts=4 sw=4


package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gomodule/redigo/redis"
)




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
	// get the value from redis for the key viewedcount
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
    count := getCount()
	fmt.Fprintf(w, "Ouch! %d - %s\n", count, r.URL.Path)
    count++
    updateCount(count)
}





//***************************************************************
//                          m a i n 
//***************************************************************

func main() {

    h := new(httpHandler)

    http.HandleFunc("/", h.base)
    http.HandleFunc("/hi", h.hi)
    http.HandleFunc("/hitme", h.hitme)

    log.Fatal(http.ListenAndServe(":9000", nil))

}
