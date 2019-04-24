// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

//  Handle HTTP Events

// Generated: 2019-04-24 11:09:33.44631 -0400 EDT m=+0.001906926


package handlers

import (
	
	_ "github.com/2kranki/go-sqlite3"
	
    "html/template"
)


var Tmpls *template.Template

func Title(i interface{}) string {
    return "Title() - NOT Implemented"
}

func Body(i interface{}) string {
    return "Body() - NOT Implemented"
}

func init() {
    funcs := map[string]interface{}{"Title":Title, "Body":Body,}
	Tmpls = template.Must(template.New("tmpls").Funcs(funcs).ParseGlob("tmpl/*.tmpl"))
}

