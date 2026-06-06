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

func collectCheckResults(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult
	results = append(results, validateRequiredSections(doc)...)
	results = append(results, validateResourceEntryFormat(doc)...)
	results = append(results, validateDescriptionRules(doc)...)
	results = append(results, validateDuplicates(doc)...)
	results = append(results, validateAlphabeticalOrder(doc)...)
	results = append(results, validateBareURLs(doc)...)
	results = append(results, validateTrackingURLs(doc)...)
	results = append(results, validateScopeKeywords(doc)...)
	return results
}

func reportResults(results []CheckResult) error {
	failureCount := 0
	warningCount := 0
	for _, result := range results {
		switch result.Level {
		case "failure":
			fmt.Fprintln(os.Stderr, result.Message)
			failureCount++
		case "warning":
			fmt.Fprintln(os.Stderr, result.Message)
			warningCount++
		}
	}
	if failureCount > 0 {
		return fmt.Errorf("%d check failure(s)", failureCount)
	}
	if warningCount > 0 {
		fmt.Fprintf(os.Stderr, "README checks passed with %d warning(s)\n", warningCount)
	} else {
		fmt.Fprintln(os.Stderr, "README checks passed")
	}
	return nil
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
	return reportResults(collectCheckResults(doc))
}
