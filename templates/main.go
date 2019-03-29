// See License.txt in main respository directory

// Program to play with templates and see what is
// available for us to use in this context.

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type xyz struct {
	T1		string
	T2		[]string
	M1		map[string]string
}

var (
	debug  		= flag.Bool("debug", false, "enable debugging")
	quiet  		= flag.Bool("quiet", false, "Quiet all messages")

	//force  	= flag.Bool("force", false, "over-write output")
	file   		= flag.String("file", "tmpl01.tmpl.txt", "input template file (default: tmpl01.tmpl.txt)")
)
var startTime	time.Time
var Tmpls       *template.Template
var	Vars		xyz
var	Funcs		template.FuncMap
var	pgmNameStr	string = "abc"


func dblClose() string {
	return "}}"
}

func dblOpen() string {
	return "{{"
}

func displayPath() string {
	return *file
}

func pgmName() string {
	return pgmNameStr
}

func pgmNameUC() string {
	return strings.ToUpper(pgmNameStr)
}

func typeOf(x interface{}) string {
	return fmt.Sprintf("%T", x)
}


// IsPathRegularFile cleans up the supplied file path
// and then checks the cleaned file path to see
// if it is an existing standard file. Return the
// cleaned up path and a potential error if it exists.
func IsPathRegularFile(fp string) (string,error) {
	var	err 	error
	var path	string

	fp = filepath.Clean(fp)
	path,err = filepath.Abs(fp)
	if err != nil {
		return path,errors.New(fmt.Sprint("Error getting absolute path for:", fp, err))
	}
	fi, err := os.Lstat(path)
	if err != nil {
		return path,errors.New("path not found")
	}
	if fi.Mode().IsRegular() {
		return path,nil
	}
	return path,errors.New("path not regular file")
}



// ExecTmplFile executes any given template while caching it for
// future use.
// Input:
//		w		io.Writer
//		path	string
//		data	interface{}				// nil allowed
//		funcs	template.FuncMap		// nil allowed
func ExecTmplFile(w io.Writer, path string, data interface{}, funcs template.FuncMap) {
	var err error

	if w == nil {
		w = os.Stdout
	}
	tmpl := Tmpls.Lookup(path)
	if tmpl == nil {
		path,err = IsPathRegularFile(path)
		if err != nil {
			log.Fatalln("Error while stating file:", path, err)
		}
		if Tmpls == nil {
			Tmpls, err = template.New(path).Funcs(funcs).ParseFiles(path)
		} else {
			Tmpls, err = Tmpls.New(path).Funcs(funcs).ParseFiles(path)
		}
		if err != nil {
			log.Fatalln("Error parsing templates for:", path, err)
		}
	}
	err = tmpl.ExecuteTemplate(w, path, data)
	if err != nil {
		log.Fatalln("Error executing template:", path, err)
	}
}



func ExecTmplStr(w io.Writer, name string, tmpl string, data interface{}, funcs template.FuncMap) {
	var err 	error
	var tmplc	*template.Template

	if w == nil {
		w = os.Stdout
	}
	if Tmpls != nil {
		tmplc = Tmpls.Lookup(name)
	}
	if tmplc == nil {
		if Tmpls == nil {
			Tmpls,err = template.New(name).Funcs(funcs).Parse(tmpl)
		} else {
			Tmpls,err = Tmpls.New(name).Funcs(funcs).Parse(tmpl)
		}
		if err != nil {
			log.Fatalln("Error parsing templates for:", name, err)
		}
	}
	err = Tmpls.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatalln("Error executing template:", name, err)
	}
}



func TestTmplStr(name string, tmpl string, data interface{}, funcs template.FuncMap, result string) {
	var sb  	*bytes.Buffer

	sb = new(bytes.Buffer)
	ExecTmplStr(sb, name, tmpl, data, funcs)
	if result != sb.String() {
		log.Fatalln("Error:", name, " test did not Work!",
				"\nResult was:'", sb.String(), "' and should have been:'", result, "'")
	}
	fmt.Println("\tSuccessfully ran:",name)
}


func main() {
	//var err 	error
	//var sb  	bytes.Buffer
	//var test   	string

	startTime = time.Now()
	flag.Parse()
	if !*quiet {
		fmt.Println("\tStart:", startTime)
	}
	if *debug {
		fmt.Println("\tIn Debug Mode")
	}
	Funcs = template.FuncMap{
		"dblClose":dblClose,
		"dblOpen":dblOpen,
		"displayPath":displayPath,
		"pgmNameUC":pgmNameUC,
		"typeOf":typeOf,
	}

	// Clean up the input file path and insure that it exists
	fmt.Println("file:", *file)
	fmt.Printf("file Type: %T\n", file)

	// Let's try some simple strings.
	TestTmplStr("test01", "{{typeOf 0}}", nil, Funcs,"int")
	TestTmplStr("test02", "{{displayPath}}", nil, Funcs,"tmpl01.tmpl.txt")
	TestTmplStr("test03", "{{pgmNameUC}}", nil, Funcs,"ABC")
	TestTmplStr("test04", "{{dblOpen}}-{{dblClose}}", nil, Funcs,"{{-}}")
}
