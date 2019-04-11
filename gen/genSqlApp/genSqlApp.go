// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Generate SQL Application programs in go

// Notes:
//	1.	The html and text templating systems require that
//		their data be separated since it is not identical.
//		So, we put them in separate files.
//	2.	The html and text templating systems access generic
//		structures with range, with, if.  They do not handle
//		structures well especially arrays of structures within
//		structures.

package genSqlApp

import (
	"../mainData"
	"../shared"
	"../util"
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
)

const (
	jsonDirCon = "./"
	// Merged from main.go
	cmdId       = "cmd"
	jsonDirId   = "jsondir"
	nameId      = "name"
	timeId      = "time"
)

// FileDefn gives the parameters needed to generate a file.  The fields of
// the struct have been simplified to allow for easy json encoding/decoding.
type FileDefn struct {
	ModelName 		string 		`json:"ModelName,omitempty"`
	FileName 		string 		`json:"FileName,omitempty"`
	FileType		string 		`json:"Type,omitempty"`			// text, sql, html
	Class  			string    	`json:"Class,omitempty"`		// single, table
	PrefixTable		bool 		`json:"PrefixTable,omitempty"`	// Prefix output file name with table name
}

// FileDefns controls what files are generated.
var FileDefns	[]FileDefn = []FileDefn {
	{"main.go.tmpl.txt",
		"main.go",
		"text",
		"one",
		false,
	},
	{"mainExec.go.tmpl.txt",
		"mainExec.go",
		"text",
		"single",
		false,
	},
	{"tableio.go.tmpl.txt",
		"tableio.go",
		"text",
		"single",
		false,
	},
}

func init() {

}

func createModelPath(fn string) (string, error) {
	var modelPath 	string
	var ok 			bool
	var err			error

	modelPath,err = util.IsPathRegularFile(fn)
	if err == nil {
		return modelPath,err
	}

	// Calculate the model path.
	modelPath = sharedData.MdlDir()
	modelPath += "/sqlapp"
	modelPath += "/"
	modelPath += fn
	//modelPath = filepath.Clean(modelPath)				// Clean() is part of Abs()
	modelPath,err = filepath.Abs(modelPath)

	return modelPath,err
}

func createOutputPath(fn string) (string, error) {
	var outPath 	string
	var ok 			bool
	var err 		error

	outPath = sharedData.MdlDir()
	outPath += "/"
	outPath += fn
	outPath, err = util.IsPathRegularFile(outPath)
	if err == nil {
		if !sharedData.Force() {
			return outPath, errors.New(fmt.Sprint("Over-write error of:", outPath))
		}
	}

	return outPath, nil
}

func GenSqlApp(inDefns map[string]interface{}) error {
	var err			error

	if sharedData.Debug() {
		log.Println("\tsql_app: In Debug Mode")
		log.Printf("\t  args: %q\n", flag.Args())
	}

	// Read in the primary json files.

	// Now handle each FileDefn creating a file for it.
	for _, def := range(FileDefns) {
		var modelPath	string
		var outPath		string

		if !sharedData.Quiet() {
			log.Println("Process file:",def.ModelName,"generating:",def.FileName,"from:",def.JsonPath,"...")
		}

		//if def.PreprocSql {
			//PreprocessSql(json.(DatabaseData))
		//}

		// Create the input model file path.
		if modelPath, err := createModelPath(def.ModelName); err !=  nil {
			return errors.New(fmt.Sprintln("Error:", modelPath, err))
		}
		if sharedData.Debug() {
			log.Println("\t\tmodelPath=", modelPath)
		}

		// Create the output path
		if outPath, err := createOutputPath(def.FileName); err != nil {
			log.Fatalln(err)
		}
		if sharedData.Debug() {
			log.Println("\t\t outPath=", outPath)
		}

		// Now generate the file.
		switch def.FileType {
			case "html":
				err = GenHtmlFile(modelPath, def.FileName, def.PrefixTable, data)
			case "text":
				err = GenTextFile(modelPath, def.FileName, def.PrefixTable, data)
			default:
				return errors.New(fmt.Sprint("Error: Invalid file type:", def.FileType,
					"for",def.JsonPath, err))
		}
		if err != nil {
			return errors.New(fmt.Sprint("Error: Processing file:", def.JsonPath, err))
		}
	}

	return nil
}

