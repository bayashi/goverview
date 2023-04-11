package main

import (
	"fmt"
	"os"
)

const (
	cmdName string = "goverview"
	version string = "0.0.1"
)

func main() {
	err := run()
	if err != nil {
		putErr(fmt.Sprintf("Err: %s", err))
		os.Exit(exitErr)
	}
	os.Exit(exitOK)
}

func run() error {
	o := parseArgs()
	overview, err := buildOverview(o)
	if err != nil {
		return err
	}

	fmt.Println(overview)

	return nil
}
