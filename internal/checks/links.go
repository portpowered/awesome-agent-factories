package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ResourceEntry is a parsed - [Name](URL) - Description. list item in a resource section.
type ResourceEntry struct {
	Name        string
	URL         string
	Description string
	Section     string
	Line        int
}

var resourceEntryPattern = regexp.MustCompile(`^\s*-\s+\[([^\]]+)\]\(([^)]+)\)\s+-\s+(.+)$`)

var markdownLinkPattern = regexp.MustCompile(`\[[^\]]*\]\([^)]*\)`)
var bareURLPattern = regexp.MustCompile(`https?://[^\s\])<>"]+`)

var trackingURLPatterns = []string{"utm_", "fbclid", "gclid", "mc_eid"}

func isResourceSection(heading string) bool {
	for _, required := range requiredResourceSections {
		if required == heading {
			return true
		}
	}
	return false
}

func looksLikeResourceEntry(line string) bool {
	trimmed := strings.TrimLeft(line, " \t")
	return strings.HasPrefix(trimmed, "- [")
}

func parseResourceEntry(line string, section string, lineNum int) (ResourceEntry, bool) {
	matches := resourceEntryPattern.FindStringSubmatch(line)
	if matches == nil {
		return ResourceEntry{}, false
	}
	return ResourceEntry{
		Name:        matches[1],
		URL:         matches[2],
		Description: matches[3],
		Section:     section,
		Line:        lineNum,
	}, true
}

func validateResourceEntryFormat(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult

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
			if _, ok := parseResourceEntry(line, section.Heading, lineNum); ok {
				continue
			}

			results = append(results, CheckResult{
				Level: "failure",
				Rule:  "entry-format",
				Message: fmt.Sprintf(
					`section %q line %d: invalid resource entry format; expected - [Name](URL) - Description.`,
					section.Heading,
					lineNum,
				),
			})
		}
	}

	return results
}

func validateDescriptionRules(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult

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
			if !ok {
				continue
			}

			if !descriptionEndsWithPeriod(entry.Description) {
				results = append(results, CheckResult{
					Level: "failure",
					Rule:  "description-period",
					Message: fmt.Sprintf(
						`section %q line %d entry %q: description must end with a period`,
						entry.Section,
						entry.Line,
						entry.Name,
					),
				})
			}

			if phrase, found := findBannedPhraseInDescription(entry.Description); found {
				results = append(results, CheckResult{
					Level: "failure",
					Rule:  "banned-marketing-phrase",
					Message: fmt.Sprintf(
						`section %q line %d entry %q: description contains banned marketing phrase %q`,
						entry.Section,
						entry.Line,
						entry.Name,
						phrase,
					),
				})
			}
		}
	}

	return results
}

func lineWithoutMarkdownLinks(line string) string {
	return markdownLinkPattern.ReplaceAllString(line, "")
}

func findBareURLsInLine(line string) []string {
	return bareURLPattern.FindAllString(lineWithoutMarkdownLinks(line), -1)
}

func findTrackingPatternInURL(rawURL string) (string, bool) {
	lower := strings.ToLower(rawURL)
	for _, pattern := range trackingURLPatterns {
		if strings.Contains(lower, pattern) {
			return pattern, true
		}
	}
	return "", false
}

func validateBareURLs(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult

	for _, section := range doc.Sections {
		if !isResourceSection(section.Heading) {
			continue
		}

		bodyLines := strings.Split(section.Body, "\n")
		for i, line := range bodyLines {
			bareURLs := findBareURLsInLine(line)
			if len(bareURLs) == 0 {
				continue
			}

			lineNum := section.Line + 1 + i
			for _, bareURL := range bareURLs {
				results = append(results, CheckResult{
					Level: "failure",
					Rule:  "bare-url",
					Message: fmt.Sprintf(
						`section %q line %d: bare URL %q is not allowed; use - [Name](URL) - Description.`,
						section.Heading,
						lineNum,
						bareURL,
					),
				})
			}
		}
	}

	return results
}

func validateTrackingURLs(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult

	for _, entry := range collectResourceEntries(doc) {
		pattern, found := findTrackingPatternInURL(entry.URL)
		if !found {
			continue
		}

		results = append(results, CheckResult{
			Level: "warning",
			Rule:  "tracking-url",
			Message: fmt.Sprintf(
				`section %q line %d entry %q: URL contains tracking or suspicious pattern %q`,
				entry.Section,
				entry.Line,
				entry.Name,
				pattern,
			),
		})
	}

	return results
}
