// See License.txt in main repository directory

// Template Functions used in generation

package genSqlApp

import (
	"log"
	"path/filepath"
	"../util"
)

func ReadJsonFileData(fn string) error {
	var err 		error
	var ok 			bool
	var jsonPath 	string
	var data  		interface{}

	if jsonPath, ok = defns[jsonDirId].(string); !ok {
		jsonPath = jsonDirCon
	}
	jsonPath += "/"
	jsonPath += fn
	jsonPath,_ = filepath.Abs(jsonPath)
	if debug {
		log.Println("json path:", jsonPath)
	}

	// Read in the json file
	if data, err = util.ReadJsonFile(jsonPath); err != nil {
		log.Fatalln("Error: unmarshalling", jsonPath, ", JSON input file:", err)
	}
	tmplData.DataJson = &data
	tmplData.Data = decodeDatabase(data)

	if debug {
		log.Println("\tJson Data:", data)
	}

	return nil
}


