package main

import (
	"fmt"
	"os"

	pt "github.com/bayashi/go-proptree"
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

	tree, err := fromLocal(o)
	if err != nil {
		return err
	}

	tree.RenderAsText(os.Stdout, pt.RenderTextDefaultOptions())

	return nil
}
