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
	"os"
	"runtime"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug  bool
	pw     string
	port   string
	server string
	user   string
	db     string
)

func exec() {
	queryDB()
}

func queryDB() {
	connStr := "sqlserver://" + user + ":" + pw + "@" + server + ":" + port + "?database=" + db + "&connection+timeout=30"
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

    flag.Usage = usage
	flag.BoolVar(&debug, "debug", false, "enable debugging")
	flag.StringVar(&pw, "pw", "Passw0rd!", "the database password")
	flag.StringVar(&port, "port", "1402", "the database port")
	flag.StringVar(&server, "server", "localhost", "the database server")
	flag.StringVar(&user, "user", "sa", "the database user")
	flag.StringVar(&db, "db", "TeachDB", "the database name")
	flag.Parse()

    // Show some intersting stuff
    fmt.Println("Environ:",os.Environ())
    fmt.Println("GOARCH:",runtime.GOARCH)
    fmt.Println("GOOS:",runtime.GOOS)

	exec()
	//fmt.Println("Query Columns Error: ", err.Error())

}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n\tconsole01 [options]\n")
	fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nNotes:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "This program gives an example of accessing a Microsoft SQL Server\n")
	fmt.Fprintf(flag.CommandLine.Output(), "from a GO program.\n")
}

