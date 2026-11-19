package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An asset uploaded via the assets API.
type StaticAssetReference struct {
	type_ string
	// The ID of an asset uploaded via the assets API (`POST /v1/assets`), as returned in that response's `asset_id`.
	AssetId         string `json:"asset_id" xml:"asset_id" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StaticAssetReference) GetAssetId() (value string) {
	if s == nil {
		return
	}
	return s.AssetId
}

func (s *StaticAssetReference) SetAssetId(assetId string) {
	s.AssetId = assetId
}

func (s *StaticAssetReference) Type() string {
	return s.type_
}

func (s *StaticAssetReference) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StaticAssetReference) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StaticAssetReference to string"
	}
	return string(jsonData)
}

func (s *StaticAssetReference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StaticAssetReference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StaticAssetReference(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "asset" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "asset", *literals.Type)
	}
	s.type_ = "asset"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StaticAssetReference) MarshalJSON() ([]byte, error) {
	type embed StaticAssetReference
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "asset",
	}
	return json.Marshal(marshaler)
}
