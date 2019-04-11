// See License.txt in main repository directory

// mainData is responsible for reading and processing the
// Main JSON file.  It reads it in and supplies the necessary
// functions used in the templating.

package mainData

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"../shared"
	"../util"
)

type MainFlag struct {
	Name 		string 		`json:"Name,omitempty"`
	Internal 	string 		`json:"Internal,omitempty"`	// Internal Data Name
	Desc 		string 		`json:"Desc,omitempty"`
	FlagType	string 		`json:"Type,omitempty"`
	Init 		string 		`json:"Init,omitempty"`
}

type MainUsage struct {
	Line		string
	Notes 		[]string
}

type MainData struct {
	Flags		[]MainFlag
	Usage 		MainUsage
}

var	mainStruct	MainData
var	mainJson	interface{}

func decodeMainFlag(i interface{}) *MainFlag {
	var f		MainFlag
	var m		map[string]interface{}
	var ok		bool

	if m, ok = i.(map[string]interface{}); !ok {
		return nil
	}
	if f.Name, ok = m["Name"].(string); !ok {
		return nil
	}
	if f.Internal, ok = m["Internal"].(string); !ok {
		return nil
	}
	if f.Desc, ok = m["Desc"].(string); !ok {
		return nil
	}
	if f.FlagType, ok = m["FlagType"].(string); !ok {
		return nil
	}
	if f.Init, ok = m["Init"].(string); !ok {
		return nil
	}
	return &f
}

func decodeMainUsage(i interface{}) *MainUsage {
	var t		MainUsage
	var m		map[string]interface{}
	var notes	[]string
	var ok		bool

	if m, ok = i.(map[string]interface{}); !ok {
		return nil
	}
	if t.Line, ok = m["Line"].(string); !ok {
		return nil
	}
	if notes, ok = m["Notes"].([]string); !ok {
		return nil
	}
	for _, v := range notes {
		t.Notes = append(t.Notes, v)
	}
	return &t
}

func decodeMainData(i interface{}) *MainData {
	var t		MainData
	var m		map[string]interface{}
	var ok		bool

	if m, ok = i.(map[string]interface{}); !ok {
		return nil
	}
	if t.Usage, ok = m["Usage"].(MainUsage); !ok {
		return nil
	}
	if t.Flags, ok = m["Flags"].([]MainFlag); !ok {
		return nil
	}
	return &t
}

func flagPrs(mp map[string]interface{}) string {
	var str			strings.Builder
	var desc		string
	var flagType	string
	var init		string
	var internal	string
	var name		string

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

	// Now create the string from the names
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

func MainJson() *interface{} {
	return &mainJson
}

func MainStruct() *MainData {
	return &mainStruct
}

// ReadJsonFileMain reads the input JSON file for main
// and stores the generic JSON Table as well as the
// decoded structs.
func ReadJsonFileMain(fn string) error {
	var err 		error
	var jsonPath 	string

	jsonPath,_ = filepath.Abs(fn)
	if sharedData.Debug() {
		log.Println("json path:", jsonPath)
	}

	// Read in the json file generically
	if mainJson, err = util.ReadJsonFile(jsonPath); err != nil {
		return errors.New(fmt.Sprintln("Error: unmarshalling", jsonPath, ", JSON input file:", err))
	}

	// Read in the json file structurally
	if err = util.ReadJsonFileToData(jsonPath, &mainStruct); err != nil {
		return errors.New(fmt.Sprintln("Error: unmarshalling", jsonPath, ", JSON input file:", err))
	}

	if sharedData.Debug() {
		log.Println("\tJson Data:", mainJson)
		log.Println("\tJson Struct:", mainStruct)
	}

	return nil
}

func rowScan(mp map[string]interface{}, ) string {
	var str			strings.Builder
	var desc		string
	//var init		string
	//var internal	string
	//var name		string
	//fieldNames		:= []string{}

	//fmt.Println("\t\tflagPrs:", mp)

	// Get the field names
	for _, field := range mp {
		if sharedData.Debug() {
			v := reflect.ValueOf(field)
			log.Println("FIELD type:", v.Kind())
		}
		//fieldNames = append(fieldNames, field.Name.(string))
	}

	// Now create the string from the flag fields
	str.WriteString(",")
	if len(desc) > 0 {
		str.WriteString(fmt.Sprintf(",\"%s\"",desc))
	}
	str.WriteString(")")

	return str.String()
}

func Setup(inDefns map[string]interface{}) {

}

