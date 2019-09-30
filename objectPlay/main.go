// vi:nu:et:sts=4 ts=4 sw=4
// See License.txt in main repository directory (Public Domain)

// Test Go's object capabilities and how it works.

// Actually, inheritacne works pretty much the same as in other
// Object Oriented Programming languages.  We have upward in-
// heritance as normal.  However, there is no way for a lower-
// level object to execute a upper-level function from a lower-
// level function. This can be done in other OOPs.

// To accomplish the lower-level doing upper-level function, we
// must careate a func within the lower-level object.

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
)

//---------------------------------------------------------------------
//								Object A
//---------------------------------------------------------------------

type objA struct {

}

func (o *objA) DoA() {
	log.Printf("objA::DoA\n")
}

func (o *objA) DoB() {
	log.Printf("objA::DoB\n")
}

func (o *objA) DoC() {
	log.Printf("objA::DoC\n")
}

func NewObjA() *objA {
	return &objA{}
}

//---------------------------------------------------------------------
//								Object B
//---------------------------------------------------------------------

type objB struct {
	objA
	Abc		func ()
}

func (o *objB) DoA() {
	log.Printf("objB::DoA\n")
	o.objA.DoA()
}

func (o *objB) DoD() {
	log.Printf("objB::DoD\n")
	if o.Abc != nil {
		o.Abc()
	}
}

func NewObjB() *objB {
	return &objB{}
}

//---------------------------------------------------------------------
//								Object C
//---------------------------------------------------------------------

type objC struct {
	objB
}

func (o *objC) DoD() {
	log.Printf("objC::DoD\n")
	o.objB.DoD()
}

func (o *objC) DoE() {
	log.Printf("objC::DoE\n")
}

func NewObjC() *objC {
	oc := &objC{}

	oc.Abc = func () {
		// This works because oc is remembered since this is an inline function.
		oc.DoE()
	}
	return oc
}

//=====================================================================
//						Main Support Functions
//=====================================================================

func usage() {
    	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
    	
	    fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
	    flag.PrintDefaults()
	    fmt.Fprintf(flag.CommandLine.Output(), "\nNotes:\n")
            
}

func main() {
	var oa	*objA
	var ob	*objB
	var oc	*objC

    // Set up flag variables

	flag.Usage = usage
    flag.BoolVar(&debug, "debug", true, "enable debugging")
	flag.BoolVar(&force, "force", true, "enable over-writes and deletions")
	flag.BoolVar(&force, "f", true, "enable over-writes and deletions")
	flag.BoolVar(&noop, "noop", true, "execute program, but do not make real changes")
	flag.BoolVar(&quiet, "quiet", true, "enable quiet mode")
	flag.BoolVar(&quiet, "q", true, "enable quiet mode")

    // Parse the flags and check them
	flag.Parse()
	if debug {
		log.Println("\tIn Debug Mode...")
	}

    //---------------------------------------------------------------------
    //							Various Tests
    //---------------------------------------------------------------------

    oa = NewObjA()
    ob = NewObjB()
    oc = NewObjC()

	log.Println("You should see: objA::DoA")
	oa.DoA()
    log.Println("You should see: objB::DoA then objA::DoA")
    ob.DoA()
	log.Println("You should see: objB::DoA then objA::DoA")
	oc.DoA()
	log.Println("You should see: objC::DoD then objB::DoD then objC::DoE")
	oc.DoD()

}


