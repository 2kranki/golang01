// See License.txt in main repository directory

// Template Functions used in generation
// See the specific template files for how the functions
// and data are used.

package dbPkg

import (
	"../shared"
	"testing"
)

var tbl0sql = "DROP TABLE Customer IF EXISTS;\nCREATE TABLE Customer (\n\tCustNo\tINT PRIMARY KEY,\n\tCustName\tNVARCHAR(30),\n\tCustCurBal\tDEC(15,2)\n);\n"
var tbl0struct = "type Customer struct {\n\tCustNo\tint64\n\tCustName\tstring\n\tCustCurBal\tfloat64\n}\n"

func TestCreate(t *testing.T) {
	var err			error
	var str			string

	sharedData.SetMainPath("./test/db.json.txt")
	if err = ReadJsonFile(sharedData.MainPath()); err != nil {
		t.Errorf("TestCreateStruct() Reading Main JSON failed: %s'\n", sharedData.MainPath())
	}

	if len(dbStruct.Tables) != 2 {
		t.Errorf("TestCreateStruct() failed: len(Tables) should be 2 but is '%d'\n", len(dbStruct.Name))
	}
	if len(dbStruct.Tables[0].Fields) != 3 {
		t.Errorf("TestCreateStruct() failed: should be 3 Tables but is %d\n", len(dbStruct.Tables[0].Fields))
	}

	str = dbStruct.Tables[0].CreateSql()
	t.Log("Table[0] CreateSql =", str)
	if str != tbl0sql {
		t.Errorf("TestCreateStruct() failed: invalid create sql generated\n")
	}

	str = dbStruct.Tables[0].CreateStruct()
	t.Log("Table[0] Struct =", str)
	if str != tbl0struct {
		t.Errorf("TestCreateStruct() failed: invalid struct generated\n")
	}

	//t.Log(logData.String())
	t.Log("Successfully, completed: TestCreateStruct")

}

func TestReadJsonFileDb(t *testing.T) {
	var err			error

	sharedData.SetMainPath("./test/db.json.txt")
	if err = ReadJsonFile(sharedData.MainPath()); err != nil {
		t.Errorf("ReadJsonFileDb() Reading Main JSON failed: %s'\n", sharedData.MainPath())
	}

	if len(dbStruct.Tables) != 2 {
		t.Errorf("ReadJsonFileDb() failed: len(Tables) should be 2 but is '%d'\n", len(dbStruct.Name))
	}
	if len(dbStruct.Tables[0].Fields) != 3 {
		t.Errorf("ReadJsonFileDb() failed: should be 3 Fields but is %d\n", len(dbStruct.Tables[0].Fields))
	}
	if len(dbStruct.Tables[1].Fields) != 8 {
		t.Errorf("ReadJsonFileDb() failed: should be 8 Fields but is %d\n", len(dbStruct.Tables[1].Fields))
	}

	//t.Log(logData.String())
	t.Log("Successfully, completed: TestReadAppJson")

}
