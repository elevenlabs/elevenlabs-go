package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceApiKeyResponseModel struct {
	Name                 string           `json:"name" xml:"name" required:"true"`
	Hint                 string           `json:"hint" xml:"hint" required:"true"`
	KeyId                string           `json:"key_id" xml:"key_id" required:"true"`
	ServiceAccountUserId string           `json:"service_account_user_id" xml:"service_account_user_id" required:"true"`
	CreatedAtUnix        *int             `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	IsDisabled           *bool            `json:"is_disabled,omitempty" xml:"is_disabled,omitempty"`
	Permissions          []PermissionType `json:"permissions,omitempty" xml:"permissions,omitempty"`
	DisableReason        *LockReason      `json:"disable_reason,omitempty" xml:"disable_reason,omitempty"`
	// Maximum number of credits allowed in the current billing period.
	CharacterLimit *int `json:"character_limit,omitempty" xml:"character_limit,omitempty"`
	// Credits already used in the current billing period.
	CharacterCount           *int     `json:"character_count,omitempty" xml:"character_count,omitempty"`
	HashedXiApiKey           string   `json:"hashed_xi_api_key" xml:"hashed_xi_api_key" required:"true"`
	AllowedIps               []string `json:"allowed_ips,omitempty" xml:"allowed_ips,omitempty" minItems:"1" maxItems:"100"`
	ThirdPartyDisableAllowed *bool    `json:"third_party_disable_allowed,omitempty" xml:"third_party_disable_allowed,omitempty"`
	// Per-API-key concurrency limits (TTS/dubbing/music). Enterprise-only.
	PlatformLimits *PlatformLimits `json:"platform_limits,omitempty" xml:"platform_limits,omitempty"`
	// Unix timestamp (seconds) after which this API key stops working. Null means the key never expires.
	ExpiresAtUnix   *int `json:"expires_at_unix,omitempty" xml:"expires_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceApiKeyResponseModel) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WorkspaceApiKeyResponseModel) GetHint() (value string) {
	if w == nil {
		return
	}
	return w.Hint
}

func (w *WorkspaceApiKeyResponseModel) GetKeyId() (value string) {
	if w == nil {
		return
	}
	return w.KeyId
}

func (w *WorkspaceApiKeyResponseModel) GetServiceAccountUserId() (value string) {
	if w == nil {
		return
	}
	return w.ServiceAccountUserId
}

func (w *WorkspaceApiKeyResponseModel) GetCreatedAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.CreatedAtUnix
}

func (w *WorkspaceApiKeyResponseModel) GetIsDisabled() (value *bool) {
	if w == nil {
		return
	}
	return w.IsDisabled
}

func (w *WorkspaceApiKeyResponseModel) GetPermissions() (value []PermissionType) {
	if w == nil {
		return
	}
	return w.Permissions
}

func (w *WorkspaceApiKeyResponseModel) GetDisableReason() (value *LockReason) {
	if w == nil {
		return
	}
	return w.DisableReason
}

func (w *WorkspaceApiKeyResponseModel) GetCharacterLimit() (value *int) {
	if w == nil {
		return
	}
	return w.CharacterLimit
}

func (w *WorkspaceApiKeyResponseModel) GetCharacterCount() (value *int) {
	if w == nil {
		return
	}
	return w.CharacterCount
}

func (w *WorkspaceApiKeyResponseModel) GetHashedXiApiKey() (value string) {
	if w == nil {
		return
	}
	return w.HashedXiApiKey
}

func (w *WorkspaceApiKeyResponseModel) GetAllowedIps() (value []string) {
	if w == nil {
		return
	}
	return w.AllowedIps
}

func (w *WorkspaceApiKeyResponseModel) GetThirdPartyDisableAllowed() (value *bool) {
	if w == nil {
		return
	}
	return w.ThirdPartyDisableAllowed
}

func (w *WorkspaceApiKeyResponseModel) GetPlatformLimits() (value *PlatformLimits) {
	if w == nil {
		return
	}
	return w.PlatformLimits
}

func (w *WorkspaceApiKeyResponseModel) GetExpiresAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.ExpiresAtUnix
}

func (w *WorkspaceApiKeyResponseModel) SetName(name string) {
	w.Name = name
}

func (w *WorkspaceApiKeyResponseModel) SetHint(hint string) {
	w.Hint = hint
}

func (w *WorkspaceApiKeyResponseModel) SetKeyId(keyId string) {
	w.KeyId = keyId
}

func (w *WorkspaceApiKeyResponseModel) SetServiceAccountUserId(serviceAccountUserId string) {
	w.ServiceAccountUserId = serviceAccountUserId
}

func (w *WorkspaceApiKeyResponseModel) SetCreatedAtUnix(createdAtUnix *int) {
	w.CreatedAtUnix = createdAtUnix
}

func (w *WorkspaceApiKeyResponseModel) SetIsDisabled(isDisabled *bool) {
	w.IsDisabled = isDisabled
}

func (w *WorkspaceApiKeyResponseModel) SetPermissions(permissions []PermissionType) {
	w.Permissions = permissions
}

func (w *WorkspaceApiKeyResponseModel) SetDisableReason(disableReason *LockReason) {
	w.DisableReason = disableReason
}

func (w *WorkspaceApiKeyResponseModel) SetCharacterLimit(characterLimit *int) {
	w.CharacterLimit = characterLimit
}

func (w *WorkspaceApiKeyResponseModel) SetCharacterCount(characterCount *int) {
	w.CharacterCount = characterCount
}

func (w *WorkspaceApiKeyResponseModel) SetHashedXiApiKey(hashedXiApiKey string) {
	w.HashedXiApiKey = hashedXiApiKey
}

func (w *WorkspaceApiKeyResponseModel) SetAllowedIps(allowedIps []string) {
	w.AllowedIps = allowedIps
}

func (w *WorkspaceApiKeyResponseModel) SetThirdPartyDisableAllowed(thirdPartyDisableAllowed *bool) {
	w.ThirdPartyDisableAllowed = thirdPartyDisableAllowed
}

func (w *WorkspaceApiKeyResponseModel) SetPlatformLimits(platformLimits *PlatformLimits) {
	w.PlatformLimits = platformLimits
}

func (w *WorkspaceApiKeyResponseModel) SetExpiresAtUnix(expiresAtUnix *int) {
	w.ExpiresAtUnix = expiresAtUnix
}

func (w *WorkspaceApiKeyResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceApiKeyResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceApiKeyResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceApiKeyResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceApiKeyResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceApiKeyResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
