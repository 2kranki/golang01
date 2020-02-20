// vi:nu:et:sts=4 ts=4 sw=4
// How to parse html in Golang using the HTML Tokenizer.
//
// Warning: The HTML Tokenizer is a one-pass parser.  It is not a tree
//          structure that you can do passes over. If you want a tree
//          like structure, then you should use html.Parse().
// 


// 1/9/2020 - I modified this to validate the array data. I needed this
//          validation for unit testing in genapp and was the reason that
//          I did this experimentation. The test is a simple check on the
//          first TD's Element a's Text to see that it is a number and
//          progresses from 0;
//          See: https://github.com/2kranki/genapp



package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strconv"
import "strings"
import "golang.org/x/net/html"

func ValidateFile(path string) error {
    var in_td       bool = false
    var num_td      int
    var num_entry   int
    var nameEven    string

    b, err := ioutil.ReadFile(path)
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
            if in_td && len(t.Data) == 1 {
                if (num_td & 1) == 1 {
                    num, err := strconv.Atoi(t.Data)
                    if err != nil {
                        log.Fatalf("%s\n", err.Error())
                    }
                    if num != num_entry {
                        fmt.Println("==>Test failed!\n\n")
                        return fmt.Errorf("Invalid integer: %d %s\n",num_entry,t.Data)
                    }
                    fmt.Println("\t\t\tNumEntry:",num_entry)
                    num_entry++
                } else {
                    if nameEven != "" && (nameEven[0] + 1) != t.Data[0] {
                        fmt.Println("==>Test failed!\n\n")
                        return fmt.Errorf("Invalid name: %s %s\n",nameEven,t.Data)
                    }
                    nameEven = t.Data
                    fmt.Println("\t\t\tNameEven:",nameEven)
                }
            }

        // A StartTagToken looks like <a>
        case html.StartTagToken:
            t := tokens.Token()
            depth++
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" StartTag:", depth, t.Data)
            if t.Data == "td" {
                in_td = true
                num_td++
            }

        // An EndTagToken looks like </a>
        case html.EndTagToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            depth--
            fmt.Println(" EndTag:", depth, t.Data)
            if t.Data == "td" {
                in_td = false
            }

         // A SelfClosingTagToken tag looks like <br/>
        case html.SelfClosingTagToken:
            t := tokens.Token()
            for i:=0; i<len(t.Attr); i++ {
                fmt.Print(" Attr:", t.Attr[i])
            }
            fmt.Println(" Self Closing:",tt)

        // A CommentToken looks like <!--x-->
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

    fmt.Println("==>Test succeeded!\n\n")
    return nil
}

func main() {
    var err error

    fmt.Println("Testing HTML2...\n")
    err = ValidateFile("./data/app01sq_list_html2.txt")
    fmt.Println("...End of HTML2 Tests...\n")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(4)
    }

    fmt.Println("Testing HTML3...\n")
    err = ValidateFile("./data/app01sq_list_html3.txt")
    fmt.Println("...End of HTML3 Tests...\n")
    if err == nil {
        fmt.Print("Error - Test for html3 file should have failed!\n")
        os.Exit(4)
    }

}

