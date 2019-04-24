// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//

package tableio

import (
    "database/sql"
	"errors"
	
	    _ "github.com/2kranki/go-sqlite3"
	
	"net/http"
	"strconv"
)


var db *sql.DB

// ioClose cleans up anything that needs to be
// accomplished before the database is closed
// and then closes the database connection.
func ioClose() {
    db.Close()
}

func ioConnect() {
	
	    db, err = sql.Open("sqlite3", host)
	
    if err != nil {
        log.Fatalln("Error: Cannot connect: ", err)
    }
    err = db.Ping()
    if err != nil {
        log.Fatalln("Error: Cannot connect: ", err)
    }
}

