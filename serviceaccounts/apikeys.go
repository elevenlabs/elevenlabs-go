package serviceaccounts

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost struct {
	ServiceAccountUserId string `json:"-" url:"-"`
	Name                 string `json:"name" xml:"name" required:"true"`
	// The permissions of the XI API.
	Permissions *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions `json:"permissions" xml:"permissions" required:"true"`
	// The character limit of the XI API key. If provided this will limit the usage of this api key to n characters per month where n is the chosen value. Requests that incur charges will fail after reaching this monthly limit.
	CharacterLimit *int `json:"character_limit,omitempty" xml:"character_limit,omitempty"`
	// List of IP addresses or CIDR ranges allowed to use this API key. Each entry may be a CIDR range (e.g. '10.0.0.0/24') or a bare IP address (normalized to /32 or /128). On create, omit or pass null to allow all IPs. On update, omit to leave the allowlist unchanged, or pass "clear" to remove it.
	AllowedIps []string `json:"allowed_ips,omitempty" xml:"allowed_ips,omitempty" minItems:"1" maxItems:"100"`
	// Whether the holder of this key may disable it via the self-disable endpoint. On create, omit or pass null to use the workspace's default (enabled for non-Enterprise plans, disabled for Enterprise plans). On update, omit to leave it unchanged, or pass "clear" to reset it to the workspace default. Only honored for workspaces with self-disable access enabled.
	ThirdPartyDisableAllowed *bool `json:"third_party_disable_allowed,omitempty" xml:"third_party_disable_allowed,omitempty"`
	// Maximum concurrent text-to-speech/speech-to-speech requests for this API key. Only available for enterprise customers.
	TtsConcurrencyLimit *int `json:"tts_concurrency_limit,omitempty" xml:"tts_concurrency_limit,omitempty"`
	// Maximum concurrent dubbing requests for this API key. Only available for enterprise customers.
	DubbingConcurrencyLimit *int `json:"dubbing_concurrency_limit,omitempty" xml:"dubbing_concurrency_limit,omitempty"`
	// Maximum concurrent music generation requests for this API key. Only available for enterprise customers.
	MusicConcurrencyLimit *int `json:"music_concurrency_limit,omitempty" xml:"music_concurrency_limit,omitempty"`
	// Unix timestamp (in seconds) after which this API key stops working. When set, it must be at least 15 minutes and at most 90 days in the future. On create, omit or pass null for a key that never expires. On update, omit to leave it unchanged, or pass "clear" to remove the expiry (never expires). A workspace may disallow never-expiring keys, or allow them only on service account API keys; creating one or clearing an expiry then fails.
	ExpiresAtUnix *int `json:"expires_at_unix,omitempty" xml:"expires_at_unix,omitempty"`
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetServiceAccountUserId() (value string) {
	if b == nil {
		return
	}
	return b.ServiceAccountUserId
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetPermissions() (value *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) {
	if b == nil {
		return
	}
	return b.Permissions
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetCharacterLimit() (value *int) {
	if b == nil {
		return
	}
	return b.CharacterLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetAllowedIps() (value []string) {
	if b == nil {
		return
	}
	return b.AllowedIps
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetThirdPartyDisableAllowed() (value *bool) {
	if b == nil {
		return
	}
	return b.ThirdPartyDisableAllowed
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetTtsConcurrencyLimit() (value *int) {
	if b == nil {
		return
	}
	return b.TtsConcurrencyLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetDubbingConcurrencyLimit() (value *int) {
	if b == nil {
		return
	}
	return b.DubbingConcurrencyLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetMusicConcurrencyLimit() (value *int) {
	if b == nil {
		return
	}
	return b.MusicConcurrencyLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetExpiresAtUnix() (value *int) {
	if b == nil {
		return
	}
	return b.ExpiresAtUnix
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetServiceAccountUserId(serviceAccountUserId string) {
	b.ServiceAccountUserId = serviceAccountUserId
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetPermissions(permissions *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) {
	b.Permissions = permissions
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetCharacterLimit(characterLimit *int) {
	b.CharacterLimit = characterLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetAllowedIps(allowedIps []string) {
	b.AllowedIps = allowedIps
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetThirdPartyDisableAllowed(thirdPartyDisableAllowed *bool) {
	b.ThirdPartyDisableAllowed = thirdPartyDisableAllowed
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetTtsConcurrencyLimit(ttsConcurrencyLimit *int) {
	b.TtsConcurrencyLimit = ttsConcurrencyLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetDubbingConcurrencyLimit(dubbingConcurrencyLimit *int) {
	b.DubbingConcurrencyLimit = dubbingConcurrencyLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetMusicConcurrencyLimit(musicConcurrencyLimit *int) {
	b.MusicConcurrencyLimit = musicConcurrencyLimit
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) SetExpiresAtUnix(expiresAtUnix *int) {
	b.ExpiresAtUnix = expiresAtUnix
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// The permissions of the XI API.
type BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions struct {
	PermissionTypeList []types.PermissionType
	AllStringLiteral   string
	typ                string
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) GetPermissionTypeList() []types.PermissionType {
	if b == nil {
		return nil
	}
	return b.PermissionTypeList
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) GetAllStringLiteral() string {
	if b == nil {
		return ""
	}
	return b.AllStringLiteral
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) UnmarshalJSON(data []byte) error {
	var valuePermissionTypeList []types.PermissionType
	if err := json.Unmarshal(data, &valuePermissionTypeList); err == nil {
		b.typ = "PermissionTypeList"
		b.PermissionTypeList = valuePermissionTypeList
		return nil
	}
	var valueAllStringLiteral string
	if err := json.Unmarshal(data, &valueAllStringLiteral); err == nil {
		b.typ = "AllStringLiteral"
		b.AllStringLiteral = valueAllStringLiteral
		if b.AllStringLiteral != "all" {
			return fmt.Errorf("unexpected value for literal on type %T; expected all got %v", b, b.AllStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) MarshalJSON() ([]byte, error) {
	if b.typ == "PermissionTypeList" || b.PermissionTypeList != nil {
		return json.Marshal(b.PermissionTypeList)
	}
	if b.typ == "AllStringLiteral" || b.AllStringLiteral != "" {
		return json.Marshal(b.AllStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsVisitor interface {
	VisitPermissionTypeList([]types.PermissionType) error
	VisitAllStringLiteral(string) error
}

func (b *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) Accept(visitor BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsVisitor) error {
	if b.typ == "PermissionTypeList" || b.PermissionTypeList != nil {
		return visitor.VisitPermissionTypeList(b.PermissionTypeList)
	}
	if b.typ == "AllStringLiteral" || b.AllStringLiteral != "" {
		return visitor.VisitAllStringLiteral(b.AllStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

func NewBodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsWithAllStringLiteral() *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions {
	return &BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions{typ: "AllStringLiteral", AllStringLiteral: "all"}
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch struct {
	ApiKeyId             string `json:"-" url:"-"`
	ServiceAccountUserId string `json:"-" url:"-"`
	// Whether to enable or disable the API key.
	IsEnabled *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	// The name of the XI API key to use (used for identification purposes only).
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The permissions of the XI API.
	Permissions *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions `json:"permissions,omitempty" xml:"permissions,omitempty"`
	// The character limit of the XI API key. If provided this will limit the usage of this api key to n characters per month where n is the chosen value. Requests that incur charges will fail after reaching this monthly limit.
	CharacterLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit `json:"character_limit,omitempty" xml:"character_limit,omitempty"`
	// List of IP addresses or CIDR ranges allowed to use this API key. Each entry may be a CIDR range (e.g. '10.0.0.0/24') or a bare IP address (normalized to /32 or /128). On create, omit or pass null to allow all IPs. On update, omit to leave the allowlist unchanged, or pass "clear" to remove it.
	AllowedIps *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps `json:"allowed_ips,omitempty" xml:"allowed_ips,omitempty"`
	// Whether the holder of this key may disable it via the self-disable endpoint. On create, omit or pass null to use the workspace's default (enabled for non-Enterprise plans, disabled for Enterprise plans). On update, omit to leave it unchanged, or pass "clear" to reset it to the workspace default. Only honored for workspaces with self-disable access enabled.
	ThirdPartyDisableAllowed *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed `json:"third_party_disable_allowed,omitempty" xml:"third_party_disable_allowed,omitempty"`
	// Maximum concurrent text-to-speech/speech-to-speech requests for this API key. Only available for enterprise customers.
	TtsConcurrencyLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit `json:"tts_concurrency_limit,omitempty" xml:"tts_concurrency_limit,omitempty"`
	// Maximum concurrent dubbing requests for this API key. Only available for enterprise customers.
	DubbingConcurrencyLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit `json:"dubbing_concurrency_limit,omitempty" xml:"dubbing_concurrency_limit,omitempty"`
	// Maximum concurrent music generation requests for this API key. Only available for enterprise customers.
	MusicConcurrencyLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit `json:"music_concurrency_limit,omitempty" xml:"music_concurrency_limit,omitempty"`
	// Unix timestamp (in seconds) after which this API key stops working. When set, it must be at least 15 minutes and at most 90 days in the future. On create, omit or pass null for a key that never expires. On update, omit to leave it unchanged, or pass "clear" to remove the expiry (never expires). A workspace may disallow never-expiring keys, or allow them only on service account API keys; creating one or clearing an expiry then fails.
	ExpiresAtUnix *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix `json:"expires_at_unix,omitempty" xml:"expires_at_unix,omitempty"`
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetApiKeyId() (value string) {
	if b == nil {
		return
	}
	return b.ApiKeyId
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetServiceAccountUserId() (value string) {
	if b == nil {
		return
	}
	return b.ServiceAccountUserId
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetIsEnabled() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) {
	if b == nil {
		return
	}
	return b.IsEnabled
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetPermissions() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) {
	if b == nil {
		return
	}
	return b.Permissions
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetCharacterLimit() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) {
	if b == nil {
		return
	}
	return b.CharacterLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetAllowedIps() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) {
	if b == nil {
		return
	}
	return b.AllowedIps
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetThirdPartyDisableAllowed() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) {
	if b == nil {
		return
	}
	return b.ThirdPartyDisableAllowed
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetTtsConcurrencyLimit() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) {
	if b == nil {
		return
	}
	return b.TtsConcurrencyLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetDubbingConcurrencyLimit() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) {
	if b == nil {
		return
	}
	return b.DubbingConcurrencyLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetMusicConcurrencyLimit() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) {
	if b == nil {
		return
	}
	return b.MusicConcurrencyLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetExpiresAtUnix() (value *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) {
	if b == nil {
		return
	}
	return b.ExpiresAtUnix
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetApiKeyId(apiKeyId string) {
	b.ApiKeyId = apiKeyId
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetServiceAccountUserId(serviceAccountUserId string) {
	b.ServiceAccountUserId = serviceAccountUserId
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetIsEnabled(isEnabled *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) {
	b.IsEnabled = isEnabled
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetName(name *string) {
	b.Name = name
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetPermissions(permissions *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) {
	b.Permissions = permissions
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetCharacterLimit(characterLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) {
	b.CharacterLimit = characterLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetAllowedIps(allowedIps *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) {
	b.AllowedIps = allowedIps
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetThirdPartyDisableAllowed(thirdPartyDisableAllowed *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) {
	b.ThirdPartyDisableAllowed = thirdPartyDisableAllowed
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetTtsConcurrencyLimit(ttsConcurrencyLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) {
	b.TtsConcurrencyLimit = ttsConcurrencyLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetDubbingConcurrencyLimit(dubbingConcurrencyLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) {
	b.DubbingConcurrencyLimit = dubbingConcurrencyLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetMusicConcurrencyLimit(musicConcurrencyLimit *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) {
	b.MusicConcurrencyLimit = musicConcurrencyLimit
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetExpiresAtUnix(expiresAtUnix *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) {
	b.ExpiresAtUnix = expiresAtUnix
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// Whether to enable or disable the API key.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled struct {
	Boolean               bool
	NoUpdateStringLiteral string
	typ                   string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) GetBoolean() bool {
	if b == nil {
		return false
	}
	return b.Boolean
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) GetNoUpdateStringLiteral() string {
	if b == nil {
		return ""
	}
	return b.NoUpdateStringLiteral
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		b.typ = "Boolean"
		b.Boolean = valueBoolean
		return nil
	}
	var valueNoUpdateStringLiteral string
	if err := json.Unmarshal(data, &valueNoUpdateStringLiteral); err == nil {
		b.typ = "NoUpdateStringLiteral"
		b.NoUpdateStringLiteral = valueNoUpdateStringLiteral
		if b.NoUpdateStringLiteral != "no_update" {
			return fmt.Errorf("unexpected value for literal on type %T; expected no_update got %v", b, b.NoUpdateStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) MarshalJSON() ([]byte, error) {
	if b.typ == "Boolean" || b.Boolean != false {
		return json.Marshal(b.Boolean)
	}
	if b.typ == "NoUpdateStringLiteral" || b.NoUpdateStringLiteral != "" {
		return json.Marshal(b.NoUpdateStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabledVisitor interface {
	VisitBoolean(bool) error
	VisitNoUpdateStringLiteral(string) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabledVisitor) error {
	if b.typ == "Boolean" || b.Boolean != false {
		return visitor.VisitBoolean(b.Boolean)
	}
	if b.typ == "NoUpdateStringLiteral" || b.NoUpdateStringLiteral != "" {
		return visitor.VisitNoUpdateStringLiteral(b.NoUpdateStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabledWithNoUpdateStringLiteral() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled {
	return &BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled{typ: "NoUpdateStringLiteral", NoUpdateStringLiteral: "no_update"}
}

// The permissions of the XI API.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions struct {
	PermissionTypeList                                                                                  []types.PermissionType
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
	typ                                                                                                 string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) GetPermissionTypeList() []types.PermissionType {
	if b == nil {
		return nil
	}
	return b.PermissionTypeList
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) UnmarshalJSON(data []byte) error {
	var valuePermissionTypeList []types.PermissionType
	if err := json.Unmarshal(data, &valuePermissionTypeList); err == nil {
		b.typ = "PermissionTypeList"
		b.PermissionTypeList = valuePermissionTypeList
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) MarshalJSON() ([]byte, error) {
	if b.typ == "PermissionTypeList" || b.PermissionTypeList != nil {
		return json.Marshal(b.PermissionTypeList)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsVisitor interface {
	VisitPermissionTypeList([]types.PermissionType) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsVisitor) error {
	if b.typ == "PermissionTypeList" || b.PermissionTypeList != nil {
		return visitor.VisitPermissionTypeList(b.PermissionTypeList)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneAll      BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne = "all"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne, error) {
	switch s {
	case "all":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneAll, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne {
	return &e
}

// The character limit of the XI API key. If provided this will limit the usage of this api key to n characters per month where n is the chosen value. Requests that incur charges will fail after reaching this monthly limit.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit struct {
	Integer                                                                                                int
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	typ                                                                                                    string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) GetInteger() int {
	if b == nil {
		return 0
	}
	return b.Integer
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		b.typ = "Integer"
		b.Integer = valueInteger
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) MarshalJSON() ([]byte, error) {
	if b.typ == "Integer" || b.Integer != 0 {
		return json.Marshal(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitVisitor interface {
	VisitInteger(int) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitVisitor) error {
	if b.typ == "Integer" || b.Integer != 0 {
		return visitor.VisitInteger(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne {
	return &e
}

// List of IP addresses or CIDR ranges allowed to use this API key. Each entry may be a CIDR range (e.g. '10.0.0.0/24') or a bare IP address (normalized to /32 or /128). On create, omit or pass null to allow all IPs. On update, omit to leave the allowlist unchanged, or pass "clear" to remove it.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps struct {
	StringList                                                                                         []string
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	typ                                                                                                string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) GetStringList() []string {
	if b == nil {
		return nil
	}
	return b.StringList
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) UnmarshalJSON(data []byte) error {
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		b.typ = "StringList"
		b.StringList = valueStringList
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) MarshalJSON() ([]byte, error) {
	if b.typ == "StringList" || b.StringList != nil {
		return json.Marshal(b.StringList)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsVisitor interface {
	VisitStringList([]string) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsVisitor) error {
	if b.typ == "StringList" || b.StringList != nil {
		return visitor.VisitStringList(b.StringList)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne {
	return &e
}

// Whether the holder of this key may disable it via the self-disable endpoint. On create, omit or pass null to use the workspace's default (enabled for non-Enterprise plans, disabled for Enterprise plans). On update, omit to leave it unchanged, or pass "clear" to reset it to the workspace default. Only honored for workspaces with self-disable access enabled.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed struct {
	Boolean                                                                                                          bool
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
	typ                                                                                                              string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) GetBoolean() bool {
	if b == nil {
		return false
	}
	return b.Boolean
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		b.typ = "Boolean"
		b.Boolean = valueBoolean
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) MarshalJSON() ([]byte, error) {
	if b.typ == "Boolean" || b.Boolean != false {
		return json.Marshal(b.Boolean)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedVisitor interface {
	VisitBoolean(bool) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedVisitor) error {
	if b.typ == "Boolean" || b.Boolean != false {
		return visitor.VisitBoolean(b.Boolean)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne {
	return &e
}

// Maximum concurrent text-to-speech/speech-to-speech requests for this API key. Only available for enterprise customers.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit struct {
	Integer                                                                                                     int
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
	typ                                                                                                         string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) GetInteger() int {
	if b == nil {
		return 0
	}
	return b.Integer
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		b.typ = "Integer"
		b.Integer = valueInteger
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if b.typ == "Integer" || b.Integer != 0 {
		return json.Marshal(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitVisitor) error {
	if b.typ == "Integer" || b.Integer != 0 {
		return visitor.VisitInteger(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne {
	return &e
}

// Maximum concurrent dubbing requests for this API key. Only available for enterprise customers.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit struct {
	Integer                                                                                                         int
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
	typ                                                                                                             string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) GetInteger() int {
	if b == nil {
		return 0
	}
	return b.Integer
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		b.typ = "Integer"
		b.Integer = valueInteger
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if b.typ == "Integer" || b.Integer != 0 {
		return json.Marshal(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitVisitor) error {
	if b.typ == "Integer" || b.Integer != 0 {
		return visitor.VisitInteger(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne {
	return &e
}

// Maximum concurrent music generation requests for this API key. Only available for enterprise customers.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit struct {
	Integer                                                                                                       int
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
	typ                                                                                                           string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) GetInteger() int {
	if b == nil {
		return 0
	}
	return b.Integer
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		b.typ = "Integer"
		b.Integer = valueInteger
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if b.typ == "Integer" || b.Integer != 0 {
		return json.Marshal(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitVisitor) error {
	if b.typ == "Integer" || b.Integer != 0 {
		return visitor.VisitInteger(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne {
	return &e
}

// Unix timestamp (in seconds) after which this API key stops working. When set, it must be at least 15 minutes and at most 90 days in the future. On create, omit or pass null for a key that never expires. On update, omit to leave it unchanged, or pass "clear" to remove the expiry (never expires). A workspace may disallow never-expiring keys, or allow them only on service account API keys; creating one or clearing an expiry then fails.
type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix struct {
	Integer                                                                                               int
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	typ                                                                                                   string
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) GetInteger() int {
	if b == nil {
		return 0
	}
	return b.Integer
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) GetBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne() BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne {
	if b == nil {
		return ""
	}
	return b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		b.typ = "Integer"
		b.Integer = valueInteger
		return nil
	}
	var valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	if err := json.Unmarshal(data, &valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne); err == nil {
		b.typ = "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne"
		b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = valueBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) MarshalJSON() ([]byte, error) {
	if b.typ == "Integer" || b.Integer != 0 {
		return json.Marshal(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne != "" {
		return json.Marshal(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixVisitor interface {
	VisitInteger(int) error
	VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne(BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne) error
}

func (b *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) Accept(visitor BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixVisitor) error {
	if b.typ == "Integer" || b.Integer != 0 {
		return visitor.VisitInteger(b.Integer)
	}
	if b.typ == "BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne" || b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne != "" {
		return visitor.VisitBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne(b.BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne string

const (
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneClear    BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = "clear"
	BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneNoUpdate BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = "no_update"
)

func NewBodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneFromString(s string) (BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne, error) {
	switch s {
	case "clear":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneClear, nil
	case "no_update":
		return BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneNoUpdate, nil
	}
	var t BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne) Ptr() *BodyEditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne {
	return &e
}
