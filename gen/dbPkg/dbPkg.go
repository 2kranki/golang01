// See License.txt in main repository directory

// dbPkg contains the data and functions to generate
// table and field data for html forms, handlers and
// table sql i/o for a specific database.  Multiple
// databases should be handled with multiple ??? of
// this package.

package dbPkg

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
	{"DEC", "float64"},						// DEC[(p[,s])]
	{"DECIMAL", "float64"},					// DECIMAL[(p[,s])]
	{"FLOAT", "float64"},					// FLOAT(p)
	{"INT", "int64"},
	{"INTEGER", "int64"},
	{"NULL", "nil"},
	{"NUMERIC", "???"},						// NUMERIC [(p[,s])]
	{"NVARCHAR", "string"},					// NVARCHAR(length)
	{"REAL", "float64"},
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
	{"DEC", "float64"},						// DEC[(p[,s])]
	{"DECIMAL", "float64"},					// DECIMAL[(p[,s])]
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
	{"DEC", "float64"},						// DEC[(p[,s])]
	{"DECIMAL", "float64"},					// DECIMAL[(p[,s])]
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
	{"DEC", "float64"},						// DEC[(p[,s])]
	{"DECIMAL", "float64"},					// DECIMAL[(p[,s])]
	{"FLOAT", "float64"},					// FLOAT(p)
	{"INT", "int64"},
	{"INTEGER", "int64"},
	{"NULL", "nil"},
	{"NUMERIC", "float64"},					// NUMERIC [(p[,s])]
	{"NVARCHAR", "string"},					// NVARCHAR(length)
	{"REAL", "float64"},
	{"SMALLINT", "int64"},
	{"TEXT", "string"},
	{"TIME", "time.Time"},
	{"TIMESTAMP", "time.Time"},
	{"VARCHAR", "string"},						// VARCHAR(length)
}

// DbField defines a Table's field mostly in terms of
// SQL.
type DbField struct {
	Name		string		`json:"Name,omitempty"`			// Field Name
	Type		string		`json:"Type,omitempty"`			// SQL Type
	Len		    int		    `json:"Len,omitempty"`			// Data Maximum Length
	Dec		    int		    `json:"Dec,omitempty"`			// Decimal Positions
	PrimaryKey  bool	    `json:"PrimaryKey,omitempty"`
	Nullable	bool		`json:"Null,omitempty"`
	List		bool	    `json:"List,omitempty"`			// Include in List Report
}

func (f *DbField) CreateSql(cm string) string {
	var str			strings.Builder
	var ft			string
	var nl			string
	var pk			string

	if f.Len > 0 {
		if f.Dec > 0 {
			ft = fmt.Sprintf("%s(%d,%d)", f.Type, f.Len, f.Dec)
		} else {
			ft = fmt.Sprintf("%s(%d)", f.Type, f.Len)
		}
	} else {
		ft = f.Type
	}
	nl = " NOT NULL"
	if f.Nullable {
		nl = ""
	}
	pk = ""
	if f.PrimaryKey {
		pk = " PRIMARY KEY"
	}
	str.WriteString(fmt.Sprintf("\t%s\t%s%s%s%s\n", f.Name, ft, nl, pk, cm))

	return str.String()
}

func (f *DbField) CreateStruct() string {
	var str			strings.Builder

	str.WriteString(fmt.Sprintf("\t%s\t", strings.Title(f.Name)))
	str.WriteString(fmt.Sprintf("%s\n", f.GoType()))

	return str.String()
}

