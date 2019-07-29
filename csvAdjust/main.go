// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// CSV File Adjustment program
// This program provides a convenient way to add a field
// with a constant value or delete one or more fields
// from a csv.

// Generated: Mon May 20, 2019 21:42

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
    "io"
	"log"
	"net/http"
	"os"
    "strconv"
)

var (
	debug     bool
	force     bool
	noop      bool
	quiet     bool
	db_pw     string
	db_port   string
	db_srvr   string
	db_user   string
	db_name   string
	execPath  string // exec json path (optional)
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])

	fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nNotes:\n")

	fmt.Fprintf(flag.CommandLine.Output(), "'exec json' is a file that defines the command line parameters \n")
	fmt.Fprintf(flag.CommandLine.Output(), "so that you can set them and then execute gen with -x or -exec\n")
	fmt.Fprintf(flag.CommandLine.Output(), "option.\n\n")
}

func main() {
	var err error
	var rcdin []string
	var rcdout []string
	var fileIn *os.File
	var fileOut *os.File
	var cnt int


	// Set up flag variables

	flag.Usage = usage
	flag.BoolVar(&debug, "debug", true, "enable debugging")
	flag.BoolVar(&force, "force", true, "enable over-writes and deletions")
	flag.BoolVar(&force, "f", true, "enable over-writes and deletions")
	flag.BoolVar(&noop, "noop", true, "execute program, but do not make real changes")
	flag.BoolVar(&quiet, "quiet", true, "enable quiet mode")
	flag.BoolVar(&quiet, "q", true, "enable quiet mode")

	flag.StringVar(&db_pw, "dbPW", "Passw0rd!", "the database password")
	flag.StringVar(&db_port, "dbPort", "1433", "the database port")
	flag.StringVar(&db_srvr, "dbServer", "localhost", "the database server")
	flag.StringVar(&db_user, "dbUser", "sa", "the database user")
	flag.StringVar(&db_name, "dbName", "", "the database name")

	// Parse the flags and check them
	flag.Parse()
	if debug {
		log.Println("\tIn Debug Mode...")
	}

	// Create the csv reader.
	fileIn, err = os.Open(path)
	if err != nil {
		log.Printf("...end CustomerHndlrTableLoadCSV(Error:400) - %s\n", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	defer fileIn.Close()

	log.Printf("\tFile, %s, is open...\n", path)
	rdr := csv.NewReader(fileIn)

	log.Printf("\tAdjusting the data...\n")
	for {
		rcdin, err = rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			str := fmt.Sprintf("ERROR: Reading row %d from csv - %s\n", cnt, err.Error())
			w.Write([]byte(str))
			return
		}
        cnt++
    }
}
