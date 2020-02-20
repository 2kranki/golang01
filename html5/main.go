// vi:nu:et:sts=4 ts=4 sw=4
// How to parse html in Golang using the HTML Parser which
// returns a tree instead of a stream of tokens which can
// be scanned multiple times and several different ways.
//
// In this example, we use a state machine to check for 
// TBody/tr/Element_a/Text <number> where we are looking 
// for consecutive numbers starting with 0.
//
// For documentation on this, see:
// https://godoc.org/golang.org/x/net/html#NodeType

package main

import "fmt"
import "io/ioutil"
import "log"
import "strconv"
import "strings"
import "golang.org/x/net/html"

var nodeStack   []*html.Node
var chk string
var state int

// Element: TBODY
//      Text
//      Element: TR
//          Text
//          Element: TD
//              Element: a
//                  Text    (number)
//          Text
//          Element: TD
//              Text        (letter)
//          Text
//
const (
    scanning state = iota           // Not in TBody
    st_any
    st_el_a         // Element a in First TD
    st_el_td_0      // First Element TD in TR
    st_el_td_1      // Second Element TD in TR
    st_el_tr        // Element TR (Possibly Multiple TRs)
    st_text_0       // TEXT in TBody
    st_text_1       // TEXT in TR
    st_text_2       // TEXT in First TD/Element a
    st_text_2       // TEXT in Second TD
)
const (
    ev_any event = iota
    ev_FirstChild
    ev_NextSibling
    ev_Return
)


func isEmpty() (bool) {
    if len(nodeStack) > 0 {
        return true
    }
    return false
}

func popNode() (*html.Node) {
    node := nodeStack[len(nodeStack)-1]
    nodeStack = nodeStack[:len(nodeStack)-1]
    return node
}

func pushNode(node *html.Node) {
    nodeStack = append(nodeStack, node)
}

//******************************************************************
//                  S t a t i c  C h e c k
// As you can see, this became much more painful that I thought that
// it would be to find the starting point.  However, it is one way
// of doing some checking on the html.
//******************************************************************
func check1(node *html.Node) (int) {

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
        var node1 *html.Node
        node = node.NextSibling
        if node == nil {
            break
        }
        if node.Type == html.TextNode {
            continue
        }
        printNode(node, 0)                  // Element: tr
        if node.Type == html.ElementNode && node.Data == "tr" {
            node1 = node.FirstChild         // Text
            node1 = node1.NextSibling       // td
            // Now we should have an Element with a following Text.
            printNode(node1, 0)             // Element: td
            if node1.Type == html.ElementNode && node1.Data == "td" {
                node1 = node1.FirstChild    // Element: a
                printNode(node1, 0)
                node1 = node1.FirstChild       // Text
                printNode(node1, 0)
                if node1.Type == html.TextNode && node1.Data == chk {
                    num, err := strconv.Atoi(chk)
                    if err != nil {
                        fmt.Errorf("%s\n", err.Error())
                    }
                    num++
                    chk = strconv.Itoa(num)
                } else {
                    fmt.Print("UNEXPECTED NODE3: ")
                    printNode(node1, 0)                  // Element: td
                    break
                }
            } else {
                fmt.Print("UNEXPECTED NODE2: ")
                printNode(node1, 0)                  // Element: td
                break
            }
        } else {
            fmt.Print("UNEXPECTED NODE1: ")
            printNode(node1, 0)                  // Element: td
            break
        }
    }

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

func visitPreorder(node *html.Node, indent int) {

    printNode(node, indent)

    switch state {
        case scanning:
            if node.Type == html.ElementNode && node.Data == "tr" {
                state = in_el_tr
            }
            break
        case in_el_a:
            if node.Type == html.TextNode {
                if node.Data == chk {
                    num, err := strconv.Atoi(chk)
                    if err != nil {
                        fmt.Errorf("%s\n", err.Error())
                    }
                    num++
                    chk = strconv.Itoa(num)
                }
            }
            state = in_el_tr
            break
        case in_el_td:
            if node.Type == html.ElementNode && node.Data == "a" {
                state = in_el_a
            }
            break
        case in_el_tr:
            if node.Type == html.ElementNode && node.Data == "td" {
                state = in_el_td
            }
            break
    }
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

    fmt.Println("Static Check:")
    indent = check1(nodeRoot)
    chk = "0"
    if indent > 0 {
        fmt.Println("...Static Check Succeeded!")
    }
    fmt.Println("End of Static Check")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")

    fmt.Println("Recursive Check:")
    indent = 0
    chk = "0"
    visitPreorder(nodeRoot, indent)
    fmt.Println("End of Recursive Check")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")

    fmt.Println("We are done!")
}

