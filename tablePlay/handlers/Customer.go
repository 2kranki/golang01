// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Generated: 2019-04-24 11:09:33.44631 -0400 EDT m=+0.001906926


package handlers

import (
	
	_ "github.com/2kranki/go-sqlite3"
	
    //"html/template"
	"net/http"
)

    // CustomerHndlrIndex handles the display of the table index.
    func CustomerHndlrIndex(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // for all rows {
            // Get the row to display.
            // Display the row index fields.
        // }

        w.Write([]byte("Show table index here..."))
    }

    // CustomerHndlrShow handles displaying of the table row form display.
    func CustomerHndlrShow(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // Get the row to display.

        // Display the row in the form.

        w.Write([]byte("Show a particular record..."))
    }

    // CustomerHndlrCreated handles an add row request which comes from
    // the row display form.
    func CustomerHndlrCreated(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // Verify any fields that need it.

        // Insert the row of data given.

        // Display the row in the form.

        w.Write([]byte("Process the form data from the row creation form..."))
    }

    // CustomerHndlrDeleted handles an delete request which comes from
    // the row display form.
    func CustomerHndlrDeleted(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // Verify any fields that need it.

        // Delete the row with data given.

        // Display the next row in the form.

        w.Write([]byte("Process the form data from the row deletion form..."))
    }

    // CustomerHndlrNext handles an next request which comes from
    // the row display form and should display the next row from the
    // current one.
    func CustomerHndlrNext(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // Get the next row to display.

        // Display the row in the form.

        w.Write([]byte("Process the form data from the row deletion form..."))
    }

    // CustomerHndlrPrev handles an previous request which comes from
    // the row display form and should display the previous row from the
    // current one.
    func CustomerHndlrPrev(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // Get the previous row to display.

        // Display the row in the form.

        w.Write([]byte("Process the form data from the row deletion form..."))
    }

    // CustomerHndlrUpdated handles an update request which comes from
    // the row display form.
    func CustomerHndlrUpdated(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
            return
        }

        // Verify any fields that need it.

        // Delete the row with data given.

        // Display the next row in the form.

        w.Write([]byte("Process the form data from the row update form..."))
    }

