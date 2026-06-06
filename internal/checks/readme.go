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

func init() {
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
