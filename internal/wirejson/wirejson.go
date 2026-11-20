// Package wirejson is the canonical JSON layer shared by the hand-written agents-as-code builders
// (the workflow builder in agentworkflows and the tools builder in agenttools).
//
// The builders keep their state in API wire format as plain Go JSON trees: nil, bool, string,
// json.Number, []any and map[string]any. This package converts Go values into such trees, reads
// object keys in the order JavaScript sees them, and writes the canonical file format that every
// ElevenLabs SDK emits byte for byte: keys sorted at every level, 2-space indentation, a trailing
// newline, and strings and numbers written as JavaScript's JSON.stringify writes them.
//
// The package is hand-written (listed in .fernignore) and internal to the module, so the builders
// can share it without making it part of the public API.
package wirejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"unicode/utf16"
)

// Decode parses JSON text into a wire tree, keeping every number as its json.Number text so no
// value is rounded through a float64.
func Decode(data []byte) (any, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var value any
	if err := dec.Decode(&value); err != nil {
		return nil, err
	}
	if _, err := dec.Token(); err != io.EOF {
		return nil, errors.New("wirejson: unexpected data after the JSON value")
	}
	return value, nil
}

// Normalize converts any JSON-marshalable Go value (generated SDK types, structs, maps, slices) to
// a wire tree. Values that are already wire trees come back as deep copies.
func Normalize(value any) (any, error) {
	if tree, ok := asTree(value); ok {
		return tree, nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return Decode(data)
}

// asTree deep-copies value when it is already made only of wire-tree types.
func asTree(value any) (any, bool) {
	switch v := value.(type) {
	case nil, bool, string, json.Number:
		return v, true
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			copied, ok := asTree(item)
			if !ok {
				return nil, false
			}
			out[i] = copied
		}
		return out, true
	case map[string]any:
		out := make(map[string]any, len(v))
		for key, item := range v {
			copied, ok := asTree(item)
			if !ok {
				return nil, false
			}
			out[key] = copied
		}
		return out, true
	}
	return nil, false
}

// Clone deep-copies a wire tree. Values of other types are returned as they are.
func Clone(value any) any {
	switch v := value.(type) {
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = Clone(item)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for key, item := range v {
			out[key] = Clone(item)
		}
		return out
	}
	return value
}

// ObjectKeys returns the keys of the JSON object data in the order JavaScript's JSON.parse gives
// them back (Object.keys): keys that are array indices ("0", "1", ... up to 2^32-2) first, in
// ascending numeric order, then every other key in document order. A key that appears twice keeps
// its first position. Builders use it where document order decides the output, such as a
// workflow's node and edge maps.
func ObjectKeys(data []byte) ([]string, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if delim, ok := tok.(json.Delim); !ok || delim != '{' {
		return nil, errors.New("wirejson: expected a JSON object")
	}
	seen := map[string]bool{}
	var keys []string
	for dec.More() {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		key, ok := tok.(string)
		if !ok {
			return nil, errors.New("wirejson: expected an object key")
		}
		var skip json.RawMessage
		if err := dec.Decode(&skip); err != nil {
			return nil, err
		}
		if !seen[key] {
			seen[key] = true
			keys = append(keys, key)
		}
	}
	return jsPropertyOrder(keys, false), nil
}

// SortKeys sorts object keys the way the reference JavaScript SDK's output orders them: the keys
// are sorted by UTF-16 code units (Array.prototype.sort), and the object built from them then
// lists array-index keys first, in ascending numeric order.
func SortKeys(keys []string) []string {
	return jsPropertyOrder(keys, true)
}

func jsPropertyOrder(keys []string, sortRest bool) []string {
	var indices []string
	var rest []string
	for _, key := range keys {
		if isArrayIndex(key) {
			indices = append(indices, key)
		} else {
			rest = append(rest, key)
		}
	}
	sort.SliceStable(indices, func(i, j int) bool {
		a, _ := strconv.ParseUint(indices[i], 10, 32)
		b, _ := strconv.ParseUint(indices[j], 10, 32)
		return a < b
	})
	if sortRest {
		sort.SliceStable(rest, func(i, j int) bool { return lessUTF16(rest[i], rest[j]) })
	}
	return append(indices, rest...)
}

// isArrayIndex reports whether key is an ECMAScript array index: the canonical decimal form of an
// integer from 0 to 2^32-2.
func isArrayIndex(key string) bool {
	if key == "" || len(key) > 10 || (len(key) > 1 && key[0] == '0') {
		return false
	}
	for i := 0; i < len(key); i++ {
		if key[i] < '0' || key[i] > '9' {
			return false
		}
	}
	n, err := strconv.ParseUint(key, 10, 64)
	return err == nil && n < 1<<32-1
}

