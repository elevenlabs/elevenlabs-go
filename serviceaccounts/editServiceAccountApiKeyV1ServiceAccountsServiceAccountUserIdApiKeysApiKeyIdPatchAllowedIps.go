package serviceaccounts

import (
	"encoding/json"
	"fmt"
)

// List of IP addresses or CIDR ranges allowed to use this API key. Each entry may be a CIDR range (e.g. '10.0.0.0/24') or a bare IP address (normalized to /32 or /128). On create, omit or pass null to allow all IPs. On update, omit to leave the allowlist unchanged, or pass "clear" to remove it.
type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps struct {
	StringList                                                                   []string
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	typ                                                                          string
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) GetStringList() []string {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps == nil {
		return nil
	}
	return editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.StringList
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps == nil {
		return ""
	}
	return editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) UnmarshalJSON(data []byte) error {
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.typ = "StringList"
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.StringList = valueStringList
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne); err == nil {
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne"
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps)
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) MarshalJSON() ([]byte, error) {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.typ == "StringList" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.StringList != nil {
		return json.Marshal(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.StringList)
	}
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne != "" {
		return json.Marshal(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps)
}

type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsVisitor interface {
	VisitStringList([]string) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne) error
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps) Accept(visitor EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsVisitor) error {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.typ == "StringList" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.StringList != nil {
		return visitor.VisitStringList(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.StringList)
	}
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIpsOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchAllowedIps)
}
