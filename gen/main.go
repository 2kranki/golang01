// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory

// The purpose of this program is to generate other programs and code
// using Golang's templating system.

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	// genGo packages
	"./cobj"
	"./genSqlApp"
	"./util"
	// External Imports
)

const (
	cmdId      = "cmd"
	debugId    = "debug"
	defineId   = "defines"
	forceId    = "force"
	jsonDirId  = "jsondir"
	mdldirId   = "mdldir"
	nameId     = "name"
	noopId     = "noop"
	outdirId   = "outdir"
	quietId    = "quiet"
	timeId     = "time"
)

var (
	debug    bool
	execPath string
	force    bool
	jsonPath string
	mainPath string
	mdldir   string
	noop     bool
	outdir   string
	quiet    bool
)

var defns map[string]interface{}

type defineFlags []string

func (t *defineFlags) String() string {
	return ""
}

func (t *defineFlags) Set(value string) error {
	*t = append(*t, value)
	return nil
}

var defnFlags defineFlags

func SetupDefns(execPath string, cmd string) error {
	var jsonOut interface{}
	var wrk interface{}
	var err error
	//var flag		bool
	var ok bool

	// Set up default definitions
	defns = map[string]interface{}{}
	defns[debugId] = debug
	defns[forceId] = force
	if len(jsonPath) > 0 {
		defns[jsonDirId] = jsonPath
	}
	if len(mdldir) > 0 {
		defns[mdldirId] = mdldir
	} else {
		dir := os.ExpandEnv("${GENGOMDL}")
		if len(dir) > 0 {
			defns[mdldirId] = dir
		} else {
			defns[mdldirId] = "./models"
		}
	}
	defns[noopId] = noop
	if len(outdir) > 0 {
		defns[outdirId] = outdir
	}
	defns[quietId] = quiet
	defns[timeId] = time.Now().Format("Mon Jan _2, 2006 15:04")

	// Now merge in cli defines.
	for _, v := range defnFlags {
		s := strings.Split(v, "=")
		if len(s) > 1 {
			defns[s[0]] = s[1]
		}
	}
	defns[cmdId] = cmd

	if len(execPath) > 0 {
		jsonOut, err = util.ReadJsonFile(execPath)
		if err != nil {
			return err
		}
		if debug {
			fmt.Println("\tData:", jsonOut)
		}
		m := jsonOut.(map[string]interface{})
		if m == nil {
			return errors.New("Error: Exec JSON file did not unmarshal properly!")
		}
		if wrk, ok = m[debugId]; ok {
			defns[debugId] = wrk.(bool)
		}
		if wrk, ok = m[forceId]; ok {
			defns[forceId] = wrk.(bool)
		}
		if wrk, ok = m[quietId]; ok {
			defns[quietId] = wrk.(bool)
		}
		if wrk, ok = m[defineId]; ok {
			s := strings.Split(wrk.(string), ",")
			for _, v := range s {
				ss := strings.Split(v, "=")
				if len(ss) > 1 {
					defns[ss[0]] = ss[1]
				}
			}
		}
		if wrk, ok = m[cmdId]; ok {
			defns[cmdId] = wrk.(string)
		}
		if wrk, ok = m[jsonDirId]; ok {
			defns[jsonDirId] = wrk.(string)
		}
	}

	return nil
}

func main() {
	var err error

	flag.Usage = usage
	flag.BoolVar(&debug, "debug", true, "enable debugging")
	flag.StringVar(&execPath, "exec", "", "exec json path (optional)")
	flag.StringVar(&execPath, "x", "", "exec json path (optional)")
	flag.BoolVar(&force, "force", true, "enable over-writes and deletions")
	flag.BoolVar(&force, "f", true, "enable over-writes and deletions")
	flag.StringVar(&mainPath, "main", "", "set json main input path")
	flag.StringVar(&jsonPath, "json", "", "set json main input path")
	flag.StringVar(&mdldir, "mdldir", "", "set model input directory")
	flag.BoolVar(&noop, "noop", true, "execute program, but do not make real changes")
	flag.StringVar(&outdir, "outdir", "", "set output directory")
	flag.BoolVar(&quiet, "quiet", true, "enable quiet mode")
	flag.BoolVar(&quiet, "q", true, "enable quiet mode")
	flag.Var(&defnFlags, "define", "enter definitions (<name>=<string>)")
	flag.Var(&defnFlags, "d", "enter definitions (<name>=<string>)")
	flag.Parse()
	if debug {
		log.Println("\tIn Debug Mode...")
	}

	err = SetupDefns(execPath, flag.Arg(0))
	if err != nil {
		log.Fatalln("Error: failed to set up main definitions:", err)
	}

	// Execute the command
	if debug {
		log.Println("\tcmd: '",defns[cmdId],"'")
	}
	if debug {
		log.Println("\tjsonPath: '",defns[jsonDirId],"'")
	}
	switch defns[cmdId] {
	case "cobj":
		err = cobj.GenCObj(defns)
	case "sqlapp":
		err = genSqlApp.GenSqlApp(defns)
	default:
		fmt.Println("\nError: command must be 'cobj' or 'sqlapp'")
	}
	if err != nil {
		log.Println("Error: generation failed:", err)
	}

	if !quiet {
		log.Println("\tEnd of Program")
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n\tgen [options] (cobj | sqlapp)\n")
	fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nNotes:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "'exec json' is a file that defines the command line parameters \n")
	fmt.Fprintf(flag.CommandLine.Output(), "so that you can set them and then execute gen with -x or -exec\n")
	fmt.Fprintf(flag.CommandLine.Output(), "option.\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "'json path' is the json file that defines the data passed to the\n")
	fmt.Fprintf(flag.CommandLine.Output(), "template engine which controls data within the generated files.\n")
}