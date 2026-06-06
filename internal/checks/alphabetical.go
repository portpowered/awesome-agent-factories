package main

// CheckResult records a hard failure or non-blocking warning from a validation rule.
type CheckResult struct {
	Level   string // "failure" or "warning"
	Rule    string
	Message string
}
