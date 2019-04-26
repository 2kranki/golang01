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
	"../dbPkg"
	"../mainData"
	"../shared"
	"../util"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)


// FileDefn gives the parameters needed to generate a file.  The fields of
// the struct have been simplified to allow for easy json encoding/decoding.
type FileDefn struct {
	ModelName	string 		`json:"ModelName,omitempty"`
	FileDir		string		`json:"FileDir,omitempty"`
	FileName  	string 		`json:"FileName,omitempty"`
	FileType  	string 		`json:"Type,omitempty"`  		// text, sql, html
	Class     	string 		`json:"Class,omitempty"` 		// single, table
	PerTable  	bool		`json:"PerTable,omitempty"` 	// true == generate one file per table
}

// FileDefns controls what files are generated.
var FileDefns []FileDefn = []FileDefn{
	{"base.html.tmpl.txt",
		"/tmpl",
		"base.html.tmpl",
		"copy",
		"one",
		false,
	},
	{"form.html",
		"/tmpl",
		"form.html",
		"copy",
		"one",
		false,
	},
	{"main.go.tmpl.txt",
		"",
		"main.go",
		"text",
		"one",
		false,
	},
	{"mainExec.go.tmpl.txt",
		"",
		"mainExec.go",
		"text",
		"single",
		false,
	},
	{"handlers.go.tmpl.txt",
		"/handlers",
		"handlers.go",
		"text",
		"single",
		false,
	},
	{"table.handlers.go.tmpl.txt",
		"/handlers",
		".go",
		"text",
		"single",
		true,
	},
	{"tableio.go.tmpl.txt",
		"/tableio",
		"tableio.go",
		"text",
		"single",
		false,
	},
	{"table.io.go.tmpl.txt",
		"/tableio",
		".go",
		"text",
		"single",
		true,
	},
}

// TmplData is used to centralize all the inputs
// to the generators.  We maintain generic JSON
// structures for the templating system which does
// not support structs.  (Not certain why yet.)
// We also maintain the data in structs for easier
// access by the generation functions.
type TmplData struct {
	//DataJson 	map[string]interface{}
	Data     	*appData.Database
	//MainJson 	map[string]interface{}
	Main     	*mainData.MainData
	Name		string						// Table Name (if present)
}

var tmplData TmplData

type TaskData struct {
	FD			FileDefn
	TD			*TmplData
	Table		appData.DbTable
	PathIn	  	string						// Input File Path
	PathOut	  	string						// Output File Path

}

func copyFile(modelPath, outPath string) (int64, error) {
	var dst *os.File
	var err error
	var src *os.File

	if _, err = util.IsPathRegularFile(modelPath); err != nil {
		return 0, errors.New(fmt.Sprint("Error - model file does not exist:", modelPath, err))
	}

	if outPath, err = util.IsPathRegularFile(outPath); err == nil {
		if sharedData.Force() {
			if err = os.Remove(outPath); err != nil {
				return 0, errors.New(fmt.Sprint("Error - could not delete:", outPath, err))
			}
		} else {
			return 0, errors.New(fmt.Sprint("Error - overwrite error of:", outPath))
		}
	}
	if dst, err = os.Create(outPath); err != nil {
		return 0, errors.New(fmt.Sprint("Error - could not create:", outPath, err))
	}
	defer dst.Close()

	if src, err = os.Open(modelPath); err != nil {
		return 0, errors.New(fmt.Sprint("Error - could not open model file:", modelPath, err))
	}
	defer src.Close()

	amt, err := io.Copy(dst, src)

	return amt, err
}

func createModelPath(fn string) (string, error) {
	var modelPath   string
	var err         error

	if modelPath, err = util.IsPathRegularFile(fn); err == nil {
		return modelPath, err
	}

	// Calculate the model path.
	modelPath = sharedData.MdlDir()
	modelPath += "/sqlapp"
	modelPath += "/"
	modelPath += fn
	//modelPath = filepath.Clean(modelPath)				// Clean() is part of Abs()
	modelPath, err = filepath.Abs(modelPath)

	return modelPath, err
}

func createOutputPath(dir string, fn string, tn string) (string, error) {
	var outPath string
	var err error

	outPath = sharedData.OutDir()
	outPath += "/"
	if len(dir) > 0 {
		outPath += dir
		outPath += "/"
	}
	if len(tn) > 0 {
		outPath += tn
	}
	outPath += fn
	outPath, err = util.IsPathRegularFile(outPath)
	if err == nil {
		if !sharedData.Force() {
			return outPath, errors.New(fmt.Sprint("Over-write error of:", outPath))
		}
	}

	return outPath, nil
}

