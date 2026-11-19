package types

import (
	"fmt"
)

// How to break down the information. Cannot be "user" or "api_key" if include_workspace_metrics is False.
type BreakdownTypes string

const (
	BreakdownTypesNone                 BreakdownTypes = "none"
	BreakdownTypesVoice                BreakdownTypes = "voice"
	BreakdownTypesVoiceMultiplier      BreakdownTypes = "voice_multiplier"
	BreakdownTypesUser                 BreakdownTypes = "user"
	BreakdownTypesGroups               BreakdownTypes = "groups"
	BreakdownTypesApiKeys              BreakdownTypes = "api_keys"
	BreakdownTypesAllApiKeys           BreakdownTypes = "all_api_keys"
	BreakdownTypesProductType          BreakdownTypes = "product_type"
	BreakdownTypesModel                BreakdownTypes = "model"
	BreakdownTypesResource             BreakdownTypes = "resource"
	BreakdownTypesRequestQueue         BreakdownTypes = "request_queue"
	BreakdownTypesRegion               BreakdownTypes = "region"
	BreakdownTypesSubresourceId        BreakdownTypes = "subresource_id"
	BreakdownTypesReportingWorkspaceId BreakdownTypes = "reporting_workspace_id"
	BreakdownTypesHasApiKey            BreakdownTypes = "has_api_key"
	BreakdownTypesRequestSource        BreakdownTypes = "request_source"
)

func NewBreakdownTypesFromString(s string) (BreakdownTypes, error) {
	switch s {
	case "none":
		return BreakdownTypesNone, nil
	case "voice":
		return BreakdownTypesVoice, nil
	case "voice_multiplier":
		return BreakdownTypesVoiceMultiplier, nil
	case "user":
		return BreakdownTypesUser, nil
	case "groups":
		return BreakdownTypesGroups, nil
	case "api_keys":
		return BreakdownTypesApiKeys, nil
	case "all_api_keys":
		return BreakdownTypesAllApiKeys, nil
	case "product_type":
		return BreakdownTypesProductType, nil
	case "model":
		return BreakdownTypesModel, nil
	case "resource":
		return BreakdownTypesResource, nil
	case "request_queue":
		return BreakdownTypesRequestQueue, nil
	case "region":
		return BreakdownTypesRegion, nil
	case "subresource_id":
		return BreakdownTypesSubresourceId, nil
	case "reporting_workspace_id":
		return BreakdownTypesReportingWorkspaceId, nil
	case "has_api_key":
		return BreakdownTypesHasApiKey, nil
	case "request_source":
		return BreakdownTypesRequestSource, nil
	}
	var t BreakdownTypes
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BreakdownTypes) Ptr() *BreakdownTypes {
	return &e
}
