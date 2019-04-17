// See License.txt in main repository directory

// Template Functions used in generation
// See the specific template files for how the functions
// and data are used.

package cobjData

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
	Name		string		`json:"Name,omitempty"`
	Type		string		`json:"Type,omitempty"`
	Desc		string		`json:"Desc,omitempty"`
	Len		    int		    `json:"Len,omitempty"`
	Offset		int		    `json:"Offset,omitempty"`
	ShiftAmt	int		    `json:"ShiftAmt,omitempty"`
}

type Database struct {
	Name	    string			`json:"Name,omitempty"`
    Fields		[]DbField       `json:"Fields,omitempty"`
}

func (t *Database) CreateInsertStr() string {

	insertStr := ""
	for _, v := range t.Fields {
		insertStr += v.Name + ","
	}
	if len(insertStr) > 0 {
		insertStr = insertStr[0:len(insertStr)-1]
	}
	return insertStr
}

var	cobjStruct	Database

func CObjStruct() *Database {
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