func lessUTF16(a, b string) bool {
	ua := utf16.Encode([]rune(a))
	ub := utf16.Encode([]rune(b))
	for i := 0; i < len(ua) && i < len(ub); i++ {
		if ua[i] != ub[i] {
			return ua[i] < ub[i]
		}
	}
	return len(ua) < len(ub)
}

// Canonical returns a copy of value (a wire tree, or anything Normalize accepts) with every number
// rewritten as JavaScript writes it. Object keys are unchanged; Go maps have no order, and String
// sorts them when writing.
func Canonical(value any) (any, error) {
	tree, err := Normalize(value)
	if err != nil {
		return nil, err
	}
	return canonicalNumbers(tree), nil
}

func canonicalNumbers(value any) any {
	switch v := value.(type) {
	case json.Number:
		return json.Number(FormatNumber(string(v)))
	case []any:
		for i, item := range v {
			v[i] = canonicalNumbers(item)
		}
		return v
	case map[string]any:
		for key, item := range v {
			v[key] = canonicalNumbers(item)
		}
		return v
	}
	return value
}

// String writes value in the canonical file format: object keys sorted (see SortKeys), 2-space
// indentation, strings and numbers as JSON.stringify writes them, and a trailing newline. It is the
// same text as JavaScript's JSON.stringify(sortKeys(value), null, 2) + "\n".
func String(value any) (string, error) {
	tree, err := Normalize(value)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := write(&buf, tree, 0, "  "); err != nil {
		return "", err
	}
	buf.WriteByte('\n')
	return buf.String(), nil
}

// Object is a JSON object that Indent writes with its keys in Keys order instead of sorted, for
// files whose field order is fixed (such as the CLI's agents.json). A key in Keys that Values
// doesn't hold is left out.
type Object struct {
	Keys   []string
	Values map[string]any
}

// Indent writes value as JavaScript's JSON.stringify(value, null, indent) writes it, with no
// trailing newline: map[string]any objects with their keys sorted (see SortKeys), *Object values in
// their own order, and strings and numbers as String writes them. value must be a wire tree, which
// may hold *Object values.
func Indent(value any, indent string) (string, error) {
	var buf bytes.Buffer
	if err := write(&buf, value, 0, indent); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func write(buf *bytes.Buffer, value any, depth int, indent string) error {
	switch v := value.(type) {
	case nil:
		buf.WriteString("null")
	case bool:
		if v {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
	case string:
		writeString(buf, v)
	case json.Number:
		buf.WriteString(FormatNumber(string(v)))
	case []any:
		if len(v) == 0 {
			buf.WriteString("[]")
			return nil
		}
		buf.WriteByte('[')
		for i, item := range v {
			if i > 0 {
				buf.WriteByte(',')
			}
			newline(buf, depth+1, indent)
			if err := write(buf, item, depth+1, indent); err != nil {
				return err
			}
		}
		newline(buf, depth, indent)
		buf.WriteByte(']')
	case map[string]any:
		if len(v) == 0 {
			buf.WriteString("{}")
			return nil
		}
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		return writeObject(buf, SortKeys(keys), v, depth, indent)
	case *Object:
		var keys []string
		for _, key := range v.Keys {
			if _, ok := v.Values[key]; ok {
				keys = append(keys, key)
			}
		}
		if len(keys) == 0 {
			buf.WriteString("{}")
			return nil
		}
		return writeObject(buf, keys, v.Values, depth, indent)
	default:
		return fmt.Errorf("wirejson: %T is not a wire value", value)
	}
	return nil
}

func writeObject(buf *bytes.Buffer, keys []string, values map[string]any, depth int, indent string) error {
	buf.WriteByte('{')
	for i, key := range keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		newline(buf, depth+1, indent)
		writeString(buf, key)
		buf.WriteString(": ")
		if err := write(buf, values[key], depth+1, indent); err != nil {
			return err
		}
	}
	newline(buf, depth, indent)
	buf.WriteByte('}')
	return nil
}

func newline(buf *bytes.Buffer, depth int, indent string) {
	buf.WriteByte('\n')
	for i := 0; i < depth; i++ {
		buf.WriteString(indent)
	}
}

// writeString quotes s as JSON.stringify does: only `"`, `\` and control characters are escaped,
// with the short escapes where JSON has them and lowercase \u00xx otherwise. Everything else,
// including "<", "/", "+" and U+2028, is written as is.
func writeString(buf *bytes.Buffer, s string) {
	const hex = "0123456789abcdef"
	buf.WriteByte('"')
	for _, r := range s {
		switch r {
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		case '\b':
			buf.WriteString(`\b`)
		case '\f':
			buf.WriteString(`\f`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		default:
			if r < 0x20 {
				buf.WriteString(`\u00`)
				buf.WriteByte(hex[r>>4])
				buf.WriteByte(hex[r&0xf])
			} else {
				buf.WriteRune(r)
			}
		}
	}
	buf.WriteByte('"')
}
