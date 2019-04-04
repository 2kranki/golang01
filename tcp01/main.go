// vi:nu:et:sts=4 ts=4 sw=4


package main

import (
    "bufio"
    "bytes"
    "fmt"
    "html/template"
    "log"
    "net"
    "strings"
)


var foundmeText = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Found Me</title>
</head>
<body>
    <p>I see that you found me!</p>
    <p>However, I see you  at {{.}}!</p>
</body>
</html>`



func lineToWords(l string) []string {
        var words []string

	scanner := bufio.NewScanner(strings.NewReader(l))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Split the line.
	for scanner.Scan() {
        word := scanner.Text()
        //fmt.Println("'",word,"'")
        words = append(words, word) 
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error splitting words:", err)
	}

    return words
}



//***************************************************************
//                      TCP Request Handlers
//***************************************************************

// Request Handlers run as independent goroutines so any shared
// data must be protected.

func baseHandler(c net.Conn) {
    var httpHeader []string
    var HTTPURL string

    defer c.Close()

    count := 0
    scanner := bufio.NewScanner(c)
    for scanner.Scan() {
        l := scanner.Text()
        if len(l) == 0 {
            break
        }
        if count == 0 {
            httpHeader = lineToWords(l)
        }
        if count == 1 {
            urls := lineToWords(l)
            HTTPURL = urls[1]
        }
       fmt.Println(l)
        count++
    }

    fmt.Println("HTTPVERB:",httpHeader[0])
    fmt.Println("HTTPURI:",httpHeader[1])

    tmpl, err := template.New("name").Parse(foundmeText)
    if err != nil {
        log.Fatalln("Error parsing template:", err)
    }
    buf := new(bytes.Buffer)
    err = tmpl.Execute(buf, HTTPURL)
    if err != nil {
        log.Fatalln("Error executing template:", err)
    }
    httpRespondOK(c, buf)
}


func hi(c net.Conn) {
	//fmt.Fprintf(w, "Hello! - %s\n", r.URL.Path)
}


func httpRespondOK(c net.Conn, html *bytes.Buffer) {

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", html.Len())
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
    if html.Len() > 0 {
	    fmt.Fprint(c, html)
    }
}


//***************************************************************
//                          m a i n 
//***************************************************************

func main() {

    // Set up to listen on the TCP Port.
    l,err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        log.Fatalln("Error trying to listen: ", err)
    }
    defer l.Close()

    // Listen and Handle connections.
    for {

        // Wait for a connection.
        conn, err := l.Accept()
        if err != nil {
            log.Fatalln("Error on net.Accept: ", err)
        }
        // Handle the new connection.
        go baseHandler(conn)
    }

    log.Fatalln("ERROR - Should never reach here!")

}
