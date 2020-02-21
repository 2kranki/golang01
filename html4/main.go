// vi:nu:et:sts=4 ts=4 sw=4
// How to parse html in Golang using the HTML Parser which
// returns a tree instead of a stream of tokens which can
// be scanned multiple times and several different ways.
//
// In this example, we hard code the check for TBody/tr/
// Element_a/Text <number> where we are looking for 
// consecutive numbers starting with 0.
//
// Initially, it was very time consuming constructing the
// scan to get to the tr's, but this method does work and
// allows for some simple checking of the generated html.
//
// The second implementation of the checking was done using
// the recursive search and was much faster to implement
// since we had the recursive search from html3 already.
//
// The book, "The Go Programming Language" starting on 
// page 121, has yet another way to scan the tree.
// However, it combines the FirstChild with NextSibling
// which I did not want for this process.  It is a
// wonderful book if you have not read it and covers
// many aspects of Go.
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

var chk int

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

//******************************************************************
//                  S t a t i c  C h e c k 1
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
        node = node.NextSibling         // Scan the sibling chain of tbody.
        if node == nil {
            break
        }
        if node.Type == html.TextNode {
            continue
        }
        printNode(node, 3)                  // Element: tr
        if node.Type == html.ElementNode && node.Data == "tr" {
            node1 = node.FirstChild         // Text
            node1 = node1.NextSibling       // td
            // Now we should have an Element with a following Text.
            printNode(node1, 6)             // Element: td
            if node1.Type == html.ElementNode && node1.Data == "td" {
                node1 = node1.FirstChild    // Element: a
                printNode(node1, 9)
                node1 = node1.FirstChild       // Text
                printNode(node1, 12)
                if node1.Type == html.TextNode {
                    num, err := strconv.Atoi(node1.Data)
                    if err != nil {
                        log.Fatalf("%s\n", err.Error())
                    }
                    if num != chk {
                        fmt.Errorf("Error: check failed for Text %d looking for %d\n", num, chk)
                        return 0
                    }
                    chk++
                    continue
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

var stopRecurse bool
var nodeTR *html.Node

func visit_Preorder_TR(node *html.Node, indent int) {

    if stopRecurse {
        return
    }
    if node.Type == html.ElementNode && node.Data == "tr" {
        nodeTR = node
        stopRecurse = true
    }
    printNode(node, indent)
    if node.FirstChild != nil {
        visit_Preorder_TR(node.FirstChild, indent+3)
    }
    if node.NextSibling != nil {
        visit_Preorder_TR(node.NextSibling, indent)
    }
}

//******************************************************************
//                  S t a t i c  C h e c k 2
// For this check, we use the recursive search to find TR node which
// was much easier to implement.  
//******************************************************************
func check2(node *html.Node) (int) {

    visit_Preorder_TR(node, 0)
    if nodeTR == nil {
        log.Fatalf("ERROR: check2 never found the TR node!")
    }

    // Scan TD's
    node = node.FirstChild
    for node != nil {
        var node1 *html.Node

        if node.Type == html.TextNode {
            node = node.NextSibling
            continue
        }
        printNode(node, 3)                  // Element: td
        if node.Type == html.ElementNode && node.Data == "td" {
            if node.FirstChild.Type == html.ElementNode && node.FirstChild.Data == "a" {
                node1 = node.FirstChild
                printNode(node1, 9)
                node1 = node1.FirstChild       // Text
                printNode(node1, 12)
                if node1.Type == html.TextNode {
                    num, err := strconv.Atoi(node1.Data)
                    if err != nil {
                        log.Fatalf("%s\n", err.Error())
                    }
                    if num != chk {
                        fmt.Errorf("Error: check failed for Text %d looking for %d\n", num, chk)
                        return 0
                    }
                    chk++
                    continue
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

        node = node.NextSibling         // Scan the sibling chain of tbody.
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

    fmt.Println("Static Check1 of HTML2...")
    chk = 0
    indent = check1(nodeRoot)
    if indent > 0 {
        fmt.Println("...Static Check Succeeded!")
    }
    fmt.Println("End of Static Check1 for HTML2!")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")

    fmt.Println("Static Check2 of HTML2...")
    chk = 0
    indent = check1(nodeRoot)
    if indent > 0 {
        fmt.Println("...Static Check Succeeded!")
    }
    fmt.Println("End of Static Check2 for HTML2!")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")

    b, err = ioutil.ReadFile("./data/app01sq_list_html3.txt")
    if err != nil {
        log.Fatal(err)
    }
    str = string(b)
    rdr = strings.NewReader(str)

    nodeRoot, err = html.Parse(rdr)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Static Check1 of HTML3...")
    chk = 0
    indent = check1(nodeRoot)
    if indent == 0 {
        fmt.Println("...Static Check Failed, but it should have failed!")
    } else {
        log.Fatalf("ERROR: Static Check succeeded and it should not have!\n")
    }
    fmt.Println("End of Static Check1 for HTML3!")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")

    fmt.Println("Static Check2 of HTML3...")
    chk = 0
    indent = check1(nodeRoot)
    if indent == 0 {
        fmt.Println("...Static Check Failed, but it should have failed!\n")
    } else {
        log.Fatalf("ERROR: Static Check succeeded and it should not have!\n")
    }
    fmt.Println("End of Static Check2 for HTML3!")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("")

     fmt.Println("We are done!")
}

