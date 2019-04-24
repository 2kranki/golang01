// See License.txt in main repository directory

// Template Functions used in generation
// See the specific template files for how the functions
// and data are used.

package appData

import (
	"../shared"
	"../util"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

type DbTypeConv struct {
	DbType		string
	GoType		string
}

var typeConvMsSql = []DbTypeConv{
	{"BLOB", "[]byte"},						// BLOB
	{"BOOLEAN", "???"},						// BOOLEAN
	{"CHAR", "???"},						// CHAR[(length)]
	{"CHARACTER", "???"},					// CHARACTER[(length)]
	{"CLOB", "[]byte"},						// CLOB
	{"DATE", "time.Time"},
	{"DATETIME", "time.Time"},
	{"DEC", "???"},							// DEC[(p[,s])]
	{"DECIMAL", "???"},						// DECIMAL[(p[,s])]
	{"FLOAT", "float64"},					// FLOAT(p)
	{"INT", "int64"},
	{"INTEGER", "int64"},
	{"NULL", "nil"},
	{"NUMERIC", "???"},						// NUMERIC [(p[,s])]
	{"NVARCHAR", "string"},					// NVARCHAR(length)
	{"REAL", "???"},
	{"SMALLINT", "???"},
	{"TEXT", "string"},
	{"TIME", "time.Time"},
	{"TIMESTAMP", "time.Time"},
	{"VARCHAR", "???"},						// VARCHAR(length)
}

var typeConvMySql = []DbTypeConv{
	{"BLOB", "[]byte"},						// BLOB
	{"BOOLEAN", "???"},						// BOOLEAN
	{"CHAR", "???"},						// CHAR[(length)]
	{"CHARACTER", "???"},					// CHARACTER[(length)]
	{"CLOB", "[]byte"},						// CLOB
	{"DATE", "time.Time"},
	{"DATETIME", "time.Time"},
	{"DEC", "???"},							// DEC[(p[,s])]
	{"DECIMAL", "???"},						// DECIMAL[(p[,s])]
	{"FLOAT", "float64"},					// FLOAT(p)
	{"INT", "int64"},
	{"INTEGER", "int64"},
	{"NULL", "nil"},
	{"NUMERIC", "???"},						// NUMERIC [(p[,s])]
	{"NVARCHAR", "string"},					// NVARCHAR(length)
	{"REAL", "???"},
	{"SMALLINT", "???"},
	{"TEXT", "string"},
	{"TIME", "time.Time"},
	{"TIMESTAMP", "time.Time"},
	{"VARCHAR", "???"},						// VARCHAR(length)
}

var typeConvPostgres = []DbTypeConv{
	{"BLOB", "[]byte"},						// BLOB
	{"BOOLEAN", "???"},						// BOOLEAN
	{"CHAR", "???"},						// CHAR[(length)]
	{"CHARACTER", "???"},					// CHARACTER[(length)]
	{"CLOB", "[]byte"},						// CLOB
	{"DATE", "time.Time"},
	{"DATETIME", "time.Time"},
	{"DEC", "???"},							// DEC[(p[,s])]
	{"DECIMAL", "???"},						// DECIMAL[(p[,s])]
	{"FLOAT", "float64"},					// FLOAT(p)
	{"INT", "int64"},
	{"INTEGER", "int64"},
	{"NULL", "nil"},
	{"NUMERIC", "???"},						// NUMERIC [(p[,s])]
	{"NVARCHAR", "string"},					// NVARCHAR(length)
	{"REAL", "???"},
	{"SMALLINT", "???"},
	{"TEXT", "string"},
	{"TIME", "time.Time"},
	{"TIMESTAMP", "time.Time"},
	{"VARCHAR", "???"},						// VARCHAR(length)
}

var typeConvSqlite = []DbTypeConv{
	{"BLOB", "[]byte"},						// BLOB
	{"BOOLEAN", "???"},						// BOOLEAN
	{"CHAR", "???"},						// CHAR[(length)]
	{"CHARACTER", "???"},					// CHARACTER[(length)]
	{"CLOB", "[]byte"},						// CLOB
	{"DATE", "time.Time"},
	{"DATETIME", "time.Time"},
	{"DEC", "???"},							// DEC[(p[,s])]
	{"DECIMAL", "???"},						// DECIMAL[(p[,s])]
	{"FLOAT", "float64"},					// FLOAT(p)
	{"INT", "int64"},
	{"INTEGER", "int64"},
	{"NULL", "nil"},
	{"NUMERIC", "???"},						// NUMERIC [(p[,s])]
	{"NVARCHAR", "???"},					// NVARCHAR(length)
	{"REAL", "???"},
	{"SMALLINT", "???"},
	{"TEXT", "string"},
	{"TIME", "time.Time"},
	{"TIMESTAMP", "time.Time"},
	{"VARCHAR", "???"},						// VARCHAR(length)
}

type DbField struct {
	Name		string		`json:"Name,omitempty"`			// Field Name
	Type		string		`json:"Type,omitempty"`			// SQL Type
	Len		    int		    `json:"Len,omitempty"`			// Data Maximum Length
	Dec		    int		    `json:"Dec,omitempty"`			// Decimal Positions
	PrimaryKey  bool	    `json:"PrimaryKey,omitempty"`
	List		bool	    `json:"List,omitempty"`			// Include in List Report
}

// DbTable stands for Database Table and defines
// the make up of the SQL Table.
// Fields should be in the order in which they are to
// be displayed in the list form and the main form.
type DbTable struct {
	Name		string		`json:"Name,omitempty"`
	Fields		[]DbField	`json:"Fields,omitempty"`
}

func (t *DbTable) CreateInsertStr() string {

	insertStr := ""
	for _, v := range t.Fields {
		insertStr += v.Name + ","
	}
	if len(insertStr) > 0 {
		insertStr = insertStr[0:len(insertStr)-1]
	}
	return insertStr
}

type Database struct {
	Name	string			`json:"Name,omitempty"`
	SqlType	string			`json:"SqlType,omitempty"`
	Tables  []DbTable		`json:"Tables,omitempty"`
}

var	appStruct	Database
var	appJson		interface{}

func AppJson() interface{} {
	return appJson
}

func AppStruct() *Database {
	return &appStruct
}

func CreateInsertSql(t interface{}) string {
	//var Fields 	[]map[string] interface{}
	//var ok		bool
	var x		string

	//insertStr := ""
	return x
}

func CreateTableStruct(t interface{}) string {
	//var Fields 	[]map[string] interface{}
	//var ok		bool
	var x		string

	//insertStr := ""
	return x
}

// decodeDbField decodes a DbField from a generic
func decodeDbField(i interface{}) *DbField {
	var f		DbField
	var m		map[string]interface{}
	var ok		bool

	if m, ok = i.(map[string]interface{}); !ok {
		return nil
	}
	if f.Name, ok = m["Name"].(string); !ok {
		return nil
	}
	if f.Type, ok = m["Type"].(string); !ok {
		return nil
	}
	if f.Len, ok = m["Len"].(int); !ok {
		return nil
	}
	if f.PrimaryKey, ok = m["PrimaryKey"].(bool); !ok {
		return nil
	}
	if f.List, ok = m["List"].(bool); !ok {
		return nil
	}
	return &f
}

func decodeDbTable(i interface{}) *DbTable {
	var t		DbTable
	var m		map[string]interface{}
	//var fields	[]interface{}
	var fields	[]DbField
	var ok		bool

	if m, ok = i.(map[string]interface{}); !ok {
		return nil
	}
	if t.Name, ok = m["Name"].(string); !ok {
		return nil
	}
	if fields, ok = m["Fields"].([]DbField); !ok {
		return nil
	}
	for _, v := range fields {
		t.Fields = append(t.Fields, *decodeDbField(v))
	}
	return &t
}

func decodeDatabase(i interface{}) *Database {
	var t		Database
	var m		map[string]interface{}
	//var fields	[]interface{}
	var tables	[]DbTable
	var ok		bool

	if m, ok = i.(map[string]interface{}); !ok {
		return nil
	}
	if t.Name, ok = m["Name"].(string); !ok {
		return nil
	}
	if t.SqlType, ok = m["SqlType"].(string); !ok {
		return nil
	}
	if tables, ok = m["Tables"].([]DbTable); !ok {
		return nil
	}
	for _, v := range tables {
		t.Tables = append(t.Tables, *decodeDbTable(v))
	}
	return &t
}

func ForTables(f func(DbTable)) {
	for _, v := range appStruct.Tables {
		f(v)
	}
}

func GenAccessFunc(t DbTable) string {
	var str			strings.Builder
	str.WriteString(fmt.Sprintf("\tfunc %sDeleteRow( ) {\n", t.Name))
	str.WriteString("\t}\n\n")
	str.WriteString(fmt.Sprintf("\tfunc %sInsertRow( ) {\n", t.Name))
	str.WriteString("\t}\n\n")
	str.WriteString(fmt.Sprintf("\tfunc %sSelect(sel string) ([]string, error) {\n", t.Name))
	/***
	  func {{title $t.Name }}Select(sel string) []string,error {
	      {{ if eq .Data.SqlType "mariadb" }}
	          ERROR - NOT IMPLEMENTED
	      {{ else if eq .Data.SqlType "mssql" }}
	      _ "github.com/2kranki/go-mssqldb"
	      {{ else if eq .Data.SqlType "mysql" }}
	          _ "github.com/go-sql-driver/mysql"
	      {{ else if eq .Data.SqlType "postgres" }}
	          _ "github.com/lib/pq"
	      {{ else if eq .Data.SqlType "sqlite" }}
	      _ "github.com/2kranki/go-sqlite3"
	      {{ end }}

	  }

	 */
	str.WriteString("\t}\n\n")
	str.WriteString(fmt.Sprintf("\tfunc %sSetupRow(r *http.Request) {\n", t.Name))
	/***
	    func {{ title $t.Name }}SetupRow(r *http.Request) {
	        data := interface{}
	        key := r.FormValue("{{$t.PrimaryKey}}")
		    if key == "" {
			    return data, errors.New("400. Bad Request.")
		    }
	        row := config.DB.QueryRow("SELECT * FROM {{$t.Name}} WHERE {{$t.PrimaryKey}} = $1", key)
	        err := row.Scan(
	                    &data.Isbn,
	                    &data.Title,
	                    &data.Author,
	                    &data.Price)
	        if err != nil {
	        	return data, err
	        }
	        	return data, nil
	    }
	*/
	str.WriteString("\t}\n\n")
	str.WriteString(fmt.Sprintf("\tfunc %sUpdateRow( ) {\n", t.Name))
	str.WriteString("\t}\n\n")
	return str.String()
}

func GenAccessFuncs() string {
	var str			strings.Builder
	for _, v := range appStruct.Tables {
		str.WriteString(GenAccessFunc(v))
	}
	return str.String()
}

func GenForm(t DbTable) string {
	var str			strings.Builder
	str.WriteString(fmt.Sprintf("<form id=\"%s\" method=\"post\">\n", t.Name))
	for _, v := range t.Fields {
		str.WriteString(fmt.Sprintf("<p>%s</p>\n",v.Name))
	}
	str.WriteString("<p/>\n<p/>\n<p/>\n")
	str.WriteString("\t<input type=submit onclick='onPrev()' value=\"Prev\">\n")
	str.WriteString("\t<input type=submit onclick='onAdd()' value=\"Add\">\n")
	str.WriteString("\t<input type=submit onclick='onDelete()' value=\"Delete\">\n")
	str.WriteString("\t<input type=submit onclick='onUpdate()' value=\"Update\">\n")
	str.WriteString("\t<input type=submit onclick='onNext()' value=\"Next\">\n")
	str.WriteString("\t<input type=reset onclick='onReset()' value=\"Reset\">\n")
	str.WriteString("</form>\n\n")
	str.WriteString("<script>\n")
	str.WriteString("\tfunction onAdd() {\n")
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").action = \"/Create/Process\";\n", t.Name))
	str.WriteString("\t}\n")
	str.WriteString("\tfunction onDelete() {\n")
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").action = \"/Delete/Process\";\n",t.Name))
	str.WriteString("\t}\n")
	str.WriteString("\tfunction onNext() {\n")
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").action = \"/Next\";\n",t.Name))
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").method = \"get\";\n",t.Name))
	str.WriteString("\t}\n")
	str.WriteString("\tfunction onPrev() {\n")
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").action = \"/Prev\";\n",t.Name))
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").method = \"get\";\n",t.Name))
	str.WriteString("\t}\n")
	str.WriteString("\tfunction onReset() {\n")
	str.WriteString("\t}\n")
	str.WriteString("\tfunction onUpdate() {\n")
	str.WriteString(fmt.Sprintf("\t\tdocument.getElementById(\"%s\").action = \"/Update/Process\";\n",t.Name))
	str.WriteString("\t}\n")
	str.WriteString("</script>\n")
	return str.String()
}

