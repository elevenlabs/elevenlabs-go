package templates

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Request body of “POST /v1/flows/templates/{template_id}/runs“.
type TemplateRunCreateRequest struct {
	// The ID of the template, as shown in the ElevenLabs app or by `GET /v1/flows/templates`.
	TemplateId string `json:"-" url:"-"`
	// Input values keyed by input port id. Every input port of the version being run must be given; a missing or unknown id is rejected. Pass `{}` for a template with no inputs.
	Inputs map[string]*types.TemplateRunInput `json:"inputs" required:"true"`
	// The template snapshot to run. Pass a specific version id to pin that snapshot, or `latest` (the default when omitted) to run the template's most recently published version. Only published versions can be pinned, except by the template's owner, who may also pin an unpublished saved snapshot to try it out before publishing. The live draft is never run through this API.
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty" minLength:"1"`
	// Include to send the run's result to the workspace's configured flows webhooks once the run's `status` reaches `completed` or `failed`. One event for the whole run: the `flows_template_run` event's `data` matches the terminal response of `GET /v1/flows/templates/{template_id}/runs/{run_id}`.
	Webhook *types.WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
}

func (t *TemplateRunCreateRequest) GetTemplateId() (value string) {
	if t == nil {
		return
	}
	return t.TemplateId
}

func (t *TemplateRunCreateRequest) GetInputs() (value map[string]*types.TemplateRunInput) {
	if t == nil {
		return
	}
	return t.Inputs
}

func (t *TemplateRunCreateRequest) GetVersionId() (value *string) {
	if t == nil {
		return
	}
	return t.VersionId
}

func (t *TemplateRunCreateRequest) GetWebhook() (value *types.WebhookTarget) {
	if t == nil {
		return
	}
	return t.Webhook
}

func (t *TemplateRunCreateRequest) SetTemplateId(templateId string) {
	t.TemplateId = templateId
}

func (t *TemplateRunCreateRequest) SetInputs(inputs map[string]*types.TemplateRunInput) {
	t.Inputs = inputs
}

func (t *TemplateRunCreateRequest) SetVersionId(versionId *string) {
	t.VersionId = versionId
}

func (t *TemplateRunCreateRequest) SetWebhook(webhook *types.WebhookTarget) {
	t.Webhook = webhook
}

func (t *TemplateRunCreateRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (t TemplateRunCreateRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateRunCreateRequest to string"
	}
	return string(jsonData)
}

func (t *TemplateRunCreateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, t)
}

func (t *TemplateRunCreateRequest) MarshalJSON() ([]byte, error) {
	type embed TemplateRunCreateRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}

var (
	listRunsRequestFieldCursor    = big.NewInt(1 << 0)
	listRunsRequestFieldPageSize  = big.NewInt(1 << 1)
	listRunsRequestFieldVersionId = big.NewInt(1 << 2)
)

type ListRunsRequest struct {
	Cursor    *string `json:"-" url:"cursor,omitempty"`
	PageSize  *int    `json:"-" url:"page_size,omitempty"`
	VersionId *string `json:"-" url:"version_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListRunsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListRunsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listRunsRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListRunsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listRunsRequestFieldPageSize)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListRunsRequest) SetVersionId(versionId *string) {
	l.VersionId = versionId
	l.require(listRunsRequestFieldVersionId)
}
