// vi:nu:et:sts=4 ts=4 sw=4
// How to parse html in Golang
//
// Warning: The HTML Tokenizer is a one-pass parser.  It is not a tree
//          structure that you can do passes over. If you want that,
//          you will need to build it yourself.
// 

package main

import "fmt"
import "io/ioutil"
import "strings"
import "golang.org/x/net/html"

func main() {

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
        fmt.Printf("Token: type:%T  ", tt)
        switch tt {
        // ErrorToken means that an error occurred during tokenization
        case html.ErrorToken:
            fmt.Println("End")
            break loop
        // TextToken means a text node
        case html.TextToken:
            t := tokens.Token()
            fmt.Println("Text:",t.Data)
        // A StartTagToken looks like <a>
        case html.StartTagToken:
            t := tokens.Token()
            depth++
            fmt.Println("StartTag:", depth, t.Data)
        // An EndTagToken looks like </a>
        case html.EndTagToken:
            t := tokens.Token()
            depth--
            fmt.Println("EndTag:", depth, t.Data)
         // A SelfClosingTagToken tag looks like <br/>
        case html.SelfClosingTagToken:
            fmt.Println("Self Closing:",tt)
        // A CommentToken looks like <!--x-->
        case html.CommentToken:
            fmt.Println("Comment:",tt)
        // A DoctypeToken looks like <!DOCTYPE x>
        case html.DoctypeToken:
            fmt.Println("DocType:",tt)
        }
    }

    fmt.Println("We are done!")
}

