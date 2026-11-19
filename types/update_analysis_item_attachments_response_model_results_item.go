package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateAnalysisItemAttachmentsResponseModelResultsItem struct {
	Status  string
	Failure *BatchFailureResponseModel
	Success *BatchSuccessResponseModelAnalysisItemAttachmentResultData
}

func (u *UpdateAnalysisItemAttachmentsResponseModelResultsItem) GetFailure() *BatchFailureResponseModel {
	if u == nil {
		return nil
	}
	if u.Failure == nil {
		return nil
	}
	return u.Failure
}

func (u *UpdateAnalysisItemAttachmentsResponseModelResultsItem) GetSuccess() *BatchSuccessResponseModelAnalysisItemAttachmentResultData {
	if u == nil {
		return nil
	}
	if u.Success == nil {
		return nil
	}
	return u.Success
}

func (u *UpdateAnalysisItemAttachmentsResponseModelResultsItem) GetStatus() string {
	if u == nil {
		return ""
	}
	return u.Status
}

func (u *UpdateAnalysisItemAttachmentsResponseModelResultsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", u)
	}
	switch unmarshaler.Status {
	case "failure":
		value := new(BatchFailureResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Failure = value
	case "success":
		value := new(BatchSuccessResponseModelAnalysisItemAttachmentResultData)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Success = value
	}
	return nil
}

func (u UpdateAnalysisItemAttachmentsResponseModelResultsItem) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Failure != nil {
		return internal.MarshalJSONWithDiscriminant(u.Failure, "status", "failure")
	}
	if u.Success != nil {
		return internal.MarshalJSONWithDiscriminant(u.Success, "status", "success")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateAnalysisItemAttachmentsResponseModelResultsItemVisitor interface {
	VisitFailure(*BatchFailureResponseModel) error
	VisitSuccess(*BatchSuccessResponseModelAnalysisItemAttachmentResultData) error
}

func (u *UpdateAnalysisItemAttachmentsResponseModelResultsItem) Accept(visitor UpdateAnalysisItemAttachmentsResponseModelResultsItemVisitor) error {
	if u.Failure != nil {
		return visitor.VisitFailure(u.Failure)
	}
	if u.Success != nil {
		return visitor.VisitSuccess(u.Success)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateAnalysisItemAttachmentsResponseModelResultsItem) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Failure != nil {
		fields = append(fields, "failure")
	}
	if u.Success != nil {
		fields = append(fields, "success")
	}
	if len(fields) == 0 {
		if u.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Status)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Status != "" {
		field := fields[0]
		if u.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Status,
				u,
			)
		}
	}
	return nil
}
