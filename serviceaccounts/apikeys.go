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
type CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions struct {
	PermissionTypeList []types.PermissionType
	AllStringLiteral   string
	typ                string
}

func (c *CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) GetPermissionTypeList() []types.PermissionType {
	if c == nil {
		return nil
	}
	return c.PermissionTypeList
}

func (c *CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) GetAllStringLiteral() string {
	if c == nil {
		return ""
	}
	return c.AllStringLiteral
}

func (c *CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) UnmarshalJSON(data []byte) error {
	var valuePermissionTypeList []types.PermissionType
	if err := json.Unmarshal(data, &valuePermissionTypeList); err == nil {
		c.typ = "PermissionTypeList"
		c.PermissionTypeList = valuePermissionTypeList
		return nil
	}
	var valueAllStringLiteral string
	if err := json.Unmarshal(data, &valueAllStringLiteral); err == nil {
		c.typ = "AllStringLiteral"
		c.AllStringLiteral = valueAllStringLiteral
		if c.AllStringLiteral != "all" {
			return fmt.Errorf("unexpected value for literal on type %T; expected all got %v", c, c.AllStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) MarshalJSON() ([]byte, error) {
	if c.typ == "PermissionTypeList" || c.PermissionTypeList != nil {
		return json.Marshal(c.PermissionTypeList)
	}
	if c.typ == "AllStringLiteral" || c.AllStringLiteral != "" {
		return json.Marshal(c.AllStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsVisitor interface {
	VisitPermissionTypeList([]types.PermissionType) error
	VisitAllStringLiteral(string) error
}

func (c *CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) Accept(visitor CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsVisitor) error {
	if c.typ == "PermissionTypeList" || c.PermissionTypeList != nil {
		return visitor.VisitPermissionTypeList(c.PermissionTypeList)
	}
	if c.typ == "AllStringLiteral" || c.AllStringLiteral != "" {
		return visitor.VisitAllStringLiteral(c.AllStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

func NewCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsWithAllStringLiteral() *CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions {
	return &CreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions{typ: "AllStringLiteral", AllStringLiteral: "all"}
}

type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch struct {
	// Whether to enable or disable the API key.
	IsEnabled *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	// The name of the XI API key to use (used for identification purposes only).
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The permissions of the XI API.
	Permissions *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions `json:"permissions,omitempty" xml:"permissions,omitempty"`
	// The character limit of the XI API key. If provided this will limit the usage of this api key to n characters per month where n is the chosen value. Requests that incur charges will fail after reaching this monthly limit.
	CharacterLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit `json:"character_limit,omitempty" xml:"character_limit,omitempty"`
	// List of IP addresses or CIDR ranges allowed to use this API key. Each entry may be a CIDR range (e.g. '10.0.0.0/24') or a bare IP address (normalized to /32 or /128). On create, omit or pass null to allow all IPs. On update, omit to leave the allowlist unchanged, or pass "clear" to remove it.
	AllowedIps *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps `json:"allowed_ips,omitempty" xml:"allowed_ips,omitempty"`
	// Whether the holder of this key may disable it via the self-disable endpoint. On create, omit or pass null to use the workspace's default (enabled for non-Enterprise plans, disabled for Enterprise plans). On update, omit to leave it unchanged, or pass "clear" to reset it to the workspace default. Only honored for workspaces with self-disable access enabled.
	ThirdPartyDisableAllowed *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed `json:"third_party_disable_allowed,omitempty" xml:"third_party_disable_allowed,omitempty"`
	// Maximum concurrent text-to-speech/speech-to-speech requests for this API key. Only available for enterprise customers.
	TtsConcurrencyLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit `json:"tts_concurrency_limit,omitempty" xml:"tts_concurrency_limit,omitempty"`
	// Maximum concurrent dubbing requests for this API key. Only available for enterprise customers.
	DubbingConcurrencyLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit `json:"dubbing_concurrency_limit,omitempty" xml:"dubbing_concurrency_limit,omitempty"`
	// Maximum concurrent music generation requests for this API key. Only available for enterprise customers.
	MusicConcurrencyLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit `json:"music_concurrency_limit,omitempty" xml:"music_concurrency_limit,omitempty"`
	// Unix timestamp (in seconds) after which this API key stops working. When set, it must be at least 15 minutes and at most 90 days in the future. On create, omit or pass null for a key that never expires. On update, omit to leave it unchanged, or pass "clear" to remove the expiry (never expires). A workspace may disallow never-expiring keys, or allow them only on service account API keys; creating one or clearing an expiry then fails.
	ExpiresAtUnix *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix `json:"expires_at_unix,omitempty" xml:"expires_at_unix,omitempty"`
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetIsEnabled() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) {
	if e == nil {
		return
	}
	return e.IsEnabled
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetName() (value *string) {
	if e == nil {
		return
	}
	return e.Name
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetPermissions() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) {
	if e == nil {
		return
	}
	return e.Permissions
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetCharacterLimit() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) {
	if e == nil {
		return
	}
	return e.CharacterLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetAllowedIps() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) {
	if e == nil {
		return
	}
	return e.AllowedIps
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetThirdPartyDisableAllowed() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) {
	if e == nil {
		return
	}
	return e.ThirdPartyDisableAllowed
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetTtsConcurrencyLimit() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) {
	if e == nil {
		return
	}
	return e.TtsConcurrencyLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetDubbingConcurrencyLimit() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) {
	if e == nil {
		return
	}
	return e.DubbingConcurrencyLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetMusicConcurrencyLimit() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) {
	if e == nil {
		return
	}
	return e.MusicConcurrencyLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetExpiresAtUnix() (value *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) {
	if e == nil {
		return
	}
	return e.ExpiresAtUnix
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetIsEnabled(isEnabled *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) {
	e.IsEnabled = isEnabled
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetName(name *string) {
	e.Name = name
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetPermissions(permissions *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) {
	e.Permissions = permissions
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetCharacterLimit(characterLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) {
	e.CharacterLimit = characterLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetAllowedIps(allowedIps *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) {
	e.AllowedIps = allowedIps
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetThirdPartyDisableAllowed(thirdPartyDisableAllowed *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) {
	e.ThirdPartyDisableAllowed = thirdPartyDisableAllowed
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetTtsConcurrencyLimit(ttsConcurrencyLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) {
	e.TtsConcurrencyLimit = ttsConcurrencyLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetDubbingConcurrencyLimit(dubbingConcurrencyLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) {
	e.DubbingConcurrencyLimit = dubbingConcurrencyLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetMusicConcurrencyLimit(musicConcurrencyLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) {
	e.MusicConcurrencyLimit = musicConcurrencyLimit
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) SetExpiresAtUnix(expiresAtUnix *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) {
	e.ExpiresAtUnix = expiresAtUnix
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (e EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch to string"
	}
	return string(jsonData)
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}

func (e *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch) MarshalJSON() ([]byte, error) {
	type embed EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}

// Whether to enable or disable the API key.
type AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled struct {
	Boolean               bool
	NoUpdateStringLiteral string
	typ                   string
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) GetNoUpdateStringLiteral() string {
	if a == nil {
		return ""
	}
	return a.NoUpdateStringLiteral
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueNoUpdateStringLiteral string
	if err := json.Unmarshal(data, &valueNoUpdateStringLiteral); err == nil {
		a.typ = "NoUpdateStringLiteral"
		a.NoUpdateStringLiteral = valueNoUpdateStringLiteral
		if a.NoUpdateStringLiteral != "no_update" {
			return fmt.Errorf("unexpected value for literal on type %T; expected no_update got %v", a, a.NoUpdateStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "NoUpdateStringLiteral" || a.NoUpdateStringLiteral != "" {
		return json.Marshal(a.NoUpdateStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabledVisitor interface {
	VisitBoolean(bool) error
	VisitNoUpdateStringLiteral(string) error
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled) Accept(visitor AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabledVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "NoUpdateStringLiteral" || a.NoUpdateStringLiteral != "" {
		return visitor.VisitNoUpdateStringLiteral(a.NoUpdateStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

func NewAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabledWithNoUpdateStringLiteral() *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled {
	return &AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchIsEnabled{typ: "NoUpdateStringLiteral", NoUpdateStringLiteral: "no_update"}
}

// The permissions of the XI API.
type AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions struct {
	PermissionTypeList                                                            []types.PermissionType
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
	typ                                                                           string
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) GetPermissionTypeList() []types.PermissionType {
	if a == nil {
		return nil
	}
	return a.PermissionTypeList
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne {
	if a == nil {
		return ""
	}
	return a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) UnmarshalJSON(data []byte) error {
	var valuePermissionTypeList []types.PermissionType
	if err := json.Unmarshal(data, &valuePermissionTypeList); err == nil {
		a.typ = "PermissionTypeList"
		a.PermissionTypeList = valuePermissionTypeList
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne); err == nil {
		a.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne"
		a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) MarshalJSON() ([]byte, error) {
	if a.typ == "PermissionTypeList" || a.PermissionTypeList != nil {
		return json.Marshal(a.PermissionTypeList)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne != "" {
		return json.Marshal(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsVisitor interface {
	VisitPermissionTypeList([]types.PermissionType) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne) error
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissions) Accept(visitor AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsVisitor) error {
	if a.typ == "PermissionTypeList" || a.PermissionTypeList != nil {
		return visitor.VisitPermissionTypeList(a.PermissionTypeList)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne string

const (
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneAll      ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne = "all"
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneNoUpdate ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne = "no_update"
)

func NewApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneFromString(s string) (ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne, error) {
	switch s {
	case "all":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneAll, nil
	case "no_update":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOneNoUpdate, nil
	}
	var t ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne) Ptr() *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchPermissionsOne {
	return &e
}

// The character limit of the XI API key. If provided this will limit the usage of this api key to n characters per month where n is the chosen value. Requests that incur charges will fail after reaching this monthly limit.
type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit struct {
	Integer                                                                          int
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	typ                                                                              string
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) GetInteger() int {
	if a == nil {
		return 0
	}
	return a.Integer
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne {
	if a == nil {
		return ""
	}
	return a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		a.typ = "Integer"
		a.Integer = valueInteger
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne); err == nil {
		a.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne"
		a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) MarshalJSON() ([]byte, error) {
	if a.typ == "Integer" || a.Integer != 0 {
		return json.Marshal(a.Integer)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne != "" {
		return json.Marshal(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitVisitor interface {
	VisitInteger(int) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne) error
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) Accept(visitor ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitVisitor) error {
	if a.typ == "Integer" || a.Integer != 0 {
		return visitor.VisitInteger(a.Integer)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne string

const (
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneClear    ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = "clear"
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneNoUpdate ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = "no_update"
)

func NewApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneFromString(s string) (ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne, error) {
	switch s {
	case "clear":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneClear, nil
	case "no_update":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOneNoUpdate, nil
	}
	var t ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne) Ptr() *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne {
	return &e
}

// List of IP addresses or CIDR ranges allowed to use this API key. Each entry may be a CIDR range (e.g. '10.0.0.0/24') or a bare IP address (normalized to /32 or /128). On create, omit or pass null to allow all IPs. On update, omit to leave the allowlist unchanged, or pass "clear" to remove it.
type AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps struct {
	StringList                                                                   []string
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	typ                                                                          string
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) GetStringList() []string {
	if a == nil {
		return nil
	}
	return a.StringList
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne {
	if a == nil {
		return ""
	}
	return a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) UnmarshalJSON(data []byte) error {
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		a.typ = "StringList"
		a.StringList = valueStringList
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne); err == nil {
		a.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne"
		a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) MarshalJSON() ([]byte, error) {
	if a.typ == "StringList" || a.StringList != nil {
		return json.Marshal(a.StringList)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne != "" {
		return json.Marshal(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsVisitor interface {
	VisitStringList([]string) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne) error
}

func (a *AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) Accept(visitor AccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsVisitor) error {
	if a.typ == "StringList" || a.StringList != nil {
		return visitor.VisitStringList(a.StringList)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne string

const (
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneClear    ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = "clear"
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneNoUpdate ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = "no_update"
)

func NewApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneFromString(s string) (ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne, error) {
	switch s {
	case "clear":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneClear, nil
	case "no_update":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOneNoUpdate, nil
	}
	var t ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne) Ptr() *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne {
	return &e
}

// Whether the holder of this key may disable it via the self-disable endpoint. On create, omit or pass null to use the workspace's default (enabled for non-Enterprise plans, disabled for Enterprise plans). On update, omit to leave it unchanged, or pass "clear" to reset it to the workspace default. Only honored for workspaces with self-disable access enabled.
type V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed struct {
	Boolean                                                                            bool
	ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
	typ                                                                                string
}

func (v *V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) GetBoolean() bool {
	if v == nil {
		return false
	}
	return v.Boolean
}

func (v *V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) GetServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne() ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne {
	if v == nil {
		return ""
	}
	return v.ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
}

func (v *V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		v.typ = "Boolean"
		v.Boolean = valueBoolean
		return nil
	}
	var valueServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
	if err := json.Unmarshal(data, &valueServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne); err == nil {
		v.typ = "ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne"
		v.ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne = valueServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, v)
}

func (v V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) MarshalJSON() ([]byte, error) {
	if v.typ == "Boolean" || v.Boolean != false {
		return json.Marshal(v.Boolean)
	}
	if v.typ == "ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne" || v.ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne != "" {
		return json.Marshal(v.ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedVisitor interface {
	VisitBoolean(bool) error
	VisitServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne(ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne) error
}

func (v *V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowed) Accept(visitor V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedVisitor) error {
	if v.typ == "Boolean" || v.Boolean != false {
		return visitor.VisitBoolean(v.Boolean)
	}
	if v.typ == "ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne" || v.ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne != "" {
		return visitor.VisitServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne(v.ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}

type ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne string

const (
	ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneClear    ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne = "clear"
	ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneNoUpdate ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne = "no_update"
)

func NewServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneFromString(s string) (ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne, error) {
	switch s {
	case "clear":
		return ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneClear, nil
	case "no_update":
		return ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOneNoUpdate, nil
	}
	var t ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne) Ptr() *ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchThirdPartyDisableAllowedOne {
	return &e
}

// Maximum concurrent text-to-speech/speech-to-speech requests for this API key. Only available for enterprise customers.
type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit struct {
	Integer                                                                            int
	KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
	typ                                                                                string
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) GetInteger() int {
	if a == nil {
		return 0
	}
	return a.Integer
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) GetKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne() KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne {
	if a == nil {
		return ""
	}
	return a.KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		a.typ = "Integer"
		a.Integer = valueInteger
		return nil
	}
	var valueKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
	if err := json.Unmarshal(data, &valueKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne); err == nil {
		a.typ = "KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne"
		a.KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne = valueKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if a.typ == "Integer" || a.Integer != 0 {
		return json.Marshal(a.Integer)
	}
	if a.typ == "KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne" || a.KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne != "" {
		return json.Marshal(a.KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne(KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne) error
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimit) Accept(visitor ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitVisitor) error {
	if a.typ == "Integer" || a.Integer != 0 {
		return visitor.VisitInteger(a.Integer)
	}
	if a.typ == "KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne" || a.KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne != "" {
		return visitor.VisitKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne(a.KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne string

const (
	KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneClear    KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne = "clear"
	KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneNoUpdate KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne = "no_update"
)

func NewKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneFromString(s string) (KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne, error) {
	switch s {
	case "clear":
		return KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneClear, nil
	case "no_update":
		return KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOneNoUpdate, nil
	}
	var t KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne) Ptr() *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchTtsConcurrencyLimitOne {
	return &e
}

// Maximum concurrent dubbing requests for this API key. Only available for enterprise customers.
type KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit struct {
	Integer                                                                             int
	V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
	typ                                                                                 string
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) GetInteger() int {
	if k == nil {
		return 0
	}
	return k.Integer
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) GetV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne() V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne {
	if k == nil {
		return ""
	}
	return k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		k.typ = "Integer"
		k.Integer = valueInteger
		return nil
	}
	var valueV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
	if err := json.Unmarshal(data, &valueV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne); err == nil {
		k.typ = "V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne"
		k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne = valueV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if k.typ == "Integer" || k.Integer != 0 {
		return json.Marshal(k.Integer)
	}
	if k.typ == "V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne" || k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne != "" {
		return json.Marshal(k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne(V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne) error
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimit) Accept(visitor KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitVisitor) error {
	if k.typ == "Integer" || k.Integer != 0 {
		return visitor.VisitInteger(k.Integer)
	}
	if k.typ == "V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne" || k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne != "" {
		return visitor.VisitV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne(k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne string

const (
	V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneClear    V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne = "clear"
	V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneNoUpdate V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne = "no_update"
)

func NewV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneFromString(s string) (V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne, error) {
	switch s {
	case "clear":
		return V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneClear, nil
	case "no_update":
		return V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOneNoUpdate, nil
	}
	var t V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne) Ptr() *V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchDubbingConcurrencyLimitOne {
	return &e
}

// Maximum concurrent music generation requests for this API key. Only available for enterprise customers.
type KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit struct {
	Integer                                                                           int
	V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
	typ                                                                               string
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) GetInteger() int {
	if k == nil {
		return 0
	}
	return k.Integer
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) GetV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne() V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne {
	if k == nil {
		return ""
	}
	return k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		k.typ = "Integer"
		k.Integer = valueInteger
		return nil
	}
	var valueV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
	if err := json.Unmarshal(data, &valueV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne); err == nil {
		k.typ = "V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne"
		k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne = valueV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) MarshalJSON() ([]byte, error) {
	if k.typ == "Integer" || k.Integer != 0 {
		return json.Marshal(k.Integer)
	}
	if k.typ == "V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne" || k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne != "" {
		return json.Marshal(k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitVisitor interface {
	VisitInteger(int) error
	VisitV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne(V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne) error
}

func (k *KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimit) Accept(visitor KeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitVisitor) error {
	if k.typ == "Integer" || k.Integer != 0 {
		return visitor.VisitInteger(k.Integer)
	}
	if k.typ == "V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne" || k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne != "" {
		return visitor.VisitV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne(k.V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne string

const (
	V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneClear    V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne = "clear"
	V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneNoUpdate V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne = "no_update"
)

func NewV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneFromString(s string) (V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne, error) {
	switch s {
	case "clear":
		return V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneClear, nil
	case "no_update":
		return V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOneNoUpdate, nil
	}
	var t V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne) Ptr() *V1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchMusicConcurrencyLimitOne {
	return &e
}

// Unix timestamp (in seconds) after which this API key stops working. When set, it must be at least 15 minutes and at most 90 days in the future. On create, omit or pass null for a key that never expires. On update, omit to leave it unchanged, or pass "clear" to remove the expiry (never expires). A workspace may disallow never-expiring keys, or allow them only on service account API keys; creating one or clearing an expiry then fails.
type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix struct {
	Integer                                                                         int
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	typ                                                                             string
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) GetInteger() int {
	if a == nil {
		return 0
	}
	return a.Integer
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne {
	if a == nil {
		return ""
	}
	return a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		a.typ = "Integer"
		a.Integer = valueInteger
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne); err == nil {
		a.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne"
		a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) MarshalJSON() ([]byte, error) {
	if a.typ == "Integer" || a.Integer != 0 {
		return json.Marshal(a.Integer)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne != "" {
		return json.Marshal(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixVisitor interface {
	VisitInteger(int) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne) error
}

func (a *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) Accept(visitor ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixVisitor) error {
	if a.typ == "Integer" || a.Integer != 0 {
		return visitor.VisitInteger(a.Integer)
	}
	if a.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne" || a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne(a.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne string

const (
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneClear    ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = "clear"
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneNoUpdate ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = "no_update"
)

func NewApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneFromString(s string) (ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne, error) {
	switch s {
	case "clear":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneClear, nil
	case "no_update":
		return ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOneNoUpdate, nil
	}
	var t ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne) Ptr() *ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne {
	return &e
}
