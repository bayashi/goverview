package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	path    string
	ignore  []string
	showAll bool
}

func parseArgs() *options {
	o := &options{}

	var flagHelp bool
	flag.BoolVarP(&flagHelp, "help", "h", false, "Show help (This message) and exit")
	flag.StringArrayVarP(&o.ignore, "ignore", "", []string{}, "Ignore path to show if a given string would match")
	flag.BoolVarP(&o.showAll, "show-all", "a", false, "Show all stuff")

	flag.Parse()

	if flagHelp {
		putHelp(fmt.Sprintf("Version v%s", version))
	}

	o.targetFile()

	return o
}

func (o *options) targetFile() {
	for _, arg := range flag.Args() {
		if o.path != "" {
			putHelp(fmt.Sprintf("Err: Wrong args. Unnecessary arg [%s]", arg))
		}
		if arg == "-" {
			continue
		}
		o.path = arg
	}

	if o.path == "" {
		putHelp("Err: Wrong args. You should specify a directory path")
	}
}
