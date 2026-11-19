package serviceaccounts

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// The permissions of the XI API.
type BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions struct {
	PermissionTypeList []types.PermissionType
	AllStringLiteral   string
	typ                string
}

func (bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) GetPermissionTypeList() []types.PermissionType {
	if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions == nil {
		return nil
	}
	return bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.PermissionTypeList
}

func (bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) GetAllStringLiteral() string {
	if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions == nil {
		return ""
	}
	return bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral
}

func (bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) UnmarshalJSON(data []byte) error {
	var valuePermissionTypeList []types.PermissionType
	if err := json.Unmarshal(data, &valuePermissionTypeList); err == nil {
		bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.typ = "PermissionTypeList"
		bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.PermissionTypeList = valuePermissionTypeList
		return nil
	}
	var valueAllStringLiteral string
	if err := json.Unmarshal(data, &valueAllStringLiteral); err == nil {
		bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.typ = "AllStringLiteral"
		bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral = valueAllStringLiteral
		if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral != "all" {
			return fmt.Errorf("unexpected value for literal on type %T; expected all got %v", bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions, bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions)
}

func (bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) MarshalJSON() ([]byte, error) {
	if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.typ == "PermissionTypeList" || bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.PermissionTypeList != nil {
		return json.Marshal(bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.PermissionTypeList)
	}
	if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.typ == "AllStringLiteral" || bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral != "" {
		return json.Marshal(bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions)
}

type BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsVisitor interface {
	VisitPermissionTypeList([]types.PermissionType) error
	VisitAllStringLiteral(string) error
}

func (bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions) Accept(visitor BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsVisitor) error {
	if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.typ == "PermissionTypeList" || bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.PermissionTypeList != nil {
		return visitor.VisitPermissionTypeList(bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.PermissionTypeList)
	}
	if bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.typ == "AllStringLiteral" || bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral != "" {
		return visitor.VisitAllStringLiteral(bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions.AllStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", bodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions)
}

func NewBodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissionsWithAllStringLiteral() *BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions {
	return &BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPostPermissions{typ: "AllStringLiteral", AllStringLiteral: "all"}
}
