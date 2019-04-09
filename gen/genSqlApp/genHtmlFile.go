// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Generate SQL Application programs for GO

package genSqlApp

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var htmlTmpls template.Template

func GenHtmlFile(mdl string, fn string, pt bool, data interface{}) error {
	var outPath string
	var err error
	var funcs template.FuncMap
	var tmpl *template.Template

	log.Printf("\tGenHtmlFile mdl:%s fn:%s ...", mdl, fn)

	// The function map is different between the text and html template packages
	funcs = template.FuncMap{"dblClose": dblClose, "dblOpen": dblOpen, "nameUC": nameUC}

	outData := strings.Builder{}
	if debug {
		log.Println("\t\texecuting template...")
		log.Println("\t\tdata:", data)
	}

	name := filepath.Base(mdl)
	tmpl, err = template.New(name).Funcs(funcs).ParseFiles(mdl)
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(&outData, name, data)
	if err != nil {
		return err
	}

	if !noop {
		// Write the file to disk
		err := ioutil.WriteFile(outPath, []byte(outData.String()), 0664)
		if err != nil {
			return errors.New(fmt.Sprint("Error:", outPath, err))
		}
	} else {
		log.Println("<<<<<<<<<<<<<<<<<<<<<<<<", fn, ">>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(outData.String())
		log.Println("<<<<<<<<<<<<<<<<<<<<<<<<", fn, ">>>>>>>>>>>>>>>>>>>>>>>>>>")
	}

	return err
}