func (f *DbField) GoType() string {
	return ConvFieldToGoType(dbStruct.SqlType, f.Type)
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

func (t *DbTable) CreateSql() string {
	var str			strings.Builder

	str.WriteString(fmt.Sprintf("DROP TABLE %s IF EXISTS;\n", t.Name))
	if dbStruct.SqlType == "mssql" {
		str.WriteString("GO\n")
	}
	str.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", t.Name))
	for i,f := range t.Fields {
		var cm  		string

		cm = ""
		if i != (len(t.Fields) - 1) {
			cm = ","
		}
		str.WriteString(f.CreateSql(cm))
	}
	str.WriteString(fmt.Sprintf(");\n"))
	if dbStruct.SqlType == "mssql" {
		str.WriteString("GO\n")
	}

	return str.String()
}

func (t *DbTable) CreateStruct( ) string {
	var str			strings.Builder

	str.WriteString(fmt.Sprintf("type %s struct {\n", t.TitledName))
	for i,_ := range t.Fields {
		str.WriteString(t.Fields[i].CreateStruct())
	}
	str.WriteString("}\n")

	return str.String()
}

func (t *DbTable) ForFields(f func(f *DbField) ) {
	for i,_ := range t.Fields {
		f(&t.Fields[i])
	}
}

func (t *DbTable) Form() string {
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

// PrimaryKey returns the first field that it finds
// that is marked as a primary key.
func (t *DbTable) PrimaryKey() *DbField {

	for i,_ := range t.Fields {
		if t.Fields[i].PrimaryKey {
			return &t.Fields[i]
		}
	}
	return nil
}

// ScanFields returns struct fields to be used in
// a row.Scan.  It assumes that the struct's name
// is "data"
func (t *DbTable) ScanFields() string {
	var str			strings.Builder

	for i,f := range t.Fields {
		cm := ","
		if i == len(t.Fields) - 1 {
			cm = ""
		}
		str.WriteString(fmt.Sprintf("&data.%s%s ", f.Name, cm))
	}
	return str.String()
}

func (t *DbTable) TitledName( ) string {
	return strings.Title(t.Name)
}

type Database struct {
	Name	string			`json:"Name,omitempty"`
	SqlType	string			`json:"SqlType,omitempty"`
	Server	string			`json:"Server,omitempty"`
	Port	string			`json:"Port,omitempty"`
	PW		string			`json:"PW,omitempty"`
	Tables  []DbTable		`json:"Tables,omitempty"`
}

func (d *Database) ForTables(f func(t *DbTable) ) {
	for i,_ := range d.Tables {
		f(&d.Tables[i])
	}
}

func (d *Database) TitledName( ) string {
	return strings.Title(d.Name)
}

var	dbStruct	Database

func DbStruct() *Database {
	return &dbStruct
}

func DefaultJsonFileName() string {
	return "db.json.txt"
}

func InsertSql(t interface{}) string {
	//var Fields 	[]map[string] interface{}
	//var ok		bool
	var x		string

	//insertStr := ""
	return x
}

func ForTables(f func(*DbTable)) {
	for i,_ := range dbStruct.Tables {
		f(&dbStruct.Tables[i])
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
	for _, v := range dbStruct.Tables {
		str.WriteString(GenAccessFunc(v))
	}
	return str.String()
}

func GenListField(f DbField) string {
	var str			strings.Builder

	if f.PrimaryKey {
		str.WriteString("<a href=\"\">")
	}

	return str.String()
}

func GenListBody(t *DbTable) string {
	var str			strings.Builder
	for _, v := range t.Fields {
		str.WriteString(GenListField(v))
	}
	return str.String()
}

// init() adds the functions needed for templating to
// shared data.
func init() {
	sharedData.SetFunc("GenAccessFuncs", GenAccessFuncs)
}

// ReadJsonFileDb reads the input JSON file for app
// and stores the generic JSON Table as well as the
// decoded structs.
func ReadJsonFile(fn string) error {
	var err		    error
	var jsonPath	string

	jsonPath,_ = filepath.Abs(fn)
	if sharedData.Debug() {
		log.Println("json path:", jsonPath)
	}

	// Read in the json file structurally
	if err = util.ReadJsonFileToData(jsonPath, &dbStruct); err != nil {
		return errors.New(fmt.Sprintln("Error: unmarshalling", jsonPath, ", JSON input file:", err))
	}

	if sharedData.Debug() {
		log.Println("\tJson Struct:", dbStruct)
	}

	return nil
}

func TableNames() []string {
	var list	[]string

	for _, v := range dbStruct.Tables {
		list = append(list, v.Name)
	}

	return list
}

func typeConv(dbType string) []DbTypeConv {
	var table  		[]DbTypeConv

	switch dbType {
	case "mariadb":
		table = typeConvMsSql
	case "mssql":
		table = typeConvMsSql
	case "mysql":
		table = typeConvMySql
	case "postgres":
		table = typeConvMySql
	case "sqlite":
		table = typeConvSqlite
	}

	return table
}

func ConvFieldToGoType(dbType string, ft string) string {
	var table  		[]DbTypeConv

	if table = typeConv(dbType); table == nil {
		return ""
	}
	for i, _ := range table {
		if strings.EqualFold(ft, table[i].DbType) {
			return table[i].GoType
		}
	}

	return ""
}

func ValidateData() error {

	if x := typeConv(dbStruct.SqlType); x == nil {
		return errors.New(fmt.Sprintf("SqlType of %s is not supported!",dbStruct.SqlType))
	}
	if dbStruct.Name == "" {
		return errors.New(fmt.Sprintf("Database Name is missing!"))
	}
	if len(dbStruct.Tables) == 0 {
		return errors.New(fmt.Sprintf("There are no tables defined for %s!", dbStruct.Name))
	}
	for i, t := range dbStruct.Tables {
		if t.Name == "" {
			return errors.New(fmt.Sprintf("%d Table Name is missing!", i))
		}
		if len(t.Fields) == 0 {
			return errors.New(fmt.Sprintf("There are no fields defined for %s!", t.Name))
		}
		for j,f := range t.Fields {
			if f.Name == "" {
				return errors.New(fmt.Sprintf("%d Field Name is missing from table %s!", j, t.Name))
			}
			if typ := ConvFieldToGoType(dbStruct.SqlType, f.Type); typ == "" {
				return errors.New(fmt.Sprintf("%s:%s Field Type is invalid!", t.Name, f.Name))
			}
		}
	}


	return nil
}
