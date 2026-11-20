package agenttools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// object is a loaded JSON object with its keys in the order JavaScript's JSON.parse gives them back
// (array-index keys first, then document order). Loading reads property, path param and header
// order from it, since that order decides a schema's `required` list. Everything the builder keeps
// is converted to a plain wire tree (map[string]any) with plain.
type object struct {
	keys []string
	vals map[string]any
}

func (o *object) get(key string) (any, bool) {
	value, ok := o.vals[key]
	return value, ok
}

// toTree reads data (JSON text as a string, []byte or json.RawMessage, or any value that marshals
// to JSON, such as a map or a generated model) as a loaded tree. A parsed Go map has no order and
// is read in sorted key order.
func toTree(data any) (any, error) {
	var raw []byte
	switch v := data.(type) {
	case string:
		raw = []byte(v)
	case []byte:
		raw = v
	case json.RawMessage:
		raw = v
	default:
		encoded, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("agenttools: can't convert %T to JSON: %w", data, err)
		}
		raw = encoded
	}
	tree, err := decodeTree(raw)
	if err != nil {
		return nil, fmt.Errorf("agenttools: invalid JSON: %w", err)
	}
	return tree, nil
}

// decodeTree parses JSON text into a loaded tree: objects are *object, numbers keep their
// json.Number text. A key that appears twice keeps its first position and its last value, as in
// JavaScript.
func decodeTree(data []byte) (any, error) {
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
		obj := &object{keys: keys, vals: make(map[string]any, len(keys))}
		for _, key := range keys {
			value, err := decodeTree(fields[key])
			if err != nil {
				return nil, err
			}
			obj.vals[key] = value
		}
		return obj, nil
	case '[':
		var items []json.RawMessage
		if err := json.Unmarshal(data, &items); err != nil {
			return nil, err
		}
		out := make([]any, len(items))
		for i, item := range items {
			value, err := decodeTree(item)
			if err != nil {
				return nil, err
			}
			out[i] = value
		}
		return out, nil
	}
	return wirejson.Decode(data)
}

// plain converts a loaded tree to a plain wire tree (map[string]any for objects).
func plain(value any) any {
	switch v := value.(type) {
	case *object:
		out := make(map[string]any, len(v.keys))
		for _, key := range v.keys {
			out[key] = plain(v.vals[key])
		}
		return out
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = plain(item)
		}
		return out
	}
	return value
}

// stripNulls drops null object entries at every level, since the API sends null for unset fields
// and the other SDKs read it as absent. Free-form JSON values (constants and dynamic variable
// placeholders) are kept as they are, and so is a null constant_value, which loadSchema keeps for a
// nullable literal. The values of name maps (properties, path_params_schema, request_headers) are
// walked as values, so a property named "constant_value" is no exception.
func stripNulls(value any) any {
	switch v := value.(type) {
	case *object:
		out := &object{vals: make(map[string]any, len(v.keys))}
		for _, key := range v.keys {
			item := v.vals[key]
			if item == nil && key != "constant_value" {
				continue
			}
			switch key {
			case "constant_value", "dynamic_variable_placeholders":
			case "properties", "path_params_schema", "request_headers":
				if names, ok := item.(*object); ok {
					item = stripNames(names)
				} else {
					item = stripNulls(item)
				}
			default:
				item = stripNulls(item)
			}
			out.keys = append(out.keys, key)
			out.vals[key] = item
		}
		return out
	case []any:
		out := make([]any, len(v))
		for i, item := range v {
			out[i] = stripNulls(item)
		}
		return out
	}
	return value
}

func stripNames(names *object) *object {
	out := &object{vals: make(map[string]any, len(names.keys))}
	for _, name := range names.keys {
		item := names.vals[name]
		if item == nil {
			continue
		}
		out.keys = append(out.keys, name)
		out.vals[name] = stripNulls(item)
	}
	return out
}

func cloneMap(m map[string]any) map[string]any {
	if m == nil {
		return map[string]any{}
	}
	return wirejson.Clone(m).(map[string]any)
}

// stringList reads a JSON array of strings, reporting false if value is another shape.
func stringList(value any) ([]string, bool) {
	items, ok := value.([]any)
	if !ok {
		return nil, false
	}
	out := make([]string, len(items))
	for i, item := range items {
		s, ok := item.(string)
		if !ok {
			return nil, false
		}
		out[i] = s
	}
	return out, true
}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
