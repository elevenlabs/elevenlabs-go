package internal

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// WebSocketFrame is a received text frame, parsed once so several message types can be tried.
type WebSocketFrame struct {
	Raw    json.RawMessage
	fields map[string]json.RawMessage
}

// ParseWebSocketFrame parses a text frame as JSON.
func ParseWebSocketFrame(data []byte) (*WebSocketFrame, error) {
	if !json.Valid(data) {
		return nil, &WebSocketFrameError{Data: string(data)}
	}
	frame := &WebSocketFrame{Raw: json.RawMessage(bytes.Clone(data))}
	_ = json.Unmarshal(data, &frame.fields)
	return frame, nil
}

// HasProperty reports whether the frame is an object carrying the named property.
func (f *WebSocketFrame) HasProperty(name string) bool {
	_, ok := f.fields[name]
	return ok
}

// HasLiteral reports whether the frame's named property equals the given JSON literal.
func (f *WebSocketFrame) HasLiteral(name string, literal string) bool {
	field, ok := f.fields[name]
	if !ok {
		return false
	}
	var actual, expected interface{}
	if json.Unmarshal(field, &actual) != nil || json.Unmarshal([]byte(literal), &expected) != nil {
		return false
	}
	return reflect.DeepEqual(actual, expected)
}

// Decode unmarshals the frame into the given value.
func (f *WebSocketFrame) Decode(value interface{}) bool {
	return json.Unmarshal(f.Raw, value) == nil
}

// MarshalWebSocketMessage marshals a message and adds each literal property the message did not
// write itself or left empty, such as a discriminator a model keeps unexported or a single-value
// enum field the caller did not set.
func MarshalWebSocketMessage(message interface{}, literals map[string]interface{}) ([]byte, error) {
	data, err := json.Marshal(message)
	if err != nil || len(literals) == 0 {
		return data, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil || fields == nil {
		return data, nil
	}
	for name, value := range literals {
		if written, ok := fields[name]; ok && !isEmptyJSONValue(written) {
			continue
		}
		encoded, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[name] = encoded
	}
	return json.Marshal(fields)
}

func isEmptyJSONValue(value json.RawMessage) bool {
	text := string(value)
	return text == `""` || text == "null"
}

// WebSocketFrameError reports a text frame that is not valid JSON.
type WebSocketFrameError struct {
	Data string
}

func (e *WebSocketFrameError) Error() string {
	return "websocket frame is not valid JSON"
}
