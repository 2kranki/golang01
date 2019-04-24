// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// SQL Application main program

// Generated: 2019-04-24 11:09:33.44631 -0400 EDT m=+0.001906926

package main

import (
    "fmt"
	"github.com/gorilla/mux"
    "log"
	"net/http"
	"time"
	"./handlers"
)

func HndlrFavIcon(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
	    http.NotFound(w, r)
	}
    http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
}

func HndlrHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Set up main menu here..."))
}

func HndlrDebug(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./tmpl/form.html")
}

func HndlrDebugAdd(w http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
	    http.NotFound(w, req)
	}
    fmt.Fprintf(w, "Good Addition!")
}

func HndlrDebugDelete(w http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
	    http.NotFound(w, req)
	}
    fmt.Fprintf(w, "Good Deletion!")
}

func HndlrDebugNext(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" {
	    http.NotFound(w, req)
	}
    fmt.Fprintf(w, "Going to Next!")
}

func HndlrDebugPrev(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" {
	    http.NotFound(w, req)
	}
    fmt.Fprintf(w, "Going to Prev!")
}

func HndlrDebugReset(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Doing a Reset!")

}

func HndlrDebugUpdate(w http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
	    http.NotFound(w, req)
	}
    fmt.Fprintf(w, "Good Update!")
}

func exec() {

	r := mux.NewRouter()
	r.HandleFunc("/", mainIndex)
	r.HandleFunc("/favicon.ico", HndlrFavIcon)
	    r.HandleFunc("/debug", HndlrDebug)
	    r.HandleFunc("/debugAdd", HndlrDebugAdd)
	    r.HandleFunc("/debugDelete", HndlrDebugDelete)
	    r.HandleFunc("/debugNext", HndlrDebugNext)
	    r.HandleFunc("/debugPrev", HndlrDebugPrev)
	    r.HandleFunc("/debugReset", HndlrDebugReset)
	    r.HandleFunc("/debugUpdate", HndlrDebugUpdate)
	r.HandleFunc("/Customer",                    handlers.CustomerHndlrIndex)
	    r.HandleFunc("/Customer/show",               handlers.CustomerHndlrShow)
	    r.HandleFunc("/Customer/created",            handlers.CustomerHndlrCreated)
	    r.HandleFunc("/Customer/deleted",            handlers.CustomerHndlrDeleted)
	    r.HandleFunc("/Customer/next",               handlers.CustomerHndlrNext)
	    r.HandleFunc("/Customer/prev",               handlers.CustomerHndlrPrev)
	    r.HandleFunc("/Customer/updated",            handlers.CustomerHndlrUpdated)
	r.HandleFunc("/Vendor",                    handlers.VendorHndlrIndex)
	    r.HandleFunc("/Vendor/show",               handlers.VendorHndlrShow)
	    r.HandleFunc("/Vendor/created",            handlers.VendorHndlrCreated)
	    r.HandleFunc("/Vendor/deleted",            handlers.VendorHndlrDeleted)
	    r.HandleFunc("/Vendor/next",               handlers.VendorHndlrNext)
	    r.HandleFunc("/Vendor/prev",               handlers.VendorHndlrPrev)
	    r.HandleFunc("/Vendor/updated",            handlers.VendorHndlrUpdated)
	srvUrl := fmt.Sprintf("%s:%s", http_srvr, http_port)
	srv := &http.Server{
		Handler:      r,
		Addr:         srvUrl,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}


func mainIndex(w http.ResponseWriter, req *http.Request) {
	//http.Redirect(w, req, "/$v.Name", http.StatusSeeOther)
}

// Tell Client that we don't have the requested file
func noFile(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}