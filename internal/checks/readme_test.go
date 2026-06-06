package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadReadmeMissing(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "README.md")

	_, err := loadReadme(path)
	if err == nil {
		t.Fatal("expected error for missing README")
	}
	if !strings.Contains(err.Error(), "README.md is missing") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseSections(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		want     []Section
	}{
		{
			name: "extracts level-2 headings and bodies",
			markdown: strings.TrimSpace(`
# Title

Preamble stays outside sections.

## Scope

Scope intro.

## Contents

- [Scope](#scope)

## Theories

Foundational ideas.
`),
			want: []Section{
				{Heading: "Scope", Body: "Scope intro.", Line: 5},
				{Heading: "Contents", Body: "- [Scope](#scope)", Line: 9},
				{Heading: "Theories", Body: "Foundational ideas.", Line: 13},
			},
		},
		{
			name:     "no sections",
			markdown: "# Only a title\n\nNo level-2 headings here.",
			want:     nil,
		},
		{
			name: "ignores level-3 headings as section boundaries",
			markdown: strings.TrimSpace(`
## Frameworks

Section intro.

### Subheading

Still part of Frameworks body.
`),
			want: []Section{
				{
					Heading: "Frameworks",
					Body:    "Section intro.\n\n### Subheading\n\nStill part of Frameworks body.",
					Line:    1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseSections(tt.markdown)
			if len(got) != len(tt.want) {
				t.Fatalf("got %d sections, want %d: %#v", len(got), len(tt.want), got)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Fatalf("section %d: got %#v, want %#v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestLoadReadmeParsesSections(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "README.md")
	content := strings.TrimSpace(`
## Alpha

First section.

## Beta

Second section.
`)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	doc, err := loadReadme(path)
	if err != nil {
		t.Fatalf("loadReadme: %v", err)
	}
	if len(doc.Sections) != 2 {
		t.Fatalf("got %d sections, want 2", len(doc.Sections))
	}
	if doc.Sections[0].Heading != "Alpha" || doc.Sections[1].Heading != "Beta" {
		t.Fatalf("unexpected headings: %#v", doc.Sections)
	}
}
