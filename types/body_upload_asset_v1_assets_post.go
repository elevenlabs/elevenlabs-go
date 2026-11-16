package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyUploadAssetV1AssetsPost struct {
	// The file to upload.
	Asset io.Reader `json:"asset" xml:"asset" required:"true"`
	// Display name for the asset.
	Name string `json:"name" xml:"name" required:"true"`
}

func (b *BodyUploadAssetV1AssetsPost) GetAsset() (value io.Reader) {
	if b == nil {
		return
	}
	return b.Asset
}

func (b *BodyUploadAssetV1AssetsPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUploadAssetV1AssetsPost) SetAsset(asset io.Reader) {
	b.Asset = asset
}

func (b *BodyUploadAssetV1AssetsPost) SetName(name string) {
	b.Name = name
}

func (b *BodyUploadAssetV1AssetsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUploadAssetV1AssetsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUploadAssetV1AssetsPost to string"
	}
	return string(jsonData)
}

func (b *BodyUploadAssetV1AssetsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUploadAssetV1AssetsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUploadAssetV1AssetsPost(tmp)
	return nil
}

func (b *BodyUploadAssetV1AssetsPost) MarshalJSON() ([]byte, error) {
	type embed BodyUploadAssetV1AssetsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
