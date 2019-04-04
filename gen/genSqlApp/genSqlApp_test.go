// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// Test Generate SQL Application Generator

package genSqlApp

import (
	"testing"
	"time"
	//"../util"
)

const jsonTestPath = "../test/test01"

func TestReadJsonFile(t *testing.T) {
	var	json		interface{}

	defns = map[string]interface{}{}
	defns[debugId] = true
	debug = true
	defns[forceId] = false
	force = false
	defns[jsonDirId] = jsonTestPath
	defns[mdldirId] = "../models/sqlapp/"
	defns[outdirId] = "../test/"
	defns[quietId] = false
	quiet = false
	defns[timeId] = time.Now().Format("Mon Jan _2, 2006 15:04")

	json,err := ReadJsonFile("app.json.txt", "")
	if err != nil {
		t.Errorf("ReadJsonFile() failed: %s\n", err)
	}
	if json.Database != "Finances" {
		t.Errorf("ReadJsonFile() failed: Name should be 'Finances' but is '%s'\n", json.Database)
	}
	if len(json.Tables) != 2 {
		t.Errorf("ReadJsonFile() failed: should be 2 tables but is %d\n", len(json.Tables))
	}
	if len(json.Tables[0].Fields) != 8 {
		t.Errorf("ReadJsonFile() failed: should be 8 fields in table[0] but is %d\n", len(json.Tables[0].Fields))
	}

	//t.Log(logData.String())
	t.Log("Successfully, completed: TestReadAppJson")
}

func TestCheckForExistingOutputFile(t *testing.T) {
	//var	defns		map[string]interface{}

	defns = map[string]interface{}{}
	defns[debugId] = true
	debug = true
	defns[forceId] = false
	force = false
	defns[jsonDirId] = jsonTestPath
	defns[mdldirId] = "../models/sqlapp/"
	defns[outdirId] = "../test/"
	defns[quietId] = false
	quiet = false
	defns[timeId] = time.Now().Format("Mon Jan _2, 2006 15:04")

	_,err := util.IsPathRegularFile("xyzzy.go")
	if err != nil {
		t.Errorf("CheckForExistingOutputFile(xyzzy.go) failed: %s\n", err)
	}

	t.Log("Successfully, completed: TestCheckForExistingOutputFile")
}

func TestGenTextFile(t *testing.T) {

	defns = map[string]interface{}{}
	defns[debugId] = true
	debug = true
	defns[forceId] = false
	force = false
	defns[jsonDirId] = jsonTestPath
	defns[noopId] = true
	noop = true
	defns[mdldirId] = "../models/sqlapp/"
	defns[outdirId] = "../test/"
	defns[quietId] = false
	quiet = false
	defns[timeId] = time.Now().Format("Mon Jan _2, 2006 15:04")

	err := GenTextFile("sqlAppMain.txt", "main.go")
	if err == nil {
		t.Errorf("GenTextFile(sqlAppMain.txt, main.go) failed: %s\n", err)
	}

	t.Log("Successfully, completed: TestGenFile")
}
