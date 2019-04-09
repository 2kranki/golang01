// See License.txt in main repository directory

// Template Functions used in generation

package genSqlApp

import (
	"fmt"
	"log"
	"reflect"
	"strings"
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

func nameUC() string {
	return strings.ToUpper(defns[nameId].(string))
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
		if debug {
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

