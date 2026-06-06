package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func captureStderr(fn func()) string {
	old := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stderr = w

	done := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.String()
	}()

	fn()
	_ = w.Close()
	os.Stderr = old
	return <-done
}

func setupTestRepo(t *testing.T, readmeContent string) {
	t.Helper()

	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(orig)
	})

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module test\n\ngo 1.24\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte(readmeContent), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
}

func TestCollectCheckResultsRunsAllValidators(t *testing.T) {
	markdown := resourceSectionFixture("Frameworks", strings.TrimSpace(`
- [Beta](https://example.com/b) - Second framework for orchestration.
- [Alpha](https://example.com/a) - First framework for orchestration.
`))
	doc := &ReadmeDocument{
		Raw:      markdown,
		Sections: parseSections(markdown),
	}

	results := collectCheckResults(doc)
	rules := make(map[string]bool)
	for _, result := range results {
		rules[result.Rule] = true
	}

	if !rules["alphabetical-order"] {
		t.Fatal("expected alphabetical-order failure from integrated validators")
	}
}

func TestReportResultsHardFailures(t *testing.T) {
	results := []CheckResult{
		{
			Level:   "failure",
			Rule:    "entry-format",
			Message: `section "Frameworks" line 42: invalid resource entry format; expected - [Name](URL) - Description.`,
		},
		{
			Level:   "failure",
			Rule:    "description-period",
			Message: `section "Frameworks" line 43: resource "Example" description must end with a period (rule: description-period).`,
		},
	}

	stderr := captureStderr(func() {
		if err := reportResults(results); err == nil {
			t.Fatal("expected error for hard failures")
		} else if !strings.Contains(err.Error(), "2 check failure(s)") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if !strings.Contains(stderr, "entry-format") && !strings.Contains(stderr, "invalid resource entry format") {
		t.Fatalf("stderr missing entry-format failure: %q", stderr)
	}
	if !strings.Contains(stderr, "description-period") {
		t.Fatalf("stderr missing description-period failure: %q", stderr)
	}
	if strings.Contains(stderr, "README checks passed") {
		t.Fatalf("stderr should not report success: %q", stderr)
	}
}

func TestReportResultsWarningsOnly(t *testing.T) {
	results := []CheckResult{
		{
			Level:   "warning",
			Rule:    "scope-keyword",
			Message: `section "Frameworks" line 42: resource "Example" description lacks agent-factory scope keyword (rule: scope-keyword).`,
		},
	}

	stderr := captureStderr(func() {
		if err := reportResults(results); err != nil {
			t.Fatalf("warnings must not fail: %v", err)
		}
	})

	if !strings.Contains(stderr, "scope-keyword") {
		t.Fatalf("stderr missing scope warning: %q", stderr)
	}
	if !strings.Contains(stderr, "README checks passed with 1 warning(s)") {
		t.Fatalf("stderr missing warning summary: %q", stderr)
	}
}

func TestRunValidEmptySections(t *testing.T) {
	setupTestRepo(t, validEmptySectionFixture())

	stderr := captureStderr(func() {
		if err := run(); err != nil {
			t.Fatalf("valid empty sections should pass: %v", err)
		}
	})
	if !strings.Contains(stderr, "README checks passed") {
		t.Fatalf("expected success summary on stderr: %q", stderr)
	}
}

func TestRunHardFailureExitsNonZero(t *testing.T) {
	markdown := resourceSectionFixture("Frameworks", `- [Bad Entry](https://example.com) missing description separator`)
	setupTestRepo(t, markdown)

	stderr := captureStderr(func() {
		if err := run(); err == nil {
			t.Fatal("expected hard failure")
		} else if !strings.Contains(err.Error(), "check failure") {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if !strings.Contains(stderr, "invalid resource entry format") {
		t.Fatalf("stderr missing format failure: %q", stderr)
	}
}

func TestRunAgainstRepositoryReadme(t *testing.T) {
	if _, err := readmePath(); err != nil {
		t.Skip("repository README.md not available from test working directory")
	}

	stderr := captureStderr(func() {
		if err := run(); err != nil {
			t.Fatalf("repository README should pass all hard checks: %v", err)
		}
	})
	if !strings.Contains(stderr, "README checks passed") {
		t.Fatalf("expected success summary: %q", stderr)
	}
}
