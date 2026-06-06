package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	path, err := readmePath()
	if err != nil {
		return err
	}
	doc, err := loadReadme(path)
	if err != nil {
		return err
	}

	var failures []CheckResult
	failures = append(failures, validateRequiredSections(doc)...)
	failures = append(failures, validateResourceEntryFormat(doc)...)
	failures = append(failures, validateDescriptionRules(doc)...)
	failures = append(failures, validateDuplicates(doc)...)
	failures = append(failures, validateAlphabeticalOrder(doc)...)

	for _, result := range failures {
		if result.Level == "failure" {
			fmt.Fprintln(os.Stderr, result.Message)
		}
	}
	if len(failures) > 0 {
		return fmt.Errorf("%d check failure(s)", len(failures))
	}
	return nil
}
