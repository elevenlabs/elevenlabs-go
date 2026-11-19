package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttentionView struct {
	AttentionId string `json:"attention_id" xml:"attention_id" required:"true"`
	CampaignId  string `json:"campaign_id" xml:"campaign_id" required:"true"`
	// The cause, such as `agent_unavailable` or `phone_number_unavailable`.
	Reason string `json:"reason" xml:"reason" required:"true"`
	// What went wrong and how to fix it.
	Detail string          `json:"detail" xml:"detail" required:"true"`
	Status AttentionStatus `json:"status" xml:"status" required:"true"`
	// Whether a retry places the calls the item holds back. Items that can't be retried are closed by resolving them.
	Retryable bool `json:"retryable" xml:"retryable" required:"true"`
	// The note recorded on resolving the item.
	Resolution      *string `json:"resolution" xml:"resolution" required:"true"`
	CreatedAtUnix   int     `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	UpdatedAtUnix   int     `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AttentionView) GetAttentionId() (value string) {
	if a == nil {
		return
	}
	return a.AttentionId
}

func (a *AttentionView) GetCampaignId() (value string) {
	if a == nil {
		return
	}
	return a.CampaignId
}

func (a *AttentionView) GetReason() (value string) {
	if a == nil {
		return
	}
	return a.Reason
}

func (a *AttentionView) GetDetail() (value string) {
	if a == nil {
		return
	}
	return a.Detail
}

func (a *AttentionView) GetStatus() (value AttentionStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AttentionView) GetRetryable() (value bool) {
	if a == nil {
		return
	}
	return a.Retryable
}

func (a *AttentionView) GetResolution() (value *string) {
	if a == nil {
		return
	}
	return a.Resolution
}

func (a *AttentionView) GetCreatedAtUnix() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnix
}

func (a *AttentionView) GetUpdatedAtUnix() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAtUnix
}

func (a *AttentionView) SetAttentionId(attentionId string) {
	a.AttentionId = attentionId
}

func (a *AttentionView) SetCampaignId(campaignId string) {
	a.CampaignId = campaignId
}

func (a *AttentionView) SetReason(reason string) {
	a.Reason = reason
}

func (a *AttentionView) SetDetail(detail string) {
	a.Detail = detail
}

func (a *AttentionView) SetStatus(status AttentionStatus) {
	a.Status = status
}

func (a *AttentionView) SetRetryable(retryable bool) {
	a.Retryable = retryable
}

func (a *AttentionView) SetResolution(resolution *string) {
	a.Resolution = resolution
}

func (a *AttentionView) SetCreatedAtUnix(createdAtUnix int) {
	a.CreatedAtUnix = createdAtUnix
}

func (a *AttentionView) SetUpdatedAtUnix(updatedAtUnix int) {
	a.UpdatedAtUnix = updatedAtUnix
}

func (a *AttentionView) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttentionView) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttentionView to string"
	}
	return string(jsonData)
}

func (a *AttentionView) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AttentionView) MarshalJSON() ([]byte, error) {
	type embed AttentionView
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
