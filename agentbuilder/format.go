package agentbuilder

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// The CLI's file format: JSON with a 4-space indent and no trailing newline (serde_json's pretty
// printer, and JavaScript's JSON.stringify(value, null, 4)).
const fileIndent = "    "

// formatIndex writes an index file (agents.json, tools.json), whose objects are *wirejson.Object
// values in the CLI's field order.
func formatIndex(value any) (string, error) {
	return wirejson.Indent(value, fileIndent)
}

// formatConfig writes a config file, with keys sorted at every level as the CLI writes them.
func formatConfig(value any) (string, error) {
	tree, err := wirejson.Canonical(value)
	if err != nil {
		return "", err
	}
	return wirejson.Indent(tree, fileIndent)
}

// isJSSpace reports whether r is whitespace to JavaScript's trim() and /\s/.
func isJSSpace(r rune) bool {
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0xa0, 0x1680, 0x2028, 0x2029, 0x202f, 0x205f, 0x3000, 0xfeff:
		return true
	}
	return r >= 0x2000 && r <= 0x200a
}

// sanitizeFilename is the CLI's sanitize_filename: a filesystem-safe file stem for an agent or
// tool name. / \ : * ? " < > | and whitespace become -, runs of - collapse, leading and trailing
// . and - are stripped, a name that started with . gets a _ prefix, an empty result becomes
// "unnamed", and names are cut at 100 characters.
func sanitizeFilename(name string) string {
	trimmed := strings.TrimFunc(name, isJSSpace)
	if trimmed == "" {
		return "unnamed"
	}
	sanitized := strings.Map(func(r rune) rune {
		if strings.ContainsRune(`/\:*?"<>|`, r) || isJSSpace(r) {
			return '-'
		}
		return r
	}, trimmed)
	for strings.Contains(sanitized, "--") {
		sanitized = strings.ReplaceAll(sanitized, "--", "-")
	}
	stripped := strings.Trim(sanitized, ".-")
	if stripped == "" {
		return "unnamed"
	}
	result := stripped
	if strings.HasPrefix(trimmed, ".") {
		result = "_" + stripped
	}
	if runes := []rune(result); len(runes) > 100 {
		result = string(runes[:100])
	}
	return result
}

// uniquePath is the CLI's generate_unique_filename: dir/<name>.json, then -1, -2... while the
// path is taken.
func uniquePath(dir, name string, taken map[string]bool) string {
	stem := sanitizeFilename(name)
	candidate := dir + "/" + stem + ".json"
	for n := 1; taken[candidate]; n++ {
		candidate = dir + "/" + stem + "-" + strconv.Itoa(n) + ".json"
	}
	return candidate
}

// normalizePath is a project-relative path in one canonical spelling: forward slashes, no ./ or
// empty segments.
func normalizePath(path string) string {
	var parts []string
	for _, part := range strings.Split(strings.ReplaceAll(path, `\`, "/"), "/") {
		if part != "" && part != "." {
			parts = append(parts, part)
		}
	}
	return strings.Join(parts, "/")
}

// decodeOrdered parses JSON text with each object as a *wirejson.Object holding its keys in the
// order JavaScript's JSON.parse gives them back, so an index file can be written back in the same
// order.
func decodeOrdered(data []byte) (any, error) {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil, errors.New("no JSON value")
	}
	switch data[0] {
	case '{':
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(data, &fields); err != nil {
			return nil, err
		}
		keys, err := wirejson.ObjectKeys(data)
		if err != nil {
			return nil, err
		}
		obj := &wirejson.Object{Keys: keys, Values: make(map[string]any, len(keys))}
		for _, key := range keys {
			value, err := decodeOrdered(fields[key])
			if err != nil {
				return nil, err
			}
			obj.Values[key] = value
		}
		return obj, nil
	case '[':
		var items []json.RawMessage
		if err := json.Unmarshal(data, &items); err != nil {
			return nil, err
		}
		out := make([]any, len(items))
		for i, item := range items {
			value, err := decodeOrdered(item)
			if err != nil {
				return nil, err
			}
			out[i] = value
		}
		return out, nil
	}
	return wirejson.Decode(data)
}

// get reads a key of a decoded object, or nil.
func get(obj *wirejson.Object, key string) any {
	if obj == nil {
		return nil
	}
	return obj.Values[key]
}

// set sets a key of an ordered object, appending it if it's new.
func set(obj *wirejson.Object, key string, value any) {
	if _, ok := obj.Values[key]; !ok {
		obj.Keys = append(obj.Keys, key)
	}
	obj.Values[key] = value
}

// orderedFields is an object with fields in the given order, leaving out empty strings and nil
// values.
func orderedFields(pairs ...any) *wirejson.Object {
	obj := &wirejson.Object{Values: map[string]any{}}
	for i := 0; i+1 < len(pairs); i += 2 {
		key := pairs[i].(string)
		switch v := pairs[i+1].(type) {
		case nil:
			continue
		case string:
			if v == "" {
				continue
			}
		case *wirejson.Object:
			if v == nil {
				continue
			}
		}
		set(obj, key, pairs[i+1])
	}
	return obj
}
