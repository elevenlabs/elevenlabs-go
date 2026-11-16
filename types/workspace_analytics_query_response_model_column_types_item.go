package types

import (
	"fmt"
)

type WorkspaceAnalyticsQueryResponseModelColumnTypesItem string

const (
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemString   WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "String"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemFloat    WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "Float"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemDateTime WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "DateTime"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemInt      WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "Int"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemBool     WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "Bool"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemJson     WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "JSON"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemMap      WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "Map"
	WorkspaceAnalyticsQueryResponseModelColumnTypesItemArray    WorkspaceAnalyticsQueryResponseModelColumnTypesItem = "Array"
)

func NewWorkspaceAnalyticsQueryResponseModelColumnTypesItemFromString(s string) (WorkspaceAnalyticsQueryResponseModelColumnTypesItem, error) {
	switch s {
	case "String":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemString, nil
	case "Float":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemFloat, nil
	case "DateTime":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemDateTime, nil
	case "Int":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemInt, nil
	case "Bool":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemBool, nil
	case "JSON":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemJson, nil
	case "Map":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemMap, nil
	case "Array":
		return WorkspaceAnalyticsQueryResponseModelColumnTypesItemArray, nil
	}
	var t WorkspaceAnalyticsQueryResponseModelColumnTypesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WorkspaceAnalyticsQueryResponseModelColumnTypesItem) Ptr() *WorkspaceAnalyticsQueryResponseModelColumnTypesItem {
	return &e
}
