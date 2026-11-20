package wirejson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFormatNumberMatchesJavaScript(t *testing.T) {
	cases := map[string]string{
		"0":                       "0",
		"-0":                      "0",
		"0.0":                     "0",
		"300":                     "300",
		"300.0":                   "300",
		"3e2":                     "300",
		"-12.50":                  "-12.5",
		"0.1":                     "0.1",
		"1E-07":                   "1e-7",
		"1e-7":                    "1e-7",
		"0.00000015":              "1.5e-7",
		"0.000001":                "0.000001",
		"1e-6":                    "0.000001",
		"1e21":                    "1e+21",
		"1.5e300":                 "1.5e+300",
		"123456789012345680000":   "123456789012345680000",
		"1.2345678901234568e20":   "123456789012345680000",
		"1e20":                    "100000000000000000000",
		"9007199254740992":        "9007199254740992",
		"9007199254740993":        "9007199254740993",
		"9007199254740992.5":      "9007199254740992.5",
		"12345678901234567890123": "12345678901234567890123",
		"1e400":                   "1e400",
		"0.30000000000000004":     "0.30000000000000004",
	}
	for in, want := range cases {
		if got := FormatNumber(in); got != want {
			t.Errorf("FormatNumber(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestIsIntegerReadsEveryDigit(t *testing.T) {
	cases := map[string]bool{
		"0":                     true,
		"-0":                    true,
		"3":                     true,
		"3.0":                   true,
		"1E2":                   true,
		"1.5e1":                 true,
		"9223372036854775808":   true,
		"1.5":                   false,
		"1e-1":                  false,
		"1.0000000000000000001": false,
		"12.5e-1":               false,
		"x":                     false,
		"":                      false,
	}
	for in, want := range cases {
		if got := IsInteger(in); got != want {
			t.Errorf("IsInteger(%q) = %v, want %v", in, got, want)
		}
	}
}

func TestStringEscapesLikeJSONStringify(t *testing.T) {
	got, err := String(map[string]any{"s": "a\"b\\c\n\t\u0001<>&/+ é"})
	if err != nil {
		t.Fatal(err)
	}
	want := "{\n  \"s\": \"a\\\"b\\\\c\\n\\t\\u0001<>&/+ é\"\n}\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStringLayout(t *testing.T) {
	got, err := String(map[string]any{
		"b":     []any{json.Number("1"), map[string]any{}},
		"a":     []any{},
		"10":    true,
		"9":     nil,
		"float": 2.5,
	})
	if err != nil {
		t.Fatal(err)
	}
	want := `{
  "9": null,
  "10": true,
  "a": [],
  "b": [
    1,
    {}
  ],
  "float": 2.5
}
`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestObjectKeysUseJavaScriptOrder(t *testing.T) {
	keys, err := ObjectKeys([]byte(`{"z": 1, "b": {"x": 1}, "2": 0, "a": [], "1": 0, "01": 0, "z": 2}`))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"1", "2", "z", "b", "a", "01"}
	if !reflect.DeepEqual(keys, want) {
		t.Errorf("got %v, want %v", keys, want)
	}
}

func TestSortKeysUsesUTF16Order(t *testing.T) {
	// U+FF5E sorts after U+1F600 in UTF-8 byte order, but before it in UTF-16 code units.
	keys := SortKeys([]string{"\U0001F600", "～", "a"})
	want := []string{"a", "\U0001F600", "～"}
	if !reflect.DeepEqual(keys, want) {
		t.Errorf("got %q, want %q", keys, want)
	}
}

func TestDecodeKeepsNumberText(t *testing.T) {
	value, err := Decode([]byte(`{"n": 9007199254740993}`))
	if err != nil {
		t.Fatal(err)
	}
	if got := value.(map[string]any)["n"]; got != json.Number("9007199254740993") {
		t.Errorf("got %v", got)
	}
	if _, err := Decode([]byte(`{} {}`)); err == nil {
		t.Error("trailing data must be rejected")
	}
}

func TestIndentKeepsObjectOrderAndSortsMaps(t *testing.T) {
	value := &Object{
		Keys: []string{"z", "missing", "a"},
		Values: map[string]any{
			"z": map[string]any{"b": json.Number("1.0"), "a": []any{}},
			"a": []any{&Object{}, "x"},
		},
	}
	got, err := Indent(value, "    ")
	if err != nil {
		t.Fatal(err)
	}
	want := "{\n    \"z\": {\n        \"a\": [],\n        \"b\": 1\n    },\n    \"a\": [\n        {},\n        \"x\"\n    ]\n}"
	if got != want {
		t.Fatalf("Indent = %q", got)
	}
}
