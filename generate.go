package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const (
	gopath = "$GOPATH"
	runArg = "run=%s"
	dollar = "$"
)

var (
	flagDir            = flag.String("i", gopath, "Directory to recurse through, default "+gopath)
	flagRun            = flag.String("run", "", "if non-empty, specifies a regular expression to select directives whose full original source text (excluding any trailing spaces and final newline) matches the expression")
	flagMatch          = flag.String("match", "", "Regexp for directories to run the command for")
	flagIgnore         = flag.String("ignore", "", "Regexp for dirs we should ignore")
	flagPrintProcessed = flag.Bool("v", false, "flag prints the names of packages and files as they are")
	flagPrintSimulate  = flag.Bool("n", false, "flag prints commands that would be executed")
	flagPrintExecute   = flag.Bool("x", false, "flag prints commands as they are executed")

	v, n, x, run string
	matchRegexp  *regexp.Regexp
	ignoreRegexp *regexp.Regexp
)

func main() {
	parseFlags()
	generate()
}

func parseFlags() {

	flag.Parse()

	var s string

	if len(*flagDir) == 0 {
		panic("**invalid Input Directory")
	}

	// ENV variable
	if (*flagDir)[:1] == dollar {
		s = filepath.Clean(os.Getenv((*flagDir)[1:]))
	} else {
		s = filepath.Clean(*flagDir)
	}

	flagDir = &s

	if *flagDir == "." {
		panic("**invalid Input Directoy '" + *flagDir + "'")
	}

	if len(*flagMatch) > 0 {

		var err error

		matchRegexp, err = regexp.Compile(*flagMatch)
		if err != nil {
			panic("**Error Compiling match Regex:" + err.Error())
		}
	}

	if len(*flagIgnore) > 0 {

		var err error

		ignoreRegexp, err = regexp.Compile(*flagIgnore)
		if err != nil {
			panic("**Error Compiling ignore Regex:" + err.Error())
		}
	}

	if *flagPrintExecute {
		x = "-x"
	}

	if *flagPrintSimulate {
		n = "-n"
	}

	if *flagPrintProcessed {
		v = "-v"
	}

	if len(*flagRun) > 0 {
		run = fmt.Sprintf(runArg, *flagRun)
	}
}

func generate() {

	walker := func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			return nil
		}

		if matchRegexp != nil && !matchRegexp.MatchString(path) {
			return nil
		}

		if ignoreRegexp != nil && ignoreRegexp.MatchString(path) {
			return nil
		}

		fmt.Println("Processing:", path)

		if err := os.Chdir(path); err != nil {
			log.Fatalf("\n**error changing DIR '%s'\n%s\n", path, err)

		}

		executeCmd("go", "generate", x, v, n, run)

		return nil
	}

	if err := filepath.Walk(*flagDir, walker); err != nil {
		log.Fatalf("\n**could not walk project path '%s'\n%s\n", *flagDir, err)
	}
}
