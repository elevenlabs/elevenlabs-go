package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type DubbingLanguageResponse struct {
	// Unique identifier of the language target.
	LanguageId string `json:"language_id" xml:"language_id" required:"true"`
	// Identifier of the parent dubbing project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// BCP-47 language tag this target is dubbed into.
	TargetLanguage string `json:"target_language" xml:"target_language" required:"true"`
	// Lifecycle status: `queued` (waiting on the project to be ready, or on a worker), `processing` while it is being dubbed, `completed` once its output is available, `stale` when the transcript changed after the output was produced, or `failed`.
	Status DubbingLanguageResponseStatus `json:"status" xml:"status" required:"true"`
	// Dubbing model this target is dubbed with, inherited from the project and not selectable per language.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Voice settings applied to every speaker in this language, or null if the defaults apply.
	VoiceSettings VoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	// Signed output URLs; null until the target has produced an output (present once `completed`, and kept while `stale` — compare `output_revision` against `revision` to tell whether the output is up to date).
	Outputs *DubbingLanguageOutputs `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Monotonic counter incremented whenever this target's transcript changes (a source edit affecting it, or an edit to its translation).
	Revision int `json:"revision" xml:"revision" required:"true"`
	// The `revision` the current dubbed output was generated from; equal to `revision` when up to date, and lower when `stale`. This is null until a generation has completed.
	OutputRevision *int `json:"output_revision,omitempty" xml:"output_revision,omitempty"`
	// Why this language failed; null unless `status` is `failed`, and also null for the few languages that failed before failure reporting was introduced. A code of `project_failed` means the parent project failed, so read the project for the underlying cause.
	Error *DubbingErrorResponse `json:"error,omitempty" xml:"error,omitempty"`
	// Non-fatal conditions raised while dubbing this language, empty when there are none. Reflects the latest generation. Conditions raised while preparing the source are reported on the project instead.
	Warnings []*VoicesNotPermittedWarning `json:"warnings,omitempty" xml:"warnings,omitempty"`
	// When the language target was created.
	CreatedAt time.Time `json:"created_at" xml:"created_at" required:"true"`
	// When the language target was last updated.
	UpdatedAt       time.Time `json:"updated_at" xml:"updated_at" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingLanguageResponse) GetLanguageId() (value string) {
	if d == nil {
		return
	}
	return d.LanguageId
}

func (d *DubbingLanguageResponse) GetProjectId() (value string) {
	if d == nil {
		return
	}
	return d.ProjectId
}

func (d *DubbingLanguageResponse) GetTargetLanguage() (value string) {
	if d == nil {
		return
	}
	return d.TargetLanguage
}

func (d *DubbingLanguageResponse) GetStatus() (value DubbingLanguageResponseStatus) {
	if d == nil {
		return
	}
	return d.Status
}

func (d *DubbingLanguageResponse) GetModelId() (value *string) {
	if d == nil {
		return
	}
	return d.ModelId
}

func (d *DubbingLanguageResponse) GetVoiceSettings() (value VoiceSettings) {
	if d == nil {
		return
	}
	return d.VoiceSettings
}

func (d *DubbingLanguageResponse) GetOutputs() (value *DubbingLanguageOutputs) {
	if d == nil {
		return
	}
	return d.Outputs
}

func (d *DubbingLanguageResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingLanguageResponse) GetOutputRevision() (value *int) {
	if d == nil {
		return
	}
	return d.OutputRevision
}

func (d *DubbingLanguageResponse) GetError() (value *DubbingErrorResponse) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingLanguageResponse) GetWarnings() (value []*VoicesNotPermittedWarning) {
	if d == nil {
		return
	}
	return d.Warnings
}

func (d *DubbingLanguageResponse) GetCreatedAt() (value time.Time) {
	if d == nil {
		return
	}
	return d.CreatedAt
}

func (d *DubbingLanguageResponse) GetUpdatedAt() (value time.Time) {
	if d == nil {
		return
	}
	return d.UpdatedAt
}

func (d *DubbingLanguageResponse) SetLanguageId(languageId string) {
	d.LanguageId = languageId
}

func (d *DubbingLanguageResponse) SetProjectId(projectId string) {
	d.ProjectId = projectId
}

func (d *DubbingLanguageResponse) SetTargetLanguage(targetLanguage string) {
	d.TargetLanguage = targetLanguage
}

func (d *DubbingLanguageResponse) SetStatus(status DubbingLanguageResponseStatus) {
	d.Status = status
}

func (d *DubbingLanguageResponse) SetModelId(modelId *string) {
	d.ModelId = modelId
}

func (d *DubbingLanguageResponse) SetVoiceSettings(voiceSettings VoiceSettings) {
	d.VoiceSettings = voiceSettings
}

func (d *DubbingLanguageResponse) SetOutputs(outputs *DubbingLanguageOutputs) {
	d.Outputs = outputs
}

func (d *DubbingLanguageResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingLanguageResponse) SetOutputRevision(outputRevision *int) {
	d.OutputRevision = outputRevision
}

func (d *DubbingLanguageResponse) SetError(error *DubbingErrorResponse) {
	d.Error = error
}

func (d *DubbingLanguageResponse) SetWarnings(warnings []*VoicesNotPermittedWarning) {
	d.Warnings = warnings
}

func (d *DubbingLanguageResponse) SetCreatedAt(createdAt time.Time) {
	d.CreatedAt = createdAt
}

func (d *DubbingLanguageResponse) SetUpdatedAt(updatedAt time.Time) {
	d.UpdatedAt = updatedAt
}

func (d *DubbingLanguageResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingLanguageResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingLanguageResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingLanguageResponse) UnmarshalJSON(data []byte) error {
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

func (d *DubbingLanguageResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingLanguageResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
