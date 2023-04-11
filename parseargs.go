package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	file       string
	ignoreDir  []string
	ignoreFile []string
}

func parseArgs() *options {
	o := &options{}

	var flagHelp bool
	flag.BoolVarP(&flagHelp, "help", "h", false, "Show help (This message) and exit")
	flag.StringArrayVarP(&o.ignoreDir, "ignore-dir", "", []string{}, "Ignore Directories to show")
	flag.StringArrayVarP(&o.ignoreFile, "ignore-dir", "", []string{}, "Ignore Files to show")

	flag.Parse()

	if flagHelp {
		putHelp(fmt.Sprintf("[%s] Version v%s", cmdName, version))
	}

	o.targetFile()

	return o
}

func (o *options) targetFile() {
	for _, arg := range flag.Args() {
		if o.file != "" {
			putHelp(fmt.Sprintf("Err: Wrong args. Unnecessary arg [%s]", arg))
		}
		if arg == "-" {
			continue
		}
		o.file = arg
	}

	if o.file == "" {
		putHelp("Err: Wrong args. You should specify a FILE")
	}
}
