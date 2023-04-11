package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	exitOK  int = 0
	exitErr int = 1
)

func putErr(message ...interface{}) {
	fmt.Fprintln(os.Stderr, message...)
}

func putUsage() {
	putErr(fmt.Sprintf("Usage: %s [OPTIONS] FILE", cmdName))
}

func putHelp(message string) {
	putErr(message)
	putUsage()
	putErr("Options:")
	flag.PrintDefaults()
	os.Exit(exitOK)
}
