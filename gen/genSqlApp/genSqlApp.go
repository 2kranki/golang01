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
	"reflect"
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
	ModelName 		string 		`json:"ModelName,omitempty"`
	FileName 		string 		`json:"FileName,omitempty"`
	FileType		string 		`json:"Type,omitempty"`			// text, sql, html
	JsonDataPath	string    	`json:"JsonDataPath,omitempty"`
	JsonMainPath	string		`json:"JsonMainPath,omitempty"`
	Class  			string    	`json:"Class,omitempty"`		// single, table
	PrefixTable		bool 		`json:"PrefixTable,omitempty"`	// Prefix output file name with table name
}

// TmplData is used to centralize all the inputs
// to the generators.  We maintain generic JSON
// structures for the templating system which does
// not support structs.  (Not certain why yet.)
// We also maintain the data in structs for easier
// access by the generation functions.
type TmplData struct {
	DataJson	*map[string] interface{}
	Data		*Database
	MainJson	*map[string] interface{}
	Main		*MainData
	Defns    	*map[string] interface{}
}

// defns is the accumulation of default flags,
// flags entered via the JSON exec file and
// flags entered on the command line.
var defns 		map[string] interface{}
// We pull the following variables from defns
// for ease of access by the other routines
var debug 		bool
var force 		bool
var noop 		bool
var quiet 		bool

// FileDefns controls what files are generated.
var FileDefns	[]FileDefn = []FileDefn {
	{"main.go.tmpl.txt",
		"main.go",
		"text",
		"main.json.txt",
		"data.json.txt",
		"one",
		false,
	},
	{"mainExec.go.tmpl.txt",
		"mainExec.go",
		"text",
		"app.json.txt",
		"data.json.txt",
		"single",
		false,
	},
	{"tableio.go.tmpl.txt",
		"mainExec.go",
		"text",
		"app.json.txt",
		"data.json.txt",
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

	// Read in the json file
	switch ot {
	case "databasexx":
		err = util.ReadJsonFileToData(jsonPath, &appData)
		if err != nil {
			log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
		}
		data = appData
	case "mainxx":
		err = util.ReadJsonFileToData(jsonPath, &mainData)
		if err != nil {
			log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
		}
		data = mainData
	default:
		jd, err := util.ReadJsonFile(jsonPath)
		if err != nil {
			log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
		}
		data = &jd
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

	if debug {
		log.Println("\tsql_app: In Debug Mode")
		log.Printf("\t  Defns: %q\n", inDefns)
		log.Printf("\t  args: %q\n", flag.Args())
	}

	// Now handle each FileDefn creating a file for it.
	for _, def := range(FileDefns) {

		if !quiet {
			log.Println("Process file:",def.ModelName,"generating:",def.FileName,"from:",def.JsonPath,"...")
		}
		// Set up the Template Data
		if json, err :=	ReadJsonFile(def.JsonPath, def.JsonType); err != nil {
			return errors.New(fmt.Sprint("Error: Reading Json Input:", def.JsonPath, err))
		}
		if debug {
			log.Println("JSON TypeOf:",reflect.TypeOf(json))
			v := reflect.ValueOf(json)
			log.Println("JSON Value Type:", v.Kind())
		}
		data := TmplData{Data:&json, Defns:&defns}
		if debug {
			log.Println("data.Data TypeOf:",reflect.TypeOf(data.Data))
			log.Println("*data.Data TypeOf:",reflect.TypeOf(*data.Data))
		}

		//if def.PreprocSql {
			//PreprocessSql(json.(DatabaseData))
		//}

		// Create the input model file path.
		if modelPath, err := createModelPath(def.ModelName); err !=  nil {
			return errors.New(fmt.Sprint("Error:", modelPath, err))
		}
		if debug {
			log.Println("\t\tmodelPath=", modelPath)
		}

		// Create the output path
		if outPath, err := createOutputPath(def.FileName); err != nil {
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

