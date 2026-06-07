package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// contentsExcludedSections matches awesome-lint sectionHeadingDenylist: these
// headings must exist but are omitted from the Contents table.
var contentsExcludedSections = map[string]bool{
	"Related Lists": true,
	"Contributing":  true,
}

// requiredResourceSections lists the ten Phase 1 README category headings.
var requiredResourceSections = []string{
	"Theories",
	"Coordination Patterns",
	"Frameworks",
	"Protocols and Interfaces",
	"Benchmarks",
	"Research Papers",
	"Blog Posts",
	"Case Studies",
	"Examples and Templates",
	"Related Lists",
}

var contentsLinkPattern = regexp.MustCompile(`^\s*-\s*\[([^\]]+)\]\((#[^)]+)\)\s*$`)

// Section represents a ## heading and its body content within the README.
type Section struct {
	Heading string
	Body    string
	Line    int
}

// parseSections extracts level-2 headings and the body text up to the next heading.
func parseSections(content string) []Section {
	lines := strings.Split(content, "\n")
	var sections []Section
	var current *Section

	flush := func() {
		if current == nil {
			return
		}
		current.Body = strings.TrimRight(current.Body, "\n")
		sections = append(sections, *current)
		current = nil
	}

	for i, line := range lines {
		if heading, ok := sectionHeading(line); ok {
			flush()
			current = &Section{
				Heading: heading,
				Line:    i + 1,
			}
			continue
		}
		if current != nil {
			if current.Body != "" {
				current.Body += "\n"
			}
			current.Body += line
		}
	}

	flush()
	return sections
}

func sectionByHeading(sections []Section, heading string) (Section, bool) {
	for _, section := range sections {
		if section.Heading == heading {
			return section, true
		}
	}
	return Section{}, false
}

func headingAnchor(heading string) string {
	var b strings.Builder
	lastWasSpace := false
	for _, r := range heading {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(unicode.ToLower(r))
			lastWasSpace = false
		case unicode.IsSpace(r) || r == '-':
			if !lastWasSpace && b.Len() > 0 {
				b.WriteByte('-')
				lastWasSpace = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}

type contentsLink struct {
	Text   string
	Anchor string
}

func parseContentsLinks(body string) []contentsLink {
	var links []contentsLink
	for _, line := range strings.Split(body, "\n") {
		matches := contentsLinkPattern.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		links = append(links, contentsLink{
			Text:   matches[1],
			Anchor: strings.TrimPrefix(matches[2], "#"),
		})
	}
	return links
}

func validateRequiredSections(doc *ReadmeDocument) []CheckResult {
	var results []CheckResult

	headings := make(map[string]Section, len(doc.Sections))
	for _, section := range doc.Sections {
		headings[section.Heading] = section
	}

	for _, required := range requiredResourceSections {
		if _, ok := headings[required]; !ok {
			results = append(results, CheckResult{
				Level:   "failure",
				Rule:    "required-section",
				Message: fmt.Sprintf("missing required section heading: ## %s", required),
			})
		}
	}

	contents, ok := sectionByHeading(doc.Sections, "Contents")
	if !ok {
		results = append(results, CheckResult{
			Level:   "failure",
			Rule:    "contents-section",
			Message: "missing required section heading: ## Contents",
		})
		return results
	}

	links := parseContentsLinks(contents.Body)
	linksByAnchor := make(map[string]contentsLink, len(links))
	for _, link := range links {
		linksByAnchor[link.Anchor] = link
	}

	for _, required := range requiredResourceSections {
		if contentsExcludedSections[required] {
			continue
		}
		expectedAnchor := headingAnchor(required)
		link, listed := linksByAnchor[expectedAnchor]
		if !listed {
			results = append(results, CheckResult{
				Level:   "failure",
				Rule:    "contents-alignment",
				Message: fmt.Sprintf("Contents missing anchor link for required section ## %s (expected #%s)", required, expectedAnchor),
			})
			continue
		}
		if _, ok := headings[required]; !ok {
			continue
		}
		if headingAnchor(required) != link.Anchor {
			results = append(results, CheckResult{
				Level:   "failure",
				Rule:    "contents-alignment",
				Message: fmt.Sprintf("Contents anchor #%s does not match section heading ## %s", link.Anchor, required),
			})
		}
	}

	for _, link := range links {
		matchesRequired := false
		for _, required := range requiredResourceSections {
			if headingAnchor(required) == link.Anchor {
				matchesRequired = true
				break
			}
		}
		if !matchesRequired {
			continue
		}
		if _, ok := headings[anchorToHeading(doc.Sections, link.Anchor)]; !ok {
			results = append(results, CheckResult{
				Level:   "failure",
				Rule:    "contents-alignment",
				Message: fmt.Sprintf("Contents anchor #%s does not match any ## heading", link.Anchor),
			})
		}
	}

	return results
}

func anchorToHeading(sections []Section, anchor string) string {
	for _, section := range sections {
		if headingAnchor(section.Heading) == anchor {
			return section.Heading
		}
	}
	return ""
}
