package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConvAiFileUploadResponseModel struct {
	FileId          string `json:"file_id" xml:"file_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConvAiFileUploadResponseModel) GetFileId() (value string) {
	if c == nil {
		return
	}
	return c.FileId
}

func (c *ConvAiFileUploadResponseModel) SetFileId(fileId string) {
	c.FileId = fileId
}

func (c *ConvAiFileUploadResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiFileUploadResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiFileUploadResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConvAiFileUploadResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConvAiFileUploadResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConvAiFileUploadResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiFileUploadResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConvAiFileUploadResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