func getTypeConv(db string) []DbTypeConv {
	switch db {
	case "mariadb":
		return typeConvMsSql
	case "mssql":
		return typeConvMsSql
	case "mysql":
		return typeConvMySql
	case "postgres":
		return typeConvMySql
	case "sqlite":
		return typeConvSqlite
	}
	return nil
}

// init() adds the functions needed for templating to
// shared data.
func init() {
	//sharedData.SetFunc("GenFlagSetup", GenFlagSetup)
	sharedData.SetFunc("GenAccessFuncs", GenAccessFuncs)
}

// ReadJsonFileApp reads the input JSON file for app
// and stores the generic JSON Table as well as the
// decoded structs.
func ReadJsonFileApp(fn string) error {
	var err		    error
	var jsonPath	string

	jsonPath,_ = filepath.Abs(fn)
	if sharedData.Debug() {
		log.Println("json path:", jsonPath)
	}

	// Read in the json file generically
	if appJson, err = util.ReadJsonFile(jsonPath); err != nil {
		return errors.New(fmt.Sprintln("Error: unmarshalling", jsonPath, ", JSON input file:", err))
	}

	// Read in the json file structurally
	if err = util.ReadJsonFileToData(jsonPath, &appStruct); err != nil {
		return errors.New(fmt.Sprintln("Error: unmarshalling", jsonPath, ", JSON input file:", err))
	}

	if sharedData.Debug() {
		log.Println("\tJson Data:", appJson)
		log.Println("\tJson Struct:", appStruct)
	}

	return nil
}

func TableNames() []string {
	var list	[]string

	for _, v := range appStruct.Tables {
		list = append(list, v.Name)
	}

	return list
}
