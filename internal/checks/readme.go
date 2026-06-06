package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// bannedMarketingPhrases lists promotional wording rejected in resource descriptions per CONTRIBUTING.md.
var bannedMarketingPhrases = []string{
	"revolutionary",
	"game-changing",
	"best",
	"ultimate",
	"cutting-edge",
}

var bannedMarketingPhrasePatterns []*regexp.Regexp

func initBannedMarketingPhrasePatterns() {
	for _, phrase := range bannedMarketingPhrases {
		escaped := regexp.QuoteMeta(phrase)
		bannedMarketingPhrasePatterns = append(
			bannedMarketingPhrasePatterns,
			regexp.MustCompile(`(?i)\b`+escaped+`\b`),
		)
	}
}

func descriptionEndsWithPeriod(description string) bool {
	return strings.HasSuffix(strings.TrimSpace(description), ".")
}

func findBannedPhraseInDescription(description string) (string, bool) {
	for i, pattern := range bannedMarketingPhrasePatterns {
		if pattern.MatchString(description) {
			return bannedMarketingPhrases[i], true
		}
	}
	return "", false
}

// scopeKeywords lists positive agent-factory angle terms from CONTRIBUTING.md.
var scopeKeywords = []string{
	"coordination",
	"orchestration",
	"delegation",
	"routing",
	"handoff",
	"shared state",
	"group-level evaluation",
}

var scopeKeywordPatterns []*regexp.Regexp

func init() {
	initBannedMarketingPhrasePatterns()
	for _, phrase := range scopeKeywords {
		var pattern *regexp.Regexp
		switch phrase {
		case "handoff":
			pattern = regexp.MustCompile(`(?i)\bhandoffs?\b`)
		case "shared state":
			pattern = regexp.MustCompile(`(?i)shared\s+state`)
		case "group-level evaluation":
			pattern = regexp.MustCompile(`(?i)group[- ]level\s+evaluation`)
		default:
			escaped := regexp.QuoteMeta(phrase)
			pattern = regexp.MustCompile(`(?i)\b` + escaped + `\b`)
		}
		scopeKeywordPatterns = append(scopeKeywordPatterns, pattern)
	}
}

func descriptionHasScopeKeyword(description string) bool {
	for _, pattern := range scopeKeywordPatterns {
		if pattern.MatchString(description) {
			return true
		}
	}
	return false
}

// ReadmeDocument holds parsed README content for downstream validation rules.
type ReadmeDocument struct {
	Raw      string
	Sections []Section
}

// loadReadme reads README content from path. It returns an error when the file is absent.
func loadReadme(path string) (*ReadmeDocument, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("README.md is missing")
		}
		return nil, fmt.Errorf("read README.md: %w", err)
	}

	content := string(data)
	return &ReadmeDocument{
		Raw:      content,
		Sections: parseSections(content),
	}, nil
}

// readmePath resolves README.md from the repository root, walking up from the working directory.
func readmePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		readme := filepath.Join(dir, "README.md")
		if _, err := os.Stat(readme); err == nil {
			return readme, nil
		}
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return "", errors.New("README.md is missing")
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("README.md is missing")
		}
		dir = parent
	}
}

func sectionHeading(line string) (string, bool) {
	trimmed := strings.TrimSpace(line)
	if !strings.HasPrefix(trimmed, "## ") {
		return "", false
	}
	heading := strings.TrimSpace(strings.TrimPrefix(trimmed, "## "))
	if heading == "" || strings.HasPrefix(heading, "#") {
		return "", false
	}
	return heading, true
}
