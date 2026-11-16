package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreateServiceAccountV1ServiceAccountsPost struct {
	Name string `json:"name" xml:"name" required:"true"`
	// List of groups with their permission levels to share with by default. Each entry should specify a group_id and a permission_level (admin, editor, or viewer).
	DefaultSharingGroups []*DefaultSharingGroupConfig `json:"default_sharing_groups,omitempty" xml:"default_sharing_groups,omitempty"`
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) GetDefaultSharingGroups() (value []*DefaultSharingGroupConfig) {
	if b == nil {
		return
	}
	return b.DefaultSharingGroups
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) SetDefaultSharingGroups(defaultSharingGroups []*DefaultSharingGroupConfig) {
	b.DefaultSharingGroups = defaultSharingGroups
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateServiceAccountV1ServiceAccountsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateServiceAccountV1ServiceAccountsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateServiceAccountV1ServiceAccountsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateServiceAccountV1ServiceAccountsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
