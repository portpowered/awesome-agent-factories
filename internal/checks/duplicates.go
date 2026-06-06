package main

import (
	"fmt"
	"net/url"
	"strings"
)

// normalizeURL prepares a resource URL for duplicate comparison.
func normalizeURL(raw string) string {
	trimmed := strings.TrimSpace(raw)
	parsed, err := url.Parse(trimmed)
	if err != nil || parsed.Host == "" {
		return trimmed
	}

	parsed.Scheme = strings.ToLower(parsed.Scheme)
	parsed.Host = strings.ToLower(parsed.Host)
	parsed.Fragment = ""
	parsed.Path = strings.TrimSuffix(parsed.Path, "/")

	normalized := url.URL{
		Scheme:   parsed.Scheme,
		Host:     parsed.Host,
		Path:     parsed.Path,
		RawQuery: parsed.RawQuery,
	}
	return normalized.String()
}

func collectResourceEntries(doc *ReadmeDocument) []ResourceEntry {
	var entries []ResourceEntry

	for _, section := range doc.Sections {
		if !isResourceSection(section.Heading) {
			continue
		}

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
	}

	return entries
}

func validateDuplicates(doc *ReadmeDocument) []CheckResult {
	entries := collectResourceEntries(doc)
	var results []CheckResult

	urlIndex := make(map[string][]ResourceEntry)
	for _, entry := range entries {
		key := normalizeURL(entry.URL)
		urlIndex[key] = append(urlIndex[key], entry)
	}
	for normalized, group := range urlIndex {
		if len(group) < 2 {
			continue
		}
		first := group[0]
		for _, other := range group[1:] {
			results = append(results, CheckResult{
				Level: "failure",
				Rule:  "duplicate-url",
				Message: fmt.Sprintf(
					`duplicate URL %q: section %q entry %q (line %d) and section %q entry %q (line %d)`,
					normalized,
					first.Section,
					first.Name,
					first.Line,
					other.Section,
					other.Name,
					other.Line,
				),
			})
		}
	}

	nameIndex := make(map[string][]ResourceEntry)
	for _, entry := range entries {
		key := strings.ToLower(strings.TrimSpace(entry.Name))
		nameIndex[key] = append(nameIndex[key], entry)
	}
	for _, group := range nameIndex {
		if len(group) < 2 {
			continue
		}
		first := group[0]
		for _, other := range group[1:] {
			results = append(results, CheckResult{
				Level: "failure",
				Rule:  "duplicate-name",
				Message: fmt.Sprintf(
					`duplicate resource name %q: section %q entry %q (line %d) and section %q entry %q (line %d)`,
					first.Name,
					first.Section,
					first.Name,
					first.Line,
					other.Section,
					other.Name,
					other.Line,
				),
			})
		}
	}

	return results
}
