package main

import "testing"

func TestLeadingPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "v prefix is detected", input: "v0.0.2", expected: "v"},
		{name: "bare version has no prefix", input: "0.0.2", expected: ""},
		{name: "multi-character prefix is detected", input: "release-1.2.3", expected: "release-"},
		{name: "tag without a digit returns the whole tag", input: "latest", expected: "latest"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leadingPrefix(tt.input); got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestResolvePrefix(t *testing.T) {
	ptr := func(s string) *string { return &s }
	tests := []struct {
		name           string
		flag           *string
		force          bool
		existingPrefix string
		hasTags        bool
		expected       string
		expectError    bool
	}{
		{name: "no tags defaults to v", flag: nil, hasTags: false, expected: "v"},
		{name: "existing v tags are inferred", flag: nil, existingPrefix: "v", hasTags: true, expected: "v"},
		{name: "existing bare tags are inferred", flag: nil, existingPrefix: "", hasTags: true, expected: ""},
		{name: "explicit flag matching existing tags is accepted", flag: ptr("v"), existingPrefix: "v", hasTags: true, expected: "v"},
		{name: "explicit flag conflicting with existing tags errors", flag: ptr(""), existingPrefix: "v", hasTags: true, expectError: true},
		{name: "force overrides a conflicting flag", flag: ptr(""), force: true, existingPrefix: "v", hasTags: true, expected: ""},
		{name: "explicit flag is used when there are no tags", flag: ptr("v"), hasTags: false, expected: "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolvePrefix(tt.flag, tt.force, tt.existingPrefix, tt.hasTags)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}
