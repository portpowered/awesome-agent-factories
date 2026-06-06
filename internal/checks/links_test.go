package main

import (
	"strings"
	"testing"
)

func resourceSectionFixture(sectionHeading, body string) string {
	var b strings.Builder
	b.WriteString("## Scope\n\nScope intro.\n\n## Contents\n\n")
	for _, heading := range requiredResourceSections {
		anchor := headingAnchor(heading)
		b.WriteString("- [")
		b.WriteString(heading)
		b.WriteString("](#")
		b.WriteString(anchor)
		b.WriteString(")\n")
	}
	for _, heading := range requiredResourceSections {
		b.WriteString("\n## ")
		b.WriteString(heading)
		b.WriteString("\n\n")
		if heading == sectionHeading {
			b.WriteString(body)
		} else {
			b.WriteString("Optional intro without resource entries.\n")
		}
	}
	return strings.TrimSpace(b.String())
}

func TestParseResourceEntry(t *testing.T) {
	entry, ok := parseResourceEntry(
		`- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.`,
		"Frameworks",
		42,
	)
	if !ok {
		t.Fatal("expected valid entry to parse")
	}
	if entry.Name != "AutoGen" || entry.URL != "https://github.com/microsoft/autogen" {
		t.Fatalf("unexpected entry fields: %#v", entry)
	}
	if entry.Section != "Frameworks" || entry.Line != 42 {
		t.Fatalf("unexpected metadata: %#v", entry)
	}
}

func TestValidateResourceEntryFormat(t *testing.T) {
	makeDoc := func(markdown string) *ReadmeDocument {
		return &ReadmeDocument{
			Raw:      markdown,
			Sections: parseSections(markdown),
		}
	}

	tests := []struct {
		name      string
		markdown  string
		wantFail  bool
		wantRules []string
	}{
		{
			name: "valid entry with intro paragraph",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
Foundational orchestration frameworks.

- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
`)),
			wantFail: false,
		},
		{
			name: "missing description separator",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen)
`)),
			wantFail:  true,
			wantRules: []string{"entry-format"},
		},
		{
			name: "colon instead of hyphen separator",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen): A framework for multi-agent applications.
`)),
			wantFail:  true,
			wantRules: []string{"entry-format"},
		},
		{
			name: "missing markdown link",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- AutoGen - A framework for multi-agent applications.
`)),
			wantFail: false,
		},
		{
			name: "multiple invalid entry shapes",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen)
- [CrewAI](https://github.com/joaomdmoura/crewAI) -- A library for role-based agent crews.
`)),
			wantFail:  true,
			wantRules: []string{"entry-format"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := validateResourceEntryFormat(makeDoc(tt.markdown))
			failures := filterFailures(results)
			if tt.wantFail && len(failures) == 0 {
				t.Fatal("expected failures, got none")
			}
			if !tt.wantFail && len(failures) > 0 {
				t.Fatalf("expected no failures, got %#v", failures)
			}
			for _, rule := range tt.wantRules {
				if !containsRule(failures, rule) {
					t.Fatalf("expected failure with rule %q, got %#v", rule, failures)
				}
			}
		})
	}
}

func TestValidateDescriptionRules(t *testing.T) {
	makeDoc := func(markdown string) *ReadmeDocument {
		return &ReadmeDocument{
			Raw:      markdown,
			Sections: parseSections(markdown),
		}
	}

	tests := []struct {
		name      string
		markdown  string
		wantFail  bool
		wantRules []string
	}{
		{
			name: "valid entry passes description rules",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
`)),
			wantFail: false,
		},
		{
			name: "missing terminal period",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications
`)),
			wantFail:  true,
			wantRules: []string{"description-period"},
		},
		{
			name: "banned phrase in description",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A revolutionary framework for multi-agent applications.
`)),
			wantFail:  true,
			wantRules: []string{"banned-marketing-phrase"},
		},
		{
			name: "banned phrase allowed in link text only",
			markdown: resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [The Best AutoGen Fork](https://example.com/best-autogen) - A maintained fork for multi-agent orchestration flows.
`)),
			wantFail: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := validateDescriptionRules(makeDoc(tt.markdown))
			failures := filterFailures(results)
			if tt.wantFail && len(failures) == 0 {
				t.Fatal("expected failures, got none")
			}
			if !tt.wantFail && len(failures) > 0 {
				t.Fatalf("expected no failures, got %#v", failures)
			}
			for _, rule := range tt.wantRules {
				if !containsRule(failures, rule) {
					t.Fatalf("expected failure with rule %q, got %#v", rule, failures)
				}
			}
		})
	}
}
