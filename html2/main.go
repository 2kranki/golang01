// vi:nu:et:sts=4 ts=4 sw=4
// How to parse html in Golang using the HTML Parser which
// returns a tree instead of a stream of tokens.
//
// For documentation on this, see:
// https://godoc.org/golang.org/x/net/html#NodeType

package main

import "fmt"
import "io/ioutil"
import "log"
import "strings"
import "golang.org/x/net/html"


func print_node(node *html.Node, indent int) {
    fmt.Print(strings.Repeat(" ",indent))

    switch node.Type {

    // ErrorNode means that an error occurred during tokenization ???
    case html.ErrorNode:
        fmt.Print("Text:",node.Data)
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" Error:", node.Data)

    // TextToken means a text node
    case html.TextNode:
        fmt.Print("Text:",node.Data)
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" Text:",node.Data)

    // ???
    case html.DocumentNode:
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" Document:", node.Data)

    // ???
    case html.ElementNode:
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" Element:", node.Data)

    // A CommentToken looks like <!--x-->
    case html.CommentNode:
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" Comment:",node.Data)

    // A DoctypeToken looks like <!DOCTYPE x>
    case html.DoctypeNode:
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" DocType:",node.Data)

    default:
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" UNKNOWN:",node.Data)
    }

}

func visit_Preorder(node *html.Node, indent int) {
    print_node(node, indent)
    if node.FirstChild != nil {
        visit_Preorder(node.FirstChild, indent+3)
    }
    if node.NextSibling != nil {
        visit_Preorder(node.NextSibling, indent)
    }
}

func main() {

    b, err := ioutil.ReadFile("./data/app01sq_list_html2.txt")
    if err != nil {
        log.Fatal(err)
    }
    str := string(b)
    rdr := strings.NewReader(str)

    tokens, err := html.Parse(rdr)
    if err != nil {
        log.Fatal(err)
    }
    visit_Preorder(tokens, 0)

    fmt.Println("We are done!")
}

