package dubbing

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type BodyCreateDubbingProjectV1DubbingProjectPost struct {
	// The source media file to dub: an audio or video file of at most 3 GiB. Provide this or `source_url`, not both.
	File *string `json:"file,omitempty" xml:"file,omitempty"`
	// Public HTTP(S) URL the source media is fetched from server-side, subject to the same size and format limits as an upload. Provide this or `file`, not both.
	SourceUrl *string `json:"source_url,omitempty" xml:"source_url,omitempty"`
	// Optional free-form string (at most 500 characters) to identify the project on your end. Stored and echoed back verbatim; it does not affect the dub.
	Reference *string `json:"reference,omitempty" xml:"reference,omitempty"`
	// BCP-47 language tag of the source media; must be a language the transcription model supports. Any region or script subtag is ignored, since transcription is per-language. Omit to auto-detect.
	SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
	// Dubbing model (`dubbing_v1` or `dubbing_v2`) every language target of this project is dubbed with. Defaults to `dubbing_v2`. Fixed at create time — the source is prepared for this model, so neither the project nor an individual target can change it later.
	ModelId *CreateProjectRequestModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Key terms to bias transcription and translation toward (for example, product or brand names). At most 1,000 terms; each term at most 50 characters and 5 words; the characters `<>{}[]\` are not allowed. Terms are trimmed and deduplicated. Repeat the field once per term, or pass a single JSON array.
	Keyterms []string `json:"keyterms,omitempty" xml:"keyterms,omitempty"`
	// IDs of workspace webhooks to notify as this project progresses — the alternative to polling, and what we recommend. Each receives a `dubbing_project_ready` or `dubbing_project_failed` event for the project, and a `dubbing_language_completed` or `dubbing_language_failed` event for every language under it; `dubbing_language_completed` carries the output download URLs. At most 3 IDs, each already configured in your workspace — see [Webhooks](https://elevenlabs.io/docs/eleven-api/resources/webhooks) for how to create one and verify its signature. Repeat the field once per ID, or pass a single JSON array or comma-separated string. Delivery is best-effort and can repeat, so we recommend handling events idempotently.
	WebhookIds []string `json:"webhook_ids,omitempty" xml:"webhook_ids,omitempty"`
	// Optional shortcut: also create a language target in this BCP-47 language, queued to start once the project is ready — equivalent to creating the project and then creating one language target. Must be one of the [languages the dubbing model supports](https://elevenlabs.io/docs/help-center/product/dubbing/which-languages-are-supported-in-dubbing), and a region-qualified tag must be one of the supported dialects. This is the first language target, so it consumes the project's minimum charge rather than adding to it. Its ID is returned in `language_ids`.
	TargetLanguage *string `json:"target_language,omitempty" xml:"target_language,omitempty"`
	// Enterprise only. Optional JSON transcript to use instead of transcribing the source: a `{"segments": [...]}` document, at most 20,000 segments and 4 MiB. See [Bring your own transcript](https://elevenlabs.io/docs/eleven-api/guides/how-to/dubbing/bring-your-own-transcript) for the segment fields and their constraints. `source_language` is required whenever a transcript is provided. If any segment carries a `translation`, `target_language` is required and every segment must carry one; those translations seed the target created via `target_language`, which then skips machine translation.
	Transcript *string `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetFile() (value *string) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetSourceUrl() (value *string) {
	if b == nil {
		return
	}
	return b.SourceUrl
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetReference() (value *string) {
	if b == nil {
		return
	}
	return b.Reference
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetSourceLanguage() (value *string) {
	if b == nil {
		return
	}
	return b.SourceLanguage
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetModelId() (value *CreateProjectRequestModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetKeyterms() (value []string) {
	if b == nil {
		return
	}
	return b.Keyterms
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetWebhookIds() (value []string) {
	if b == nil {
		return
	}
	return b.WebhookIds
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetTargetLanguage() (value *string) {
	if b == nil {
		return
	}
	return b.TargetLanguage
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetTranscript() (value *string) {
	if b == nil {
		return
	}
	return b.Transcript
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetFile(file *string) {
	b.File = file
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetSourceUrl(sourceUrl *string) {
	b.SourceUrl = sourceUrl
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetReference(reference *string) {
	b.Reference = reference
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetSourceLanguage(sourceLanguage *string) {
	b.SourceLanguage = sourceLanguage
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetModelId(modelId *CreateProjectRequestModelId) {
	b.ModelId = modelId
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetKeyterms(keyterms []string) {
	b.Keyterms = keyterms
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetWebhookIds(webhookIds []string) {
	b.WebhookIds = webhookIds
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetTargetLanguage(targetLanguage *string) {
	b.TargetLanguage = targetLanguage
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) SetTranscript(transcript *string) {
	b.Transcript = transcript
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateDubbingProjectV1DubbingProjectPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateDubbingProjectV1DubbingProjectPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateDubbingProjectV1DubbingProjectPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateDubbingProjectV1DubbingProjectPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// Dubbing model (`dubbing_v1` or `dubbing_v2`) every language target of this project is dubbed with. Defaults to `dubbing_v2`. Fixed at create time — the source is prepared for this model, so neither the project nor an individual target can change it later.
type CreateProjectRequestModelId struct {
	CreateProjectRequestModelIdZero CreateProjectRequestModelIdZero
	String                          string
	typ                             string
}

func (c *CreateProjectRequestModelId) GetCreateProjectRequestModelIdZero() CreateProjectRequestModelIdZero {
	if c == nil {
		return ""
	}
	return c.CreateProjectRequestModelIdZero
}

func (c *CreateProjectRequestModelId) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *CreateProjectRequestModelId) UnmarshalJSON(data []byte) error {
	var valueCreateProjectRequestModelIdZero CreateProjectRequestModelIdZero
	if err := json.Unmarshal(data, &valueCreateProjectRequestModelIdZero); err == nil {
		c.typ = "CreateProjectRequestModelIdZero"
		c.CreateProjectRequestModelIdZero = valueCreateProjectRequestModelIdZero
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateProjectRequestModelId) MarshalJSON() ([]byte, error) {
	if c.typ == "CreateProjectRequestModelIdZero" || c.CreateProjectRequestModelIdZero != "" {
		return json.Marshal(c.CreateProjectRequestModelIdZero)
	}
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateProjectRequestModelIdVisitor interface {
	VisitCreateProjectRequestModelIdZero(CreateProjectRequestModelIdZero) error
	VisitString(string) error
}

func (c *CreateProjectRequestModelId) Accept(visitor CreateProjectRequestModelIdVisitor) error {
	if c.typ == "CreateProjectRequestModelIdZero" || c.CreateProjectRequestModelIdZero != "" {
		return visitor.VisitCreateProjectRequestModelIdZero(c.CreateProjectRequestModelIdZero)
	}
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateProjectRequestModelIdZero string

const (
	CreateProjectRequestModelIdZeroDubbingV1 CreateProjectRequestModelIdZero = "dubbing_v1"
	CreateProjectRequestModelIdZeroDubbingV2 CreateProjectRequestModelIdZero = "dubbing_v2"
)

func NewCreateProjectRequestModelIdZeroFromString(s string) (CreateProjectRequestModelIdZero, error) {
	switch s {
	case "dubbing_v1":
		return CreateProjectRequestModelIdZeroDubbingV1, nil
	case "dubbing_v2":
		return CreateProjectRequestModelIdZeroDubbingV2, nil
	}
	var t CreateProjectRequestModelIdZero
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateProjectRequestModelIdZero) Ptr() *CreateProjectRequestModelIdZero {
	return &e
}

type ListProjectRequestSortDirection string

const (
	ListProjectRequestSortDirectionAscending  ListProjectRequestSortDirection = "ASCENDING"
	ListProjectRequestSortDirectionDescending ListProjectRequestSortDirection = "DESCENDING"
)

func NewListProjectRequestSortDirectionFromString(s string) (ListProjectRequestSortDirection, error) {
	switch s {
	case "ASCENDING":
		return ListProjectRequestSortDirectionAscending, nil
	case "DESCENDING":
		return ListProjectRequestSortDirectionDescending, nil
	}
	var t ListProjectRequestSortDirection
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListProjectRequestSortDirection) Ptr() *ListProjectRequestSortDirection {
	return &e
}

var (
	listProjectRequestFieldCursor        = big.NewInt(1 << 0)
	listProjectRequestFieldPageSize      = big.NewInt(1 << 1)
	listProjectRequestFieldStatus        = big.NewInt(1 << 2)
	listProjectRequestFieldSortDirection = big.NewInt(1 << 3)
)

type ListProjectRequest struct {
	Cursor        *string                          `json:"-" url:"cursor,omitempty"`
	PageSize      *int                             `json:"-" url:"page_size,omitempty"`
	Status        *string                          `json:"-" url:"status,omitempty"`
	SortDirection *ListProjectRequestSortDirection `json:"-" url:"sort_direction,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListProjectRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProjectRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listProjectRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProjectRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listProjectRequestFieldPageSize)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProjectRequest) SetStatus(status *string) {
	l.Status = status
	l.require(listProjectRequestFieldStatus)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProjectRequest) SetSortDirection(sortDirection *ListProjectRequestSortDirection) {
	l.SortDirection = sortDirection
	l.require(listProjectRequestFieldSortDirection)
}
