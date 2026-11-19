package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"reflect"
	"sort"
	"strings"
)

const defaultMultipartContentType = "application/octet-stream"

var readerType = reflect.TypeOf((*io.Reader)(nil)).Elem()

// MultipartWriter writes a multipart/form-data request body.
type MultipartWriter struct {
	buffer *bytes.Buffer
	writer *multipart.Writer
}

// NewMultipartWriter returns a writer backed by an in-memory buffer.
func NewMultipartWriter() *MultipartWriter {
	buffer := bytes.NewBuffer(nil)
	return &MultipartWriter{buffer: buffer, writer: multipart.NewWriter(buffer)}
}

// Buffer returns the encoded body.
func (w *MultipartWriter) Buffer() *bytes.Buffer {
	return w.buffer
}

// ContentType returns the Content-Type header value, including the boundary.
func (w *MultipartWriter) ContentType() string {
	return w.writer.FormDataContentType()
}

// Close writes the trailing boundary.
func (w *MultipartWriter) Close() error {
	return w.writer.Close()
}

// WriteStruct writes every file field of value as a file part and every other
// JSON property as a form field. Nil values write nothing.
func (w *MultipartWriter) WriteStruct(value interface{}) error {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return nil
	}
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("multipart: expected a struct, got %s", v.Kind())
	}
	withoutFiles := reflect.New(v.Type()).Elem()
	withoutFiles.Set(v)
	files := map[string]bool{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		name := multipartFieldName(field)
		if name == "" || !isFileType(field.Type) {
			continue
		}
		files[name] = true
		if err := w.writeFiles(name, v.Field(i)); err != nil {
			return err
		}
		withoutFiles.Field(i).Set(reflect.Zero(field.Type))
	}
	encoded, err := json.Marshal(withoutFiles.Addr().Interface())
	if err != nil {
		return err
	}
	var properties map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &properties); err != nil {
		return err
	}
	keys := make([]string, 0, len(properties))
	for key := range properties {
		if !files[key] {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	for _, key := range keys {
		if err := w.writeProperty(key, properties[key]); err != nil {
			return err
		}
	}
	return nil
}

func multipartFieldName(field reflect.StructField) string {
	if !field.IsExported() {
		return ""
	}
	name := strings.Split(field.Tag.Get("json"), ",")[0]
	if name == "-" {
		return ""
	}
	return name
}

func isFileType(t reflect.Type) bool {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 {
		return true
	}
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
	}
	return t == readerType
}

func (w *MultipartWriter) writeFiles(name string, v reflect.Value) error {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Uint8 {
		return w.WriteFile(name, bytes.NewReader(v.Bytes()))
	}
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if err := w.writeFiles(name, v.Index(i)); err != nil {
				return err
			}
		}
		return nil
	}
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		if inner := v.Elem(); inner.Kind() == reflect.Pointer && inner.IsNil() {
			return nil
		}
	}
	reader, ok := v.Interface().(io.Reader)
	if !ok || reader == nil {
		return nil
	}
	return w.WriteFile(name, reader)
}

// WriteFile writes reader as a file part, using its Filename/Name and ContentType methods when present.
func (w *MultipartWriter) WriteFile(name string, reader io.Reader) error {
	filename := name
	if named, ok := reader.(interface{ Filename() string }); ok && named.Filename() != "" {
		filename = named.Filename()
	} else if named, ok := reader.(interface{ Name() string }); ok && named.Name() != "" {
		filename = named.Name()
	}
	contentType := defaultMultipartContentType
	if typed, ok := reader.(interface{ ContentType() string }); ok && typed.ContentType() != "" {
		contentType = typed.ContentType()
	}
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name=%q; filename=%q`, name, filename))
	header.Set("Content-Type", contentType)
	part, err := w.writer.CreatePart(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, reader)
	return err
}

func (w *MultipartWriter) writeProperty(key string, raw json.RawMessage) error {
	var decoded interface{}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return err
	}
	switch value := decoded.(type) {
	case nil:
		return nil
	case string:
		return w.writer.WriteField(key, value)
	case bool, float64:
		return w.writer.WriteField(key, string(raw))
	case []interface{}:
		if !allScalars(value) {
			return w.writer.WriteField(key, string(raw))
		}
		for _, item := range value {
			text, err := scalarText(item)
			if err != nil {
				return err
			}
			if err := w.writer.WriteField(key, text); err != nil {
				return err
			}
		}
		return nil
	default:
		return w.writer.WriteField(key, string(raw))
	}
}

func allScalars(items []interface{}) bool {
	for _, item := range items {
		switch item.(type) {
		case string, bool, float64:
		default:
			return false
		}
	}
	return true
}

func scalarText(item interface{}) (string, error) {
	if text, ok := item.(string); ok {
		return text, nil
	}
	encoded, err := json.Marshal(item)
	return string(encoded), err
}
