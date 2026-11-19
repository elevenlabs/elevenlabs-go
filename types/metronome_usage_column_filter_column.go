package types

import (
	"fmt"
)

type MetronomeUsageColumnFilterColumn string

const (
	MetronomeUsageColumnFilterColumnModel                MetronomeUsageColumnFilterColumn = "model"
	MetronomeUsageColumnFilterColumnRequestSource        MetronomeUsageColumnFilterColumn = "request_source"
	MetronomeUsageColumnFilterColumnFiatChargeType       MetronomeUsageColumnFilterColumn = "fiat_charge_type"
	MetronomeUsageColumnFilterColumnFiatCurrency         MetronomeUsageColumnFilterColumn = "fiat_currency"
	MetronomeUsageColumnFilterColumnUserId               MetronomeUsageColumnFilterColumn = "user_id"
	MetronomeUsageColumnFilterColumnHashedXiApiKey       MetronomeUsageColumnFilterColumn = "hashed_xi_api_key"
	MetronomeUsageColumnFilterColumnRegion               MetronomeUsageColumnFilterColumn = "region"
	MetronomeUsageColumnFilterColumnReportingWorkspaceId MetronomeUsageColumnFilterColumn = "reporting_workspace_id"
	MetronomeUsageColumnFilterColumnBillingGroupId       MetronomeUsageColumnFilterColumn = "billing_group_id"
	MetronomeUsageColumnFilterColumnProductType          MetronomeUsageColumnFilterColumn = "product_type"
	MetronomeUsageColumnFilterColumnSurface              MetronomeUsageColumnFilterColumn = "surface"
	MetronomeUsageColumnFilterColumnActor                MetronomeUsageColumnFilterColumn = "actor"
)

func NewMetronomeUsageColumnFilterColumnFromString(s string) (MetronomeUsageColumnFilterColumn, error) {
	switch s {
	case "model":
		return MetronomeUsageColumnFilterColumnModel, nil
	case "request_source":
		return MetronomeUsageColumnFilterColumnRequestSource, nil
	case "fiat_charge_type":
		return MetronomeUsageColumnFilterColumnFiatChargeType, nil
	case "fiat_currency":
		return MetronomeUsageColumnFilterColumnFiatCurrency, nil
	case "user_id":
		return MetronomeUsageColumnFilterColumnUserId, nil
	case "hashed_xi_api_key":
		return MetronomeUsageColumnFilterColumnHashedXiApiKey, nil
	case "region":
		return MetronomeUsageColumnFilterColumnRegion, nil
	case "reporting_workspace_id":
		return MetronomeUsageColumnFilterColumnReportingWorkspaceId, nil
	case "billing_group_id":
		return MetronomeUsageColumnFilterColumnBillingGroupId, nil
	case "product_type":
		return MetronomeUsageColumnFilterColumnProductType, nil
	case "surface":
		return MetronomeUsageColumnFilterColumnSurface, nil
	case "actor":
		return MetronomeUsageColumnFilterColumnActor, nil
	}
	var t MetronomeUsageColumnFilterColumn
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MetronomeUsageColumnFilterColumn) Ptr() *MetronomeUsageColumnFilterColumn {
	return &e
}