func genFile(task TaskData) {
	var err         error

	// Now generate the file.
	switch task.FD.FileType {
	case "copy":
		if sharedData.Noop() {
			if !sharedData.Quiet() {
				log.Printf("\tShould have copied from %s to %s\n", task.PathIn, task.PathOut)
			}
		} else {
			if amt, err := copyFile(task.PathIn, task.PathOut); err == nil {
				if !sharedData.Quiet() {
					log.Printf("\tCopied %d bytes from %s to %s\n", amt, task.PathIn, task.PathOut)
				}
			} else {
				log.Fatalf("Error - Copied %d bytes from %s to %s with error %s\n",
					amt, task.PathIn, task.PathOut, err)
			}
		}
	case "html":
		if err = GenHtmlFile(task.PathIn, task.PathOut, task); err == nil {
			if !sharedData.Quiet() {
				log.Printf("\tGenerated HTML from %s to %s\n", task.PathIn, task.PathOut)
			}
		} else {
			log.Fatalf("Error - Generated HTML from %s to %s with error %s\n",
				task.PathIn, task.PathOut, err)
		}
	case "text":
		if err = GenTextFile(task.PathIn, task.PathOut, task); err == nil {
			if !sharedData.Quiet() {
				log.Printf("\tGenerated HTML from %s to %s\n", task.PathIn, task.PathOut)
			}
		} else {
			log.Fatalf("Error - Generated HTML from %s to %s with error %s\n",
				task.PathIn, task.PathOut, err)
		}
	default:
		log.Fatalln("Error: Invalid file type:", task.FD.FileType, "for", task.FD.ModelName, err)
	}


}

// readJsonFiles reads in the two JSON files that define the
// application to be generated.
func readJsonFiles() error {
	var err error

	if err = mainData.ReadJsonFileMain(sharedData.MainPath()); err != nil {
		return errors.New(fmt.Sprintln("Error: Reading Main Json Input:", sharedData.MainPath(), err))
	}

	if err = appData.ReadJsonFileApp(sharedData.DataPath()); err != nil {
		return errors.New(fmt.Sprintln("Error: Reading Main Json Input:", sharedData.DataPath(), err))
	}

    return nil
}

func GenSqlApp(inDefns map[string]interface{}) error {
	var err 	error
	var pathIn	string
	//var ok 		bool

	if sharedData.Debug() {
		log.Println("\tsql_app: In Debug Mode")
		log.Printf("\t  args: %q\n", flag.Args())
	}

    // Read the JSON files.
    if err = readJsonFiles(); err != nil {
		log.Fatalln(err)
    }

	// Set up template data
	//if tmplData.MainJson, ok = mainData.MainJson().(map[string]interface{}); !ok {
		//log.Fatalln("Error - Could not type assert mainData.MainJson()")
	//}
	tmplData.Main = mainData.MainStruct()
	//if tmplData.DataJson, ok = appData.AppJson().(map[string]interface{}); !ok {
	//	log.Fatalln("Error - Could not type assert appData.AppJson()")
	//}
	tmplData.Data = appData.AppStruct()

	// Set up the output directory structure
    if !sharedData.Noop() {
        tmpName := path.Clean(sharedData.OutDir())
        if err = os.RemoveAll(tmpName); err != nil {
            log.Fatalln("Error: Could not remove output directory:", tmpName, err)
        }
        tmpName = path.Clean(sharedData.OutDir() + "/tmpl")
        if err = os.MkdirAll(tmpName, os.ModeDir+0777); err != nil {
            log.Fatalln("Error: Could not create output directory:", tmpName, err)
        }
        tmpName = path.Clean(sharedData.OutDir() + "/handlers")
        if err = os.MkdirAll(tmpName, os.ModeDir+0777); err != nil {
            log.Fatalln("Error: Could not create output directory:", tmpName, err)
        }
        tmpName = path.Clean(sharedData.OutDir() + "/tableio")
        if err = os.MkdirAll(tmpName, os.ModeDir+0777); err != nil {
            log.Fatalln("Error: Could not create output directory:", tmpName, err)
        }
    }

	// Setup the worker queue.
	done := make(chan bool)
	inputQueue := util.Workers(
					func(a interface{}) {
						genFile(a.(TaskData))
					},
					func() {
						done <- true
					},
					5)

	for _, def := range FileDefns {

		if !sharedData.Quiet() {
			log.Println("Process file:", def.ModelName, "generating:", def.FileName, "...")
		}

		// Create the input model file path.
		if pathIn, err = createModelPath(def.ModelName); err != nil {
			return errors.New(fmt.Sprintln("Error:", pathIn, err))
		}
		if sharedData.Debug() {
			log.Println("\t\tmodelPath=", pathIn)
		}

		// Now generate the file.
		if def.PerTable {
			appData.ForTables(
				func(v appData.DbTable) {
					data := TaskData{FD: def, TD:&tmplData, Table:v, PathIn:pathIn}
					if data.PathOut, err = createOutputPath(def.FileDir, def.FileName, v.Name); err != nil {
						log.Fatalln(err)
					}
					if sharedData.Debug() {
						log.Println("\t\t outPath=", data.PathOut)
					}
					// Generate the file.
					inputQueue <- data
				})
		} else {
			data := TaskData{FD:def, TD:&tmplData, PathIn:pathIn}
			// Create the output path
			if data.PathOut, err = createOutputPath(def.FileDir, def.FileName, ""); err != nil {
				log.Fatalln(err)
			}
			if sharedData.Debug() {
				log.Println("\t\t outPath=", data.PathOut)
			}
			// Generate the file.
			inputQueue <- data
		}
	}
	close(inputQueue)

	<-done

	return nil
}
