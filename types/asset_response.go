package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The public Asset object.
type AssetResponse struct {
	// Unique identifier for the asset.
	AssetId string `json:"asset_id" xml:"asset_id" required:"true"`
	// Display name of the asset.
	Name string `json:"name" xml:"name" required:"true"`
	// MIME type of the uploaded file (e.g. `audio/mpeg`).
	MimeType string `json:"mime_type" xml:"mime_type" required:"true"`
	// Unix timestamp (seconds) the asset was created.
	CreatedAtUnix int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// Signed URL to fetch the asset's content. May be `null` if the asset has not finished processing. Do not rely on it being valid for more than 1 hour; fetch the asset again for a fresh URL.
	ContentUrl      *string `json:"content_url" xml:"content_url" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AssetResponse) GetAssetId() (value string) {
	if a == nil {
		return
	}
	return a.AssetId
}

func (a *AssetResponse) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AssetResponse) GetMimeType() (value string) {
	if a == nil {
		return
	}
	return a.MimeType
}

func (a *AssetResponse) GetCreatedAtUnix() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnix
}

func (a *AssetResponse) GetContentUrl() (value *string) {
	if a == nil {
		return
	}
	return a.ContentUrl
}

func (a *AssetResponse) SetAssetId(assetId string) {
	a.AssetId = assetId
}

func (a *AssetResponse) SetName(name string) {
	a.Name = name
}

func (a *AssetResponse) SetMimeType(mimeType string) {
	a.MimeType = mimeType
}

func (a *AssetResponse) SetCreatedAtUnix(createdAtUnix int) {
	a.CreatedAtUnix = createdAtUnix
}

func (a *AssetResponse) SetContentUrl(contentUrl *string) {
	a.ContentUrl = contentUrl
}

func (a *AssetResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AssetResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssetResponse to string"
	}
	return string(jsonData)
}

func (a *AssetResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AssetResponse) MarshalJSON() ([]byte, error) {
	type embed AssetResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
