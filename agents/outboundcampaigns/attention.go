package outboundcampaigns

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type AttentionResolutionRequest struct {
	AttentionId string `json:"-" url:"-"`
	// How the item was handled.
	Resolution string `json:"resolution" xml:"resolution" required:"true" maxLength:"1000" minLength:"1"`
}

func (a *AttentionResolutionRequest) GetAttentionId() (value string) {
	if a == nil {
		return
	}
	return a.AttentionId
}

func (a *AttentionResolutionRequest) GetResolution() (value string) {
	if a == nil {
		return
	}
	return a.Resolution
}

func (a *AttentionResolutionRequest) SetAttentionId(attentionId string) {
	a.AttentionId = attentionId
}

func (a *AttentionResolutionRequest) SetResolution(resolution string) {
	a.Resolution = resolution
}

func (a *AttentionResolutionRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AttentionResolutionRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttentionResolutionRequest to string"
	}
	return string(jsonData)
}

func (a *AttentionResolutionRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AttentionResolutionRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AttentionResolutionRequest(tmp)
	return nil
}

func (a *AttentionResolutionRequest) MarshalJSON() ([]byte, error) {
	type embed AttentionResolutionRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

var (
	listAttentionRequestFieldCampaignId      = big.NewInt(1 << 0)
	listAttentionRequestFieldIncludeResolved = big.NewInt(1 << 1)
	listAttentionRequestFieldPageSize        = big.NewInt(1 << 2)
	listAttentionRequestFieldCursor          = big.NewInt(1 << 3)
)

type ListAttentionRequest struct {
	CampaignId      *string `json:"-" url:"campaign_id,omitempty"`
	IncludeResolved *bool   `json:"-" url:"include_resolved,omitempty"`
	PageSize        *int    `json:"-" url:"page_size,omitempty"`
	Cursor          *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListAttentionRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCampaignId sets the CampaignId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAttentionRequest) SetCampaignId(campaignId *string) {
	l.CampaignId = campaignId
	l.require(listAttentionRequestFieldCampaignId)
}

// SetIncludeResolved sets the IncludeResolved field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAttentionRequest) SetIncludeResolved(includeResolved *bool) {
	l.IncludeResolved = includeResolved
	l.require(listAttentionRequestFieldIncludeResolved)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAttentionRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listAttentionRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAttentionRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listAttentionRequestFieldCursor)
}
