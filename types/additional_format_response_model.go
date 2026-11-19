package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AdditionalFormatResponseModel struct {
	// The requested format.
	RequestedFormat string `json:"requested_format" xml:"requested_format" required:"true"`
	// The file extension of the additional format.
	FileExtension string `json:"file_extension" xml:"file_extension" required:"true"`
	// The content type of the additional format.
	ContentType string `json:"content_type" xml:"content_type" required:"true"`
	// Whether the content is base64 encoded.
	IsBase64Encoded bool `json:"is_base64_encoded" xml:"is_base64_encoded" required:"true"`
	// The content of the additional format.
	Content         string `json:"content" xml:"content" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AdditionalFormatResponseModel) GetRequestedFormat() (value string) {
	if a == nil {
		return
	}
	return a.RequestedFormat
}

func (a *AdditionalFormatResponseModel) GetFileExtension() (value string) {
	if a == nil {
		return
	}
	return a.FileExtension
}

func (a *AdditionalFormatResponseModel) GetContentType() (value string) {
	if a == nil {
		return
	}
	return a.ContentType
}

func (a *AdditionalFormatResponseModel) GetIsBase64Encoded() (value bool) {
	if a == nil {
		return
	}
	return a.IsBase64Encoded
}

func (a *AdditionalFormatResponseModel) GetContent() (value string) {
	if a == nil {
		return
	}
	return a.Content
}

func (a *AdditionalFormatResponseModel) SetRequestedFormat(requestedFormat string) {
	a.RequestedFormat = requestedFormat
}

func (a *AdditionalFormatResponseModel) SetFileExtension(fileExtension string) {
	a.FileExtension = fileExtension
}

func (a *AdditionalFormatResponseModel) SetContentType(contentType string) {
	a.ContentType = contentType
}

func (a *AdditionalFormatResponseModel) SetIsBase64Encoded(isBase64Encoded bool) {
	a.IsBase64Encoded = isBase64Encoded
}

func (a *AdditionalFormatResponseModel) SetContent(content string) {
	a.Content = content
}

func (a *AdditionalFormatResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AdditionalFormatResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AdditionalFormatResponseModel to string"
	}
	return string(jsonData)
}

func (a *AdditionalFormatResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AdditionalFormatResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AdditionalFormatResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AdditionalFormatResponseModel) MarshalJSON() ([]byte, error) {
	type embed AdditionalFormatResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
