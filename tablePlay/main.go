// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Generated: 2019-04-24 11:09:33.44631 -0400 EDT m=+0.001906926


package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

var (
	debug    	bool
	force    	bool
	noop     	bool
	quiet    	bool
	db_pw       string
	db_port     string
	db_srvr     string
	db_user     string
	db_name     string
	http_srvr   string
	http_port   string
	execPath	string	// exec json path (optional)
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

    // Set up flag variables

	flag.Usage = usage
flag.BoolVar(&debug, "debug", true, "enable debugging")
	flag.BoolVar(&force, "force", true, "enable over-writes and deletions")
	flag.BoolVar(&force, "f", true, "enable over-writes and deletions")
	flag.BoolVar(&noop, "noop", true, "execute program, but do not make real changes")
	flag.BoolVar(&quiet, "quiet", true, "enable quiet mode")
	flag.BoolVar(&quiet, "q", true, "enable quiet mode")
flag.StringVar(&execPath,"exec","","exec json path (optional)")

flag.StringVar(&db_pw, "dbpw", "Passw0rd!", "the database password")
	flag.StringVar(&db_port, "dbport", "1401", "the database port")
	flag.StringVar(&db_srvr, "dbserver", "localhost", "the database server")
	flag.StringVar(&db_user, "dbuser", "sa", "the database user")
	flag.StringVar(&db_name, "dbname", "", "the database name")
	flag.StringVar(&http_port, "port", "8080", "server port")
	flag.StringVar(&http_srvr, "server", "localhost", "server site")

    // Parse the flags and check them
	flag.Parse()
	if debug {
		log.Println("\tIn Debug Mode...")
	}

    // Execute the main process.
	exec()
}


