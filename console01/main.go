// vi:nu:et:sts=4 ts=4 sw=4

// The purpose of this program is to learn about accessing data in our
// Docker MS Sql Server. I was able to use quite a bit of the examples
// in the go-mssqldb repository.

// See: https://github.com/denisenkom/go-mssqldb

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug  = flag.Bool("debug", false, "enable debugging")
	pw     = flag.String("pw", "Passw0rd!", "the database password")
	port   = flag.String("port", "1401", "the database port")
	server = flag.String("server", "localhost", "the database server")
	user   = flag.String("user", "sa", "the database user")
	db     = flag.String("db", "TeachDB", "the database name")
)

func queryDB() {
	connStr := "sqlserver://" + *user + ":" + *pw + "@" + *server + ":" + *port + "?database=" + *db + "&connection+timeout=30"
	fmt.Println("connStr:", connStr)
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dbo.Orders")
	if err != nil {
		fmt.Println("Query Error: ", err.Error())
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Query Columns Error: ", err.Error())
		return
	}
	if cols == nil {
		return
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()
	}
	if rows.Err() != nil {
		//return rows.Err()
	}
}

func printValue(pval *interface{}) {

	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}

func main() {

	flag.Parse()
	queryDB()
	//fmt.Println("Query Columns Error: ", err.Error())

}
