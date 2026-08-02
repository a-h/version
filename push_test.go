package main

import "testing"

func TestBuildTag(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		version  string
		expected string
	}{
		{name: "v prefix applied to bare version", prefix: "v", version: "0.0.95", expected: "v0.0.95"},
		{name: "v prefix applied to v-prefixed version does not double", prefix: "v", version: "v0.0.95", expected: "v0.0.95"},
		{name: "empty prefix applied to v-prefixed version strips template prefix", prefix: "", version: "v0.0.95", expected: "0.0.95"},
		{name: "vv prefix applied to v-prefixed version does not produce vvv", prefix: "vv", version: "v0.0.95", expected: "vv0.0.95"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTag(tt.prefix, tt.version); got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestInferPrefix(t *testing.T) {
	tests := []struct {
		name            string
		tags            []string
		expectedPrefix  string
		expectedHasTags bool
	}{
		{
			name:            "no tags",
			tags:            []string{},
			expectedHasTags: false,
		},
		{
			name:            "all tags share v prefix",
			tags:            []string{"v0.0.91", "v0.0.89", "v0.0.87"},
			expectedPrefix:  "v",
			expectedHasTags: true,
		},
		{
			name:            "malformed tags are outvoted by correct tags",
			tags:            []string{"vvvv0.0.97", "vvv0.0.95", "v0.0.91", "v0.0.89", "v0.0.87", "v0.0.84"},
			expectedPrefix:  "v",
			expectedHasTags: true,
		},
		{
			name:            "bare tags are inferred",
			tags:            []string{"0.0.91", "0.0.89"},
			expectedPrefix:  "",
			expectedHasTags: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix, gotHasTags := inferPrefix(tt.tags)
			if gotHasTags != tt.expectedHasTags {
				t.Fatalf("hasTags: got %v, expected %v", gotHasTags, tt.expectedHasTags)
			}
			if gotPrefix != tt.expectedPrefix {
				t.Errorf("prefix: got %q, expected %q", gotPrefix, tt.expectedPrefix)
			}
		})
	}
}

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
		{name: "empty string returns empty prefix", input: "", expected: ""},
		{name: "string of only digits has no prefix", input: "123", expected: ""},
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
