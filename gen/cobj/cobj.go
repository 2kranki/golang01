// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in repository directory

// This module is responsible for generating objects in the
// C programming language.

package cobj

import (
	"errors"
	"fmt"
	"log"

	//"os"
	"strings"
	"text/template"
	//"tmpl"
)

const (
	mdlpathCon = "./models/cobj/"
	// Merged from main.go
	cmdId      = "cmd"
	debugId    = "debug"
	forceId    = "force"
	jsonPathId = "jsonpath"
	mdldirId   = "mdlpath"
	nameId     = "name"
	noopId     = "noop"
	outdirId   = "outdir"
	quietId    = "quiet"
	timeId     = "time"
)

// FileDefn gives the parameters needed to generate a file.  The fields of
// the struct have been simplified to allow for easy json encoding/decoding.
type FileDefn struct {
	ModelName 	string 		`json:"ModelName,omitempty"`
	FileName 	string 		`json:"FileName,omitempty"`
	FileType	string 		`json:"Type,omitempty"`			// text, sql, html
	JsonPath  	string    	`json:"JsonPath,omitempty"`		// Json definition path
	Class  		string    	`json:"Class,omitempty"`		// single, table
	OutputType	string		`json:"OutputType,omitempty"`	// database, main or nil
	PrefixTable	bool 		`json:"PrefixTable,omitempty"`	// Prefix output file name with table name
}

type object_generator struct {
	defns map[string]interface{}
	funcs template.FuncMap
	debug bool
	force bool
	quiet bool
}

var genobj object_generator = object_generator{}

func NameUC() string {
	return strings.ToUpper(genobj.defns[nameId].(string))
}

func (c *object_generator) Gen(mdl string, out string) {

}

var CObjGen object_generator

func genFile(mdl string) {
	var mf string

	//mf = modelPrefix + "obj_h.txt"
	//out,err := os.file.Create(genobj.defns[outdirId])
	//if err != nil {
	//	return errors.New(fmt.Sprintf("Missing '%s' in definitions", nameID))
	//}
	//in = tmpl.ExecTmplFile(modelPrefix+"obj_h.txt")
	log.Println(mf)
}

func GenCObj(defns map[string]interface{}) error {
	var ok bool
	var err error

	genobj.defns = defns
	genobj.funcs = template.FuncMap{"NameUC": NameUC}
	genobj.debug = defns[debugId].(bool)
	genobj.force = defns[forceId].(bool)
	genobj.quiet = defns[quietId].(bool)

	if _, ok = defns[nameId]; !ok {
		return errors.New(fmt.Sprintf("Missing 'name' in definitions"))
	}

	return err
}
