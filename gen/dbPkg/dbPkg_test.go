// See License.txt in main repository directory

// Template Functions used in generation
// See the specific template files for how the functions
// and data are used.

package dbPkg

import (
	"../shared"
	"testing"
)

func TestReadJsonFileDb(t *testing.T) {
	var err			error

	sharedData.SetMainPath("./test/db.json.txt")
	if err = ReadJsonFileDb(sharedData.MainPath()); err != nil {
		t.Errorf("ReadJsonFileDb() Reading Main JSON failed: %s'\n", sharedData.MainPath())
	}

	if len(dbStruct.Tables) != 2 {
		t.Errorf("ReadJsonFileDb() failed: len(Tables) should be 2 but is '%d'\n", len(dbStruct.Name))
	}
	if len(dbStruct.Tables[0].Fields) != 8 {
		t.Errorf("ReadJsonFileDb() failed: should be 8 Tables but is %d\n", len(dbStruct.Tables[0].Fields))
	}

	//t.Log(logData.String())
	t.Log("Successfully, completed: TestReadAppJson")

}
