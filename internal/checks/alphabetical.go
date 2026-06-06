package main

import (
	"fmt"
	"strings"
)

// CheckResult records a hard failure or non-blocking warning from a validation rule.
type CheckResult struct {
	Level   string // "failure" or "warning"
	Rule    string
	Message string
}

func collectSectionEntries(section Section) []ResourceEntry {
	var entries []ResourceEntry

	bodyLines := strings.Split(section.Body, "\n")
	for i, line := range bodyLines {
		if !looksLikeResourceEntry(line) {
			continue
		}

		lineNum := section.Line + 1 + i
		entry, ok := parseResourceEntry(line, section.Heading, lineNum)
		if ok {
			entries = append(entries, entry)
		}
	}

	return entries
}

func validateAlphabeticalOrder(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult

	for _, section := range doc.Sections {
		if !isResourceSection(section.Heading) {
			continue
		}

		entries := collectSectionEntries(section)
		if len(entries) <= 1 {
			continue
		}

		for i := 1; i < len(entries); i++ {
			prev := entries[i-1]
			curr := entries[i]
			if strings.Compare(strings.ToLower(prev.Name), strings.ToLower(curr.Name)) > 0 {
				results = append(results, CheckResult{
					Level: "failure",
					Rule:  "alphabetical-order",
					Message: fmt.Sprintf(
						`section %q: entries not in alphabetical order by link text: %q (line %d) and %q (line %d)`,
						section.Heading,
						prev.Name,
						prev.Line,
						curr.Name,
						curr.Line,
					),
				})
			}
		}
	}

	return results
}
