package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OrderMediaResponse struct {
	MediaId MediaId `json:"media_id" xml:"media_id" required:"true" pattern:"^prodmedia_[a-z0-9]{26}$"`
	// The original filename of the uploaded media.
	Name string `json:"name" xml:"name" required:"true"`
	// The MIME type of the media file (e.g. 'video/mp4').
	ContentType string `json:"content_type" xml:"content_type" required:"true"`
	// The detected or declared language of the media, if available.
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// A time-limited URL to download the media file.
	SignedUrl       string `json:"signed_url" xml:"signed_url" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OrderMediaResponse) GetMediaId() (value MediaId) {
	if o == nil {
		return
	}
	return o.MediaId
}

func (o *OrderMediaResponse) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OrderMediaResponse) GetContentType() (value string) {
	if o == nil {
		return
	}
	return o.ContentType
}

func (o *OrderMediaResponse) GetLanguage() (value *string) {
	if o == nil {
		return
	}
	return o.Language
}

func (o *OrderMediaResponse) GetSignedUrl() (value string) {
	if o == nil {
		return
	}
	return o.SignedUrl
}

func (o *OrderMediaResponse) SetMediaId(mediaId MediaId) {
	o.MediaId = mediaId
}

func (o *OrderMediaResponse) SetName(name string) {
	o.Name = name
}

func (o *OrderMediaResponse) SetContentType(contentType string) {
	o.ContentType = contentType
}

func (o *OrderMediaResponse) SetLanguage(language *string) {
	o.Language = language
}

func (o *OrderMediaResponse) SetSignedUrl(signedUrl string) {
	o.SignedUrl = signedUrl
}

func (o *OrderMediaResponse) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrderMediaResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrderMediaResponse to string"
	}
	return string(jsonData)
}

func (o *OrderMediaResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OrderMediaResponse) MarshalJSON() ([]byte, error) {
	type embed OrderMediaResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
