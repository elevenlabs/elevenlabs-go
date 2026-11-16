package orders

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost struct {
	// The ID of the order to which this media will be attached.
	OrderId OrderId `json:"-" url:"-"`
	// The language code of the media content (e.g. 'en', 'es-ES'). Must be a supported source language for some order item kind.
	DeclaredLanguage string `json:"declared_language" xml:"declared_language" required:"true"`
	// The media file to upload. Mutually exclusive with media_url.
	Media *string `json:"media,omitempty" xml:"media,omitempty"`
	// A URL to fetch the media file from. Mutually exclusive with media.
	MediaUrl *string `json:"media_url,omitempty" xml:"media_url,omitempty"`
	// The filename for URL-sourced media (e.g. 'example.mp4'). Required when using media_url.
	MediaUrlFilename *string `json:"media_url_filename,omitempty" xml:"media_url_filename,omitempty"`
	// The MIME type for URL-sourced media (e.g. 'video/mp4'). Required when using media_url.
	MediaUrlContentType *string `json:"media_url_content_type,omitempty" xml:"media_url_content_type,omitempty"`
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetOrderId() (value OrderId) {
	if b == nil {
		return
	}
	return b.OrderId
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetDeclaredLanguage() (value string) {
	if b == nil {
		return
	}
	return b.DeclaredLanguage
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetMedia() (value *string) {
	if b == nil {
		return
	}
	return b.Media
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetMediaUrl() (value *string) {
	if b == nil {
		return
	}
	return b.MediaUrl
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetMediaUrlFilename() (value *string) {
	if b == nil {
		return
	}
	return b.MediaUrlFilename
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetMediaUrlContentType() (value *string) {
	if b == nil {
		return
	}
	return b.MediaUrlContentType
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) SetOrderId(orderId OrderId) {
	b.OrderId = orderId
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) SetDeclaredLanguage(declaredLanguage string) {
	b.DeclaredLanguage = declaredLanguage
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) SetMedia(media *string) {
	b.Media = media
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) SetMediaUrl(mediaUrl *string) {
	b.MediaUrl = mediaUrl
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) SetMediaUrlFilename(mediaUrlFilename *string) {
	b.MediaUrlFilename = mediaUrlFilename
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) SetMediaUrlContentType(mediaUrlContentType *string) {
	b.MediaUrlContentType = mediaUrlContentType
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost to string"
	}
	return string(jsonData)
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost) MarshalJSON() ([]byte, error) {
	type embed BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
