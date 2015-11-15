package main

import (
	"os"
	"testing"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestMain(m *testing.M) {

	// setup

	os.Exit(m.Run())

	// teardown
}

func TestBadDir(t *testing.T) {

	i := ""
	flagDir = &i

	run := ""
	flagRun = &run

	match := ""
	flagMatch = &match

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	PanicMatches(t, func() { main() }, "**invalid Input Directory")

	i = "."
	flagDir = &i

	PanicMatches(t, func() { main() }, "**invalid Input Directoy '.'")
}

func TestGenerateDir(t *testing.T) {

	pwd, _ := os.Getwd()

	i := pwd
	flagDir = &i

	run := ""
	flagRun = &run

	match := ""
	flagMatch = &match

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	main()
}

func TestGenerateByGOPATH(t *testing.T) {

	i := "$GOPATH"
	flagDir = &i

	run := ""
	flagRun = &run

	match := ""
	flagMatch = &match

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	main()
}

func TestMatchGenerate(t *testing.T) {

	pwd, _ := os.Getwd()

	i := pwd
	flagDir = &i

	run := ""
	flagRun = &run

	match := "dir"
	flagMatch = &match

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	main()
}

func TestBadMatchGenerate(t *testing.T) {

	pwd, _ := os.Getwd()

	i := pwd
	flagDir = &i

	run := ""
	flagRun = &run

	match := "([efferf"
	flagMatch = &match

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	PanicMatches(t, func() { main() }, "**Error Compiling match Regex:error parsing regexp: missing closing ]: `[efferf`")
}

func TestIgnoreGenerate(t *testing.T) {

	pwd, _ := os.Getwd()

	i := pwd
	flagDir = &i

	run := ""
	flagRun = &run

	match := ""
	flagMatch = &match

	ignore := "generate"
	flagIgnore = &ignore

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	main()
}

func TestBadIgnoreGenerate(t *testing.T) {

	pwd, _ := os.Getwd()

	i := pwd
	flagDir = &i

	run := ""
	flagRun = &run

	match := ""
	flagMatch = &match

	ignore := "[(rftt"
	flagIgnore = &ignore

	v := false
	flagPrintProcessed = &v

	n := false
	flagPrintSimulate = &n

	x := false
	flagPrintExecute = &x

	PanicMatches(t, func() { main() }, "**Error Compiling ignore Regex:error parsing regexp: missing closing ]: `[(rftt`")
}

func TestGenerateArgs(t *testing.T) {

	pwd, _ := os.Getwd()

	i := pwd
	flagDir = &i

	run := "statics.*"
	flagRun = &run

	match := ""
	flagMatch = &match

	ignore := ""
	flagIgnore = &ignore

	v := true
	flagPrintProcessed = &v

	n := true
	flagPrintSimulate = &n

	x := true
	flagPrintExecute = &x

	main()
}
