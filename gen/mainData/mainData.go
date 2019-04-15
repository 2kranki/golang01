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

// genFlagVar generates the flag.~Var definition for given
// CLI variable definition
func genFlagVar(flg MainFlag) string {
	var str			strings.Builder

	// Now create the string from the names
	switch flg.FlagType {
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
		if flg.FlagType == "" {
			str.WriteString("flagType: EMPTY ")
		} else {
			str.WriteString("flagType:")
			str.WriteString(flg.FlagType)
		}
	}
	if len(flg.Internal) > 0 {
		str.WriteString(flg.Internal)
	} else {
		str.WriteString(flg.Name)
	}
	str.WriteString(",")
	if len(flg.Init) > 0 {
		if flg.FlagType == "string" {
			str.WriteString(fmt.Sprintf("\"%s\"",flg.Init))
		} else {
			str.WriteString(flg.Init)
		}
	}
	if len(flg.Desc) > 0 {
		str.WriteString(fmt.Sprintf(",\"%s\"",flg.Desc))
	}
	str.WriteString(")\n")

	return str.String()
}

// GenFlagSetup generates the flag.~Var definitions for the
// CLI variables
func GenFlagSetup() string {
	s := ""
	for _, v := range mainStruct.Flags {
		s += genFlagVar(v)
	}
	return s
}

// GenVarDefns generate the CLI variable definitions
func GenVarDefns() string {
	s := "\t"
	for _, v := range mainStruct.Flags {
		if len(v.Internal) > 0 {
			s += fmt.Sprintf("%s\t", v.Internal)
		} else {
			s += fmt.Sprintf("%s\t", v.Name)
		}
		s += fmt.Sprintf("%s\n", v.FlagType)
	}
	return s
}

// init() adds the functions needed for templating to
// shared data.
func init() {
	sharedData.SetFunc("GenFlagSetup", GenFlagSetup)
	sharedData.SetFunc("GenVarDefns", GenVarDefns)
}

func MainJson() interface{} {
	return mainJson
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

