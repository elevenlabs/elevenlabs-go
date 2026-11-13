package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type DubbingProjectResponse struct {
	// Unique identifier of the dubbing project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// Lifecycle status of the project: `queued` before the source is picked up, `preparing` while it is transcribed, `ready` once transcription is done and language targets can start, or `failed`. A project is never reported as `processing` — that value belongs to language targets.
	Status DubbingProjectResponseStatus `json:"status" xml:"status" required:"true"`
	// The free-form string you supplied as `reference` when creating the project, or null if you supplied none.
	Reference *string `json:"reference,omitempty" xml:"reference,omitempty"`
	// BCP-47 language tag of the source media (null if auto-detected).
	SourceLanguage *string `json:"source_language,omitempty" xml:"source_language,omitempty"`
	// Dubbing model every language target of this project is dubbed with. Fixed at create time and not selectable per language.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Source media metadata, populated once the source has been fetched and decoded (shortly after create, before the project is `ready`); null until then.
	Media *DubbingSourceMediaInfo `json:"media,omitempty" xml:"media,omitempty"`
	// Identifiers of the language targets under this project. Populated when a single project is fetched, and on create when `target_language` creates one. Always empty in list responses — list the project's language targets instead.
	LanguageIds []string `json:"language_ids,omitempty" xml:"language_ids,omitempty"`
	// IDs of the workspace webhooks notified as this project and its languages reach `ready`, `completed`, or `failed`.
	WebhookIds []string `json:"webhook_ids,omitempty" xml:"webhook_ids,omitempty"`
	// Monotonic counter incremented whenever the source transcript is edited (segment add/edit/delete).
	Revision int `json:"revision" xml:"revision" required:"true"`
	// Why the project failed; null unless `status` is `failed`. Also null for the few projects that failed before failure reporting was introduced.
	Error *DubbingErrorResponse `json:"error,omitempty" xml:"error,omitempty"`
	// Non-fatal conditions raised while preparing the source, empty when there are none. Reflects the latest preparation. Conditions raised while dubbing a particular language are reported on that language instead.
	Warnings []*VoicesNotPermittedWarning `json:"warnings,omitempty" xml:"warnings,omitempty"`
	// When the project was created.
	CreatedAt time.Time `json:"created_at" xml:"created_at" required:"true"`
	// When the project was last updated.
	UpdatedAt       time.Time `json:"updated_at" xml:"updated_at" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingProjectResponse) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingProjectResponse) GetStatus() (value DubbingProjectResponseStatus) {
	if d == nil {
		return
	}
	return d.Status
}

func (d *DubbingProjectResponse) GetReference() (value *string) {
	if d == nil {
		return
	}
	return d.Reference
}

func (d *DubbingProjectResponse) GetSourceLanguage() (value *string) {
	if d == nil {
		return
	}
	return d.SourceLanguage
}

func (d *DubbingProjectResponse) GetModelId() (value *string) {
	if d == nil {
		return
	}
	return d.ModelId
}

func (d *DubbingProjectResponse) GetMedia() (value *DubbingSourceMediaInfo) {
	if d == nil {
		return
	}
	return d.Media
}

func (d *DubbingProjectResponse) GetLanguageIds() (value []string) {
	if d == nil {
		return
	}
	return d.LanguageIds
}

func (d *DubbingProjectResponse) GetWebhookIds() (value []string) {
	if d == nil {
		return
	}
	return d.WebhookIds
}

func (d *DubbingProjectResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingProjectResponse) GetError() (value *DubbingErrorResponse) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingProjectResponse) GetWarnings() (value []*VoicesNotPermittedWarning) {
	if d == nil {
		return
	}
	return d.Warnings
}

func (d *DubbingProjectResponse) GetCreatedAt() (value time.Time) {
	if d == nil {
		return
	}
	return d.CreatedAt
}

func (d *DubbingProjectResponse) GetUpdatedAt() (value time.Time) {
	if d == nil {
		return
	}
	return d.UpdatedAt
}

func (d *DubbingProjectResponse) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingProjectResponse) SetStatus(status DubbingProjectResponseStatus) {
	d.Status = status
}

func (d *DubbingProjectResponse) SetReference(reference *string) {
	d.Reference = reference
}

func (d *DubbingProjectResponse) SetSourceLanguage(sourceLanguage *string) {
	d.SourceLanguage = sourceLanguage
}

func (d *DubbingProjectResponse) SetModelId(modelId *string) {
	d.ModelId = modelId
}

func (d *DubbingProjectResponse) SetMedia(media *DubbingSourceMediaInfo) {
	d.Media = media
}

func (d *DubbingProjectResponse) SetLanguageIds(languageIds []string) {
	d.LanguageIds = languageIds
}

func (d *DubbingProjectResponse) SetWebhookIds(webhookIds []string) {
	d.WebhookIds = webhookIds
}

func (d *DubbingProjectResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingProjectResponse) SetError(error *DubbingErrorResponse) {
	d.Error = error
}

func (d *DubbingProjectResponse) SetWarnings(warnings []*VoicesNotPermittedWarning) {
	d.Warnings = warnings
}

func (d *DubbingProjectResponse) SetCreatedAt(createdAt time.Time) {
	d.CreatedAt = createdAt
}

func (d *DubbingProjectResponse) SetUpdatedAt(updatedAt time.Time) {
	d.UpdatedAt = updatedAt
}

func (d *DubbingProjectResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingProjectResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingProjectResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingProjectResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingProjectResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingProjectResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
