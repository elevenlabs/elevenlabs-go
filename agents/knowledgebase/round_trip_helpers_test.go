package knowledgebase

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func assertRoundTrip(t *testing.T, wire string, target any) {
	t.Helper()
	if err := json.Unmarshal([]byte(wire), target); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	encoded, err := json.Marshal(target)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var want, got map[string]any
	if err := json.Unmarshal([]byte(wire), &want); err != nil {
		t.Fatalf("unmarshal want: %v", err)
	}
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatalf("unmarshal got: %v", err)
	}
	exposed := exportedJSONKeys(reflect.TypeOf(target))
	for key, wantValue := range want {
		if !exposed[key] {
			continue
		}
		gotValue, ok := got[key]
		if !ok {
			t.Fatalf("key %q was dropped on the round trip\n got: %s\nwant: %s", key, encoded, wire)
		}
		if !reflect.DeepEqual(gotValue, wantValue) {
			t.Fatalf("key %q changed on the round trip\n got: %#v\nwant: %#v", key, gotValue, wantValue)
		}
	}
	for key := range got {
		if _, ok := want[key]; !ok {
			t.Fatalf("encoding added key %q the wire value did not have\n got: %s", key, encoded)
		}
	}
}

// exportedJSONKeys is the set of JSON keys a struct type exposes through exported fields.
func exportedJSONKeys(t reflect.Type) map[string]bool {
	keys := map[string]bool{}
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return keys
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" || field.Anonymous {
			continue
		}
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := strings.SplitN(tag, ",", 2)[0]
		if name == "" {
			name = field.Name
		}
		keys[name] = true
	}
	return keys
}
