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

var nodeStack   []*html.Node

func popNode() (*html.Node) {
    node := nodeStack[len(nodeStack)-1]
    nodeStack = nodeStack[:len(nodeStack)-1]
    return node
}

func pushNode(node *html.Node) {
    nodeStack = append(nodeStack, node)
}

func check1(node *html.Node) (int) {

    fmt.Println("Check1:")
    printNode(node, 0)                  // Document
    node = node.FirstChild
    printNode(node, 0)                  // Comment
    node = node.NextSibling
    printNode(node, 0)                  // DocType
    node = node.NextSibling
    printNode(node, 0)                  // Element: html
    node = node.FirstChild
    printNode(node, 0)                  // Element: head
    node = node.NextSibling
    printNode(node, 0)                  // Text:
    node = node.NextSibling
    printNode(node, 0)                  // Element: body
    node = node.FirstChild
    printNode(node, 0)                  // Text:
    node = node.NextSibling
    printNode(node, 0)                  // Element: form
    node = node.FirstChild
    printNode(node, 0)                  // Text:
    node = node.NextSibling
    printNode(node, 0)                  // Element: table 
    node = node.FirstChild
    printNode(node, 0)                  // Text:
    node = node.NextSibling
    printNode(node, 0)                  // Comment: The Data List
    node = node.NextSibling
    printNode(node, 0)                  // Text:
    node = node.NextSibling
    printNode(node, 0)                  // Element: tbody
    node = node.FirstChild
    printNode(node, 0)                  // Text:
    for node != nil {
        node = node.NextSibling
        if node == nil {
            break
        }
        printNode(node, 0)                  // Element: tr
        if node.Type == html.ElementNode && node.Data == "tr" {
            node = node.NextSibling     // Skip Text
        } else {
            break
        }
    }
    fmt.Println("End of Check1")

    return 1
}

func printNode(node *html.Node, indent int) {
    fmt.Print(strings.Repeat(" ",indent))

    if node == nil {
        fmt.Println(" NIL")
        return
    }

    switch node.Type {

    // ErrorNode means that an error occurred during tokenization ???
    case html.ErrorNode:
        for i:=0; i<len(node.Attr); i++ {
            fmt.Print(" Attr:", node.Attr[i])
        }
        fmt.Println(" Error:", node.Data)

    // TextToken means a text node
    case html.TextNode:
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
        if node != nil {
            for i:=0; i<len(node.Attr); i++ {
                fmt.Print(" Attr:", node.Attr[i])
            }
            fmt.Println(" UNKNOWN:",node.Data)
        } else {
            fmt.Println(" NIL")
        }
    }

}

func nextPreorder(node *html.Node) (*html.Node) {
    if !node.Visited {
        printNode(node, (len(nodeStack) * 4))
        node.Visited = true
    }
    if node.FirstChild != nil {
        node.Visited = false
        nodeStack = append(nodeStack, node.FirstChild)
        return node.FirstChild
    }
    if node.NextSibling != nil {
        node.Visited = false
        return node.NextSibling
    }
    node = nodeStack[len(nodeStack)-1]
    nodeStack = nodeStack[:len(nodeStack)-1]
    return node
}

func visitPreorder(node *html.Node, indent int) {
    printNode(node, indent)
    if node.FirstChild != nil {
        fmt.Println("FirstChild:")
        visitPreorder(node.FirstChild, indent+3)
    }
    if node.NextSibling != nil {
        fmt.Println("NextSibling:")
        visitPreorder(node.NextSibling, indent)
    }
    fmt.Println("Return")
}

func main() {
    var nodeRoot *html.Node
    var node *html.Node
    var indent int

    b, err := ioutil.ReadFile("./data/app01sq_list_html2.txt")
    if err != nil {
        log.Fatal(err)
    }
    str := string(b)
    rdr := strings.NewReader(str)

    nodeRoot, err = html.Parse(rdr)
    if err != nil {
        log.Fatal(err)
    }

    indent = check1(nodeRoot)
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("Recursive Preorder:")
    indent = 0
    visitPreorder(nodeRoot, indent)
    fmt.Println("End of Recursive Preorder")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("Stacked Preorder:")
    node = nodeRoot
    for node != nil {
        node = nextPreorder(node)
    }
    fmt.Println("End of Stacked Preorder")

    fmt.Println("We are done!")
}

