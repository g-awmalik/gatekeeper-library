package main

import (
	"flag"
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"
	"gatekeeperlibrary/cmd/pkg/ingestor/suite"
	"gatekeeperlibrary/cmd/pkg/presenter"
	"os"
	"sort"
)

var (
	output = flag.String("output", "", "a file to write to instead of stdout")
)

func main() {
	flag.Parse()

	roots := []string{"."}
	if flag.NArg() != 0 {
		roots = flag.Args()
	}

	// Walk the file tree, generating a slice of paths.  Paths are any
	// directory that contains a `template.yaml`.  The library allows a
	// starting directory to be passed, but defaults to the directory it's run
	// from.

	var docs []*apis.ConstraintTemplateDoc
	for _, root := range roots {
		var ds []*apis.ConstraintTemplateDoc
		var err error

		fmt.Println("root: ", root)

		ds, err = suite.Ingest(root)

		if err != nil {
			panic(fmt.Errorf("ingestion failed: %w", err))
		}

		fmt.Printf("DIRECTORIES: %v\n", len(ds))
		if len(ds) == 0 {
			panic(fmt.Sprintf("generated no documentation data from input dirs: %v", roots))
		}

		docs = append(docs, ds...)
	}

	// Lexicographically sort by name
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].Name < docs[j].Name
	})

	// Pass these structs to the presenter, receive text output in return
	text, err := presenter.Present(docs)
	if err != nil {
		panic(fmt.Sprintf("failed to Present: %s", err))
	}

	// write this text output to stdout or to a file if the flag includes that
	if *output == "" {
		fmt.Print(text)
	} else {
		err := os.WriteFile(*output, []byte(text), 0027)
		if err != nil {
			panic(err)
		}
	}
}
