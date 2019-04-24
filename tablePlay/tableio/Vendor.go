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

func VendorRowDelete() {

}

func VendorRowGet(r *http.Request) (interface{}, error) {
    key := r.FormValue("{{$t.PrimaryKey}}")
    if key == "" {
	    return data, errors.New("400. Bad Request.")
    }

}

func VendorRowInsert() {

}

func VendorRowUpdate() {

}

