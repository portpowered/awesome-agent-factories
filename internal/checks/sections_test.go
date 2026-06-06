package main

import (
	"fmt"
	"strings"
	"testing"
)

func validEmptySectionFixture() string {
	var b strings.Builder
	b.WriteString("## Scope\n\nScope intro.\n\n## Contents\n\n")
	for _, heading := range requiredResourceSections {
		anchor := headingAnchor(heading)
		fmt.Fprintf(&b, "- [%s](#%s)\n", heading, anchor)
	}
	for _, heading := range requiredResourceSections {
		fmt.Fprintf(&b, "\n## %s\n\nOptional intro without resource entries.\n", heading)
	}
	return strings.TrimSpace(b.String())
}

func TestHeadingAnchor(t *testing.T) {
	tests := []struct {
		heading string
		want    string
	}{
		{"Theories", "theories"},
		{"Coordination Patterns", "coordination-patterns"},
		{"Protocols and Interfaces", "protocols-and-interfaces"},
		{"Examples and Templates", "examples-and-templates"},
	}

	for _, tt := range tests {
		if got := headingAnchor(tt.heading); got != tt.want {
			t.Fatalf("headingAnchor(%q) = %q, want %q", tt.heading, got, tt.want)
		}
	}
}

func TestValidateRequiredSections(t *testing.T) {
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
			name:     "valid empty sections",
			markdown: validEmptySectionFixture(),
			wantFail: false,
		},
		{
			name: "missing required section heading",
			markdown: strings.TrimSpace(`
## Contents

- [Theories](#theories)

## Theories

Intro only.
`),
			wantFail:  true,
			wantRules: []string{"required-section"},
		},
		{
			name: "missing contents section",
			markdown: strings.TrimSpace(`
## Theories

Intro only.
`),
			wantFail:  true,
			wantRules: []string{"required-section", "contents-section"},
		},
		{
			name: "missing contents entry for required section",
			markdown: strings.TrimSpace(`
## Contents

- [Theories](#theories)

## Theories

Intro only.

## Coordination Patterns

Intro only.

## Frameworks

Intro only.

## Protocols and Interfaces

Intro only.

## Benchmarks

Intro only.

## Research Papers

Intro only.

## Blog Posts

Intro only.

## Case Studies

Intro only.

## Examples and Templates

Intro only.

## Related Lists

Intro only.
`),
			wantFail:  true,
			wantRules: []string{"contents-alignment"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := validateRequiredSections(makeDoc(tt.markdown))
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

func filterFailures(results []CheckResult) []CheckResult {
	var failures []CheckResult
	for _, result := range results {
		if result.Level == "failure" {
			failures = append(failures, result)
		}
	}
	return failures
}

func containsRule(results []CheckResult, rule string) bool {
	for _, result := range results {
		if result.Rule == rule {
			return true
		}
	}
	return false
}
