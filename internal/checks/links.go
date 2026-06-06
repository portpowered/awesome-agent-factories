package main

// ResourceEntry is a parsed - [Name](URL) - Description. list item in a resource section.
type ResourceEntry struct {
	Name        string
	URL         string
	Description string
	Section     string
	Line        int
}
