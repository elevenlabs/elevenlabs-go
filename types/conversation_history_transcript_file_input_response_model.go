package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptFileInputResponseModel struct {
	FileId           string `json:"file_id" xml:"file_id" required:"true"`
	OriginalFilename string `json:"original_filename" xml:"original_filename" required:"true"`
	MimeType         string `json:"mime_type" xml:"mime_type" required:"true"`
	FileUrl          string `json:"file_url" xml:"file_url" required:"true"`
	extraProperties  map[string]interface{}
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) GetFileId() (value string) {
	if c == nil {
		return
	}
	return c.FileId
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) GetOriginalFilename() (value string) {
	if c == nil {
		return
	}
	return c.OriginalFilename
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) GetMimeType() (value string) {
	if c == nil {
		return
	}
	return c.MimeType
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) GetFileUrl() (value string) {
	if c == nil {
		return
	}
	return c.FileUrl
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) SetFileId(fileId string) {
	c.FileId = fileId
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) SetOriginalFilename(originalFilename string) {
	c.OriginalFilename = originalFilename
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) SetMimeType(mimeType string) {
	c.MimeType = mimeType
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) SetFileUrl(fileUrl string) {
	c.FileUrl = fileUrl
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptFileInputResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptFileInputResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationHistoryTranscriptFileInputResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationHistoryTranscriptFileInputResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptFileInputResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptFileInputResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
