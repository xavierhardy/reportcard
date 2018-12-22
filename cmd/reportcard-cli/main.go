package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/xavierhardy/reportcard/check"
)

var (
	dir     = flag.String("d", ".", "Root directory of your Go application")
	verbose = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	result, err := check.Run(*dir)
	if err != nil {
		log.Fatalf("Fatal error checking %s: %s", *dir, err.Error())
	}

	fmt.Printf("Grade: %s (%.1f%%)\n", result.Grade, result.Average*100)
	fmt.Printf("Files: %d\n", result.Files)
	fmt.Printf("Issues: %d\n", result.Issues)

	for _, c := range result.Checks {
		fmt.Printf("%s: %d%%\n", c.Name, int64(c.Percentage*100))
		if *verbose && len(c.FileSummaries) > 0 {
			for _, f := range c.FileSummaries {
				fmt.Printf("\t%s\n", f.Filename)
				for _, e := range f.Errors {
					fmt.Printf("\t\tLine %d: %s\n", e.LineNumber, e.ErrorString)
				}
			}
		}
	}
}
