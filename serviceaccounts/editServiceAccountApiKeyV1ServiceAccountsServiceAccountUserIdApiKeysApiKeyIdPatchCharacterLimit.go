package serviceaccounts

import (
	"encoding/json"
	"fmt"
)

// The character limit of the XI API key. If provided this will limit the usage of this api key to n characters per month where n is the chosen value. Requests that incur charges will fail after reaching this monthly limit.
type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit struct {
	Integer                                                                          int
	ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	typ                                                                              string
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) GetInteger() int {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit == nil {
		return 0
	}
	return editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.Integer
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) GetApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne() ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit == nil {
		return ""
	}
	return editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) UnmarshalJSON(data []byte) error {
	var valueInteger int
	if err := json.Unmarshal(data, &valueInteger); err == nil {
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.typ = "Integer"
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.Integer = valueInteger
		return nil
	}
	var valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
	if err := json.Unmarshal(data, &valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne); err == nil {
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.typ = "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne"
		editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne = valueApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit)
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) MarshalJSON() ([]byte, error) {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.typ == "Integer" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.Integer != 0 {
		return json.Marshal(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.Integer)
	}
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne != "" {
		return json.Marshal(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit)
}

type EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitVisitor interface {
	VisitInteger(int) error
	VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne(ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne) error
}

func (editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit *EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit) Accept(visitor EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitVisitor) error {
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.typ == "Integer" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.Integer != 0 {
		return visitor.VisitInteger(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.Integer)
	}
	if editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.typ == "ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne" || editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne != "" {
		return visitor.VisitApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne(editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit.ApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimitOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", editServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatchCharacterLimit)
}
