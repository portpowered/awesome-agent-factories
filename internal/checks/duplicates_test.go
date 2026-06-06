package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		raw  string
		want string
	}{
		{
			raw:  "https://Example.com/foo/",
			want: "https://example.com/foo",
		},
		{
			raw:  "https://example.com/foo#section",
			want: "https://example.com/foo",
		},
	}

	for _, tt := range tests {
		if got := normalizeURL(tt.raw); got != tt.want {
			t.Fatalf("normalizeURL(%q) = %q, want %q", tt.raw, got, tt.want)
		}
	}
}

func resourceSectionsFixture(sectionBodies map[string]string) string {
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
		body, ok := sectionBodies[heading]
		if ok {
			b.WriteString(body)
		} else {
			b.WriteString("Optional intro without resource entries.\n")
		}
	}
	return strings.TrimSpace(b.String())
}

func TestValidateDuplicates(t *testing.T) {
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
			name: "distinct entries pass",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
- [CrewAI](https://github.com/joaomdmoura/crewAI) - A library for role-based agent crews.
`),
			}),
			wantFail: false,
		},
		{
			name: "duplicate URL across sections",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
`),
				"Benchmarks": strings.TrimSpace(`
- [AutoGen Benchmarks](https://github.com/microsoft/autogen) - Benchmarks for multi-agent orchestration flows.
`),
			}),
			wantFail:  true,
			wantRules: []string{"duplicate-url"},
		},
		{
			name: "duplicate URL with normalization differences",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [AutoGen](https://Example.com/foo/) - A framework for multi-agent applications.
- [AutoGen Mirror](https://example.com/foo#details) - A mirrored homepage for the same project.
`),
			}),
			wantFail:  true,
			wantRules: []string{"duplicate-url"},
		},
		{
			name: "duplicate resource name case-insensitively",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
- [autogen](https://example.com/autogen-docs) - Documentation for the AutoGen project.
`),
			}),
			wantFail:  true,
			wantRules: []string{"duplicate-name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := validateDuplicates(makeDoc(tt.markdown))
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
