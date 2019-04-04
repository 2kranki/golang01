// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Generate SQL Application programs for GO

package genSqlApp

import (
	"../util"
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

const (
	jsonDirCon = "./"
	mdldirCon   = "./models/sqlapp/"
	outdirCon   = "./test/"
	// Merged from main.go
	cmdId       = "cmd"
	debugId     = "debug"
	forceId     = "force"
	jsonDirId   = "jsondir"
	mdldirId    = "mdldir"
	nameId      = "name"
	noopId      = "noop"
	outdirId    = "outdir"
	quietId     = "quiet"
	timeId      = "time"
)

// FileDefn gives the parameters needed to generate a file.  The fields of
// the struct have been simplified to allow for easy json encoding/decoding.
type FileDefn struct {
	ModelName 	string 		`json:"ModelName,omitempty"`
	FileName 	string 		`json:"FileName,omitempty"`
	FileType	string 		`json:"Type,omitempty"`			// text, sql, html
	JsonPath  	string    	`json:"JsonPath,omitempty"`
	Class  		string    	`json:"Class,omitempty"`		// single, table
	OutputType	string		`json:"OutputType,omitempty"`	// database, main or nil
	PrefixTable	bool 		`json:"PrefixTable,omitempty"`	// Prefix output file name with table name
}

type DbField struct {
	Name 		string 		`json:"Name,omitempty"`
	Type 		string 		`json:"Type,omitempty"`
	Len  		int    		`json:"Len,omitempty"`
	PrimaryKey  bool    	`json:"PrimaryKey,omitempty"`
	List  		bool    	`json:"List,omitempty"`			// Include in List Report
}

type DbTable struct {
	Name   string `json:"Name,omitempty"`
	Fields []DbField
}

type DatabaseData struct {
	Database string 		`json:"Database,omitempty"`
	SqlType  string 		`json:"SqlType,omitempty"`
	Tables   []DbTable
}

type MainFlag struct {
	Name 		string 			`json:"Name,omitempty"`
	Internal 	string 			`json:"Internal,omitempty"`	// Internal Data Name
	Desc 		string 			`json:"Desc,omitempty"`
	FlagType	string 			`json:"Type,omitempty"`
	Init 		string 			`json:"Init,omitempty"`
}

type MainUsage struct {
	Line		string
	Notes 		[]string
}

type MainData struct {
	Flags		[]MainFlag
	Usage 		MainUsage
}

type TmplData struct {
	Data 		*interface{}
	Defns    	*map[string] interface{}
}

var defns 		map[string] interface{}
//var inputData 	inputData
//var mainData	MainData
var debug 		bool
var force 		bool
var noop 		bool
var quiet 		bool
var FileDefns	[]FileDefn = []FileDefn {
	{"main.go.tmpl.txt","main.go","text","main.json.txt", "one", "MainData", false},
	{"mainExec.go.tmpl.txt","mainExec.go","text","app.json.txt", "single", "DatabaseData", false},
}

func init() {

}

func dblClose() string {
	return "}}"
}

func dblOpen() string {
	return "{{"
}

func flagPrs(mp map[string]interface{}) string {
	var str			strings.Builder
	var desc		string
	var flagType	string
	var init		string
	var internal	string
	var name		string

	//fmt.Println("\t\tflagPrs:", mp)

	// unmap the values
	if mp["Desc"] != nil {
		desc = mp["Desc"].(string)
	}
	if mp["Type"] != nil {
		flagType = mp["Type"].(string)
	}
	if mp["Init"] != nil {
		init = mp["Init"].(string)
	}
	if mp["Internal"] != nil {
		internal = mp["Internal"].(string)
	}
	if mp["Name"] != nil {
		name = mp["Name"].(string)
	}

	// Now create the string from the flag fields
	switch flagType {
	case "bool":
		str.WriteString("flag.BoolVar(&")
	case "int":
		str.WriteString("flag.IntVar(&")
	case "string":
		str.WriteString("flag.StringVar(&")
	case "var":
		str.WriteString("flag.Var(&")
	default:
		str.WriteString("flag type ERROR:")
		if flagType == "" {
			str.WriteString("flagType: EMPTY ")
		} else {
			str.WriteString("flagType:")
			str.WriteString(flagType)
		}
	}
	if len(internal) > 0 {
		str.WriteString(internal)
	} else {
		str.WriteString(name)
	}
	str.WriteString(",")
	if len(init) > 0 {
		if flagType == "string" {
			str.WriteString(fmt.Sprintf("\"%s\"",init))
		} else {
			str.WriteString(init)
		}
	}
	if len(desc) > 0 {
		str.WriteString(fmt.Sprintf(",\"%s\"",desc))
	}
	str.WriteString(")")

	return str.String()
}

func nameUC() string {
	return strings.ToUpper(defns[nameId].(string))
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
	modelPath, ok = defns[mdldirId].(string)
	if ok {
		modelPath += "/sqlapp"
	} else {
		modelPath = mdldirCon
	}
	modelPath += "/"
	modelPath += fn
	//modelPath = filepath.Clean(modelPath)				// Clean() is part of Abs()
	modelPath,err = filepath.Abs(modelPath)

	return modelPath,err
}

func createOutputPath(fn string) (string, error) {
	var outPath string
	var ok bool
	var err error

	outPath, ok = defns[outdirId].(string)
	if !ok {
		outPath = "./test"
	}
	outPath += "/"
	outPath += fn
	outPath = filepath.Clean(outPath)

	outPath, err = util.IsPathRegularFile(outPath)
	if err == nil {
		if !force {
			return outPath, errors.New(fmt.Sprint("Over-write error of:", outPath))
		}
	}

	return outPath, nil
}

func ReadJsonFile(fn string, ot string) (interface{}, error) {
	var err 		error
	var ok 			bool
	var jsonPath 	string
	var data  		interface{}
	var appData		DatabaseData
	var mainData  	MainData

	jsonPath,ok = defns[jsonDirId].(string)
	if !ok {
		jsonPath = jsonDirCon
	}
	jsonPath += "/"
	jsonPath += fn
	jsonPath,_ = filepath.Abs(jsonPath)
	if debug {
		log.Println("json path:", jsonPath)
	}

	// Open the input template file
	switch ot {
	case "database":
		err = util.ReadJsonFileToData(jsonPath, &appData)
		if err != nil {
			log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
		}
		data = appData
	case "main":
		err = util.ReadJsonFileToData(jsonPath, &mainData)
		if err != nil {
			log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
		}
		data = mainData
	default:
		data,err = util.ReadJsonFile(jsonPath)
		if err != nil {
			log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
		}
	}

	if debug {
		log.Println("\tJson Data:", data)
	}

	return data,nil
}

func GenSqlApp(inDefns map[string]interface{}) error {

	defns = inDefns
	debug = inDefns[debugId].(bool)
	force = inDefns[forceId].(bool)
	noop  = inDefns[noopId].(bool)
	quiet = inDefns[quietId].(bool)

	/*	if genobj.Name,ok = inDefns[nameId]; !ok {
			log.Fatalln("Error: missing application name!")
		}
	*/

	if debug {
		log.Println("\tsql_app: In Debug Mode")
		log.Printf("\t  Defns: %q\n", inDefns)
		log.Printf("\t  args: %q\n", flag.Args())
	}

	// Now handle each FileDefn creating a file for it.
	for _,def := range(FileDefns) {

		// Set up the Template Data
		json,err :=	ReadJsonFile(def.JsonPath, def.OutputType)
		if err != nil {
			return errors.New(fmt.Sprint("Error: Reading Json Input:", def.JsonPath, err))
		}
		data := TmplData{&json, &defns}

		// Create the input model file path.
		modelPath,err := createModelPath(def.ModelName)
		if err !=  nil {
			return errors.New(fmt.Sprint("Error:", modelPath, err))
		}
		if debug {
			log.Println("\t\tmodelPath=", modelPath)
		}

		// Create the output path
		outPath, err := createOutputPath(def.FileName)
		if err != nil {
			log.Fatalln(err)
		}
		if debug {
			log.Println("\t\t outPath=", outPath)
		}

		// Now generate the file.
		switch def.FileType {
		case "html":
			err = GenHtmlFile(modelPath, def.FileName, def.PrefixTable, data)
		case "text":
			err = GenTextFile(modelPath, def.FileName, def.PrefixTable, data)
		}
		if err != nil {
			return errors.New(fmt.Sprint("Error: Processing file:", def.JsonPath, err))
		}
	}

	return nil
}
