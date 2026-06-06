package main

import (
	"strings"
	"testing"
)

func TestValidateAlphabeticalOrder(t *testing.T) {
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
			name: "in-order entries pass",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
- [CrewAI](https://github.com/joaomdmoura/crewAI) - A library for role-based agent crews.
- [LangGraph](https://github.com/langchain-ai/langgraph) - A graph-based agent orchestration library.
`),
			}),
			wantFail: false,
		},
		{
			name: "out-of-order entries fail",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [CrewAI](https://github.com/joaomdmoura/crewAI) - A library for role-based agent crews.
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
`),
			}),
			wantFail:  true,
			wantRules: []string{"alphabetical-order"},
		},
		{
			name: "single-entry section passes",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [AutoGen](https://github.com/microsoft/autogen) - A framework for multi-agent applications.
`),
			}),
			wantFail: false,
		},
		{
			name: "empty section passes",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": "Optional intro without resource entries.\n",
			}),
			wantFail: false,
		},
		{
			name: "case-insensitive alphabetical order",
			markdown: resourceSectionsFixture(map[string]string{
				"Frameworks": strings.TrimSpace(`
- [beta tool](https://example.com/beta) - A beta orchestration toolkit.
- [Alpha Tool](https://example.com/alpha) - An alpha orchestration toolkit.
`),
			}),
			wantFail:  true,
			wantRules: []string{"alphabetical-order"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := validateAlphabeticalOrder(makeDoc(tt.markdown))
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
