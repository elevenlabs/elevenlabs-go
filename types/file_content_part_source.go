package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FileContentPartSource struct {
	Type   string
	Base64 *Base64FileSource
	FileId *FileIdSource
}

func (f *FileContentPartSource) GetBase64() *Base64FileSource {
	if f == nil {
		return nil
	}
	if f.Base64 == nil {
		return nil
	}
	return f.Base64
}

func (f *FileContentPartSource) GetFileId() *FileIdSource {
	if f == nil {
		return nil
	}
	if f.FileId == nil {
		return nil
	}
	return f.FileId
}

func (f *FileContentPartSource) GetType() string {
	if f == nil {
		return ""
	}
	return f.Type
}

func (f *FileContentPartSource) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	f.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", f)
	}
	switch unmarshaler.Type {
	case "base64":
		value := new(Base64FileSource)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		f.Base64 = value
	case "file_id":
		value := new(FileIdSource)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		f.FileId = value
	}
	return nil
}

func (f FileContentPartSource) MarshalJSON() ([]byte, error) {
	if err := f.validate(); err != nil {
		return nil, err
	}
	if f.Base64 != nil {
		return internal.MarshalJSONWithDiscriminant(f.Base64, "type", "base64")
	}
	if f.FileId != nil {
		return internal.MarshalJSONWithDiscriminant(f.FileId, "type", "file_id")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", f)
}

type FileContentPartSourceVisitor interface {
	VisitBase64(*Base64FileSource) error
	VisitFileId(*FileIdSource) error
}

func (f *FileContentPartSource) Accept(visitor FileContentPartSourceVisitor) error {
	if f.Base64 != nil {
		return visitor.VisitBase64(f.Base64)
	}
	if f.FileId != nil {
		return visitor.VisitFileId(f.FileId)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", f)
}

func (f *FileContentPartSource) validate() error {
	if f == nil {
		return fmt.Errorf("type %T is nil", f)
	}
	var fields []string
	if f.Base64 != nil {
		fields = append(fields, "base64")
	}
	if f.FileId != nil {
		fields = append(fields, "file_id")
	}
	if len(fields) == 0 {
		if f.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", f, f.Type)
		}
		return fmt.Errorf("type %T is empty", f)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", f, fields)
	}
	if f.Type != "" {
		field := fields[0]
		if f.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				f,
				f.Type,
				f,
			)
		}
	}
	return nil
}
