package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FileIdSource struct {
	type_ string
	// The ID that `POST /v1/convai/conversations/{conversation_id}/files` returned for this conversation.
	FileId          string `json:"file_id" xml:"file_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (f *FileIdSource) GetFileId() (value string) {
	if f == nil {
		return
	}
	return f.FileId
}

func (f *FileIdSource) SetFileId(fileId string) {
	f.FileId = fileId
}

func (f *FileIdSource) Type() string {
	return f.type_
}

func (f *FileIdSource) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FileIdSource) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FileIdSource to string"
	}
	return string(jsonData)
}

func (f *FileIdSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FileIdSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FileIdSource(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "file_id" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", f, "file_id", *literals.Type)
	}
	f.type_ = "file_id"
	extraProperties, err := internal.ExtractExtraProperties(data, *f, "type")
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FileIdSource) MarshalJSON() ([]byte, error) {
	type embed FileIdSource
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*f),
		Type:  "file_id",
	}
	return json.Marshal(marshaler)
}
