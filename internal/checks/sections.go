package main

import "strings"

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
