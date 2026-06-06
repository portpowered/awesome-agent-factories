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

	var results []CheckResult
	results = append(results, validateRequiredSections(doc)...)
	results = append(results, validateResourceEntryFormat(doc)...)
	results = append(results, validateDescriptionRules(doc)...)
	results = append(results, validateDuplicates(doc)...)
	results = append(results, validateAlphabeticalOrder(doc)...)
	results = append(results, validateBareURLs(doc)...)
	results = append(results, validateTrackingURLs(doc)...)
	results = append(results, validateScopeKeywords(doc)...)

	failureCount := 0
	for _, result := range results {
		switch result.Level {
		case "failure":
			fmt.Fprintln(os.Stderr, result.Message)
			failureCount++
		case "warning":
			fmt.Fprintln(os.Stderr, result.Message)
		}
	}
	if failureCount > 0 {
		return fmt.Errorf("%d check failure(s)", failureCount)
	}
	return nil
}
