package serviceaccounts

import (
	"encoding/json"
	"fmt"
)

// Unix timestamp (in seconds) after which this API key stops working. When set, it must be at least 15 minutes and at most 90 days in the future. On create, omit or pass null for a key that never expires. On update, omit to leave it unchanged, or pass "clear" to remove the expiry (never expires). A workspace may disallow never-expiring keys, or allow them only on service account API keys; creating one or clearing an expiry then fails.
type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix struct {
	Integer                                                                         int
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	typ                                                                             string
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) GetInteger() int {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix == nil {
		return 0
	}
	return editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.Integer
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix == nil {
		return ""
	}
	return editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.typ = "Integer"
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.Integer = valueInteger
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne); err == nil {
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne"
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix)
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) MarshalJSON() ([]byte, error) {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.typ == "Integer" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.Integer != 0 {
		return json.Marshal(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.Integer)
	}
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne != "" {
		return json.Marshal(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix)
}

type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixVisitor interface {
	VisitInteger(int) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne) error
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix) Accept(visitor EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixVisitor) error {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.typ == "Integer" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.Integer != 0 {
		return visitor.VisitInteger(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.Integer)
	}
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnixOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchExpiresAtUnix)
}
