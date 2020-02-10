// vi:nu:et:sts=4 ts=4 sw=4
// How to parse html in Golang using the HTML Tokenizer.
//
// Before compiling/running this program, you need to run:
//
//      go get golang.org/x/net/html
//
// Warning: The HTML Tokenizer is a one-pass parser.  It is not a tree
//          structure that you can do passes over. If you want a tree
//          like structure, then you should use html.Parse().
// 



package main

import "fmt"
import "io/ioutil"
import "strings"
import "golang.org/x/net/html"

func main() {
    var err error

    b, err := ioutil.ReadFile("./data/app01sq_list_html2.txt")
    if err != nil {
        fmt.Print(err)
    }
    str := string(b)
    rdr := strings.NewReader(str)

    tokens := html.NewTokenizer(rdr)

    depth := 0
loop:
    for {
        tt := tokens.Next()
        fmt.Printf("Token: type:%V  ", tt)
        switch tt {

        // ErrorToken means that an error occurred during tokenization
        case html.ErrorToken:
            fmt.Println(" End:", tokens.Err().Error())
            break loop

        // TextToken means a text node
        case html.TextToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" Text:",t.Data)

        // A StartTagToken looks like <a>
        case html.StartTagToken:
            t := tokens.Token()
            depth++
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" StartTag:", depth, t.Data)

        // An EndTagToken looks like </a>
        case html.EndTagToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            depth--
            fmt.Println(" EndTag:", depth, t.Data)

         // A SelfClosingTagToken tag looks like <br/>, <p/>, etc...
        case html.SelfClosingTagToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" Self Closing:",tt)

        // A CommentToken looks like <!-- x -->
        case html.CommentToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" Comment:",t.Data)

        // A DoctypeToken looks like <!DOCTYPE x>
        case html.DoctypeToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" DocType:",t.Data)
        }
    }

    fmt.Println("==>Scan completed!\n\n")
}

