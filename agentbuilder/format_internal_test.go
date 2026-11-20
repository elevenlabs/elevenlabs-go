package agentbuilder

import (
	"strings"
	"testing"
)

func TestFileNamesFollowTheCLI(t *testing.T) {
	for name, want := range map[string]string{
		"  Support line / EU ":   "Support-line-EU",
		".hidden":                "_hidden",
		"...":                    "unnamed",
		"":                       "unnamed",
		"a:b*c?d\"e<f>g|h\\i":    "a-b-c-d-e-f-g-h-i",
		"--x--":                  "x",
		"tab\tand nbsp":          "tab-and-nbsp",
		strings.Repeat("é", 120): strings.Repeat("é", 100),
	} {
		if got := sanitizeFilename(name); got != want {
			t.Errorf("sanitizeFilename(%q) = %q, want %q", name, got, want)
		}
	}
	taken := map[string]bool{"agent_configs/A.json": true, "agent_configs/A-1.json": true}
	if got := uniquePath("agent_configs", "A", taken); got != "agent_configs/A-2.json" {
		t.Fatalf("uniquePath = %q", got)
	}
	if got := normalizePath(`./agent_configs\\a.json`); got != "agent_configs/a.json" {
		t.Fatalf("normalizePath = %q", got)
	}
}

func TestFileFormat(t *testing.T) {
	text, err := formatConfig(map[string]any{"b": []any{}, "a": map[string]any{"y": 1, "x": "<&>"}})
	if err != nil {
		t.Fatal(err)
	}
	if want := "{\n    \"a\": {\n        \"x\": \"<&>\",\n        \"y\": 1\n    },\n    \"b\": []\n}"; text != want {
		t.Fatalf("formatConfig = %q", text)
	}
	index, err := formatIndex(orderedFields("type", "client", "config", "c.json", "id", ""))
	if err != nil {
		t.Fatal(err)
	}
	if want := "{\n    \"type\": \"client\",\n    \"config\": \"c.json\"\n}"; index != want {
		t.Fatalf("formatIndex = %q", index)
	}
}
