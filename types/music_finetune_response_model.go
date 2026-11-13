package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type MusicFinetuneResponseModel struct {
	// Unique identifier of the finetune.
	Id string `json:"id" xml:"id" required:"true"`
	// Name of the finetune.
	Name string `json:"name" xml:"name" required:"true"`
	// Tags associated with the finetune.
	Tags []string `json:"tags" xml:"tags" required:"true"`
	// Primary musical genre of the finetune.
	PrimaryGenre *string `json:"primary_genre,omitempty" xml:"primary_genre,omitempty"`
	// The base music model the finetune was trained on.
	ModelId string `json:"model_id" xml:"model_id" required:"true"`
	// When the finetune was created (UTC).
	CreatedAt  time.Time           `json:"created_at" xml:"created_at" required:"true"`
	Visibility FinetuneVisibility  `json:"visibility" xml:"visibility" required:"true"`
	CreatedBy  FinetuneCreatedBy   `json:"created_by" xml:"created_by" required:"true"`
	Status     MusicFinetuneStatus `json:"status" xml:"status" required:"true"`
	// Training progress from 0.0 to 1.0.
	TrainingProgress float64 `json:"training_progress" xml:"training_progress" required:"true"`
	// Reason the finetune failed or was blocked, if applicable.
	FailureReason   *MusicFinetuneFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MusicFinetuneResponseModel) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MusicFinetuneResponseModel) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *MusicFinetuneResponseModel) GetTags() (value []string) {
	if m == nil {
		return
	}
	return m.Tags
}

func (m *MusicFinetuneResponseModel) GetPrimaryGenre() (value *string) {
	if m == nil {
		return
	}
	return m.PrimaryGenre
}

func (m *MusicFinetuneResponseModel) GetModelId() (value string) {
	if m == nil {
		return
	}
	return m.ModelId
}

func (m *MusicFinetuneResponseModel) GetCreatedAt() (value time.Time) {
	if m == nil {
		return
	}
	return m.CreatedAt
}

func (m *MusicFinetuneResponseModel) GetVisibility() (value FinetuneVisibility) {
	if m == nil {
		return
	}
	return m.Visibility
}

func (m *MusicFinetuneResponseModel) GetCreatedBy() (value FinetuneCreatedBy) {
	if m == nil {
		return
	}
	return m.CreatedBy
}

func (m *MusicFinetuneResponseModel) GetStatus() (value MusicFinetuneStatus) {
	if m == nil {
		return
	}
	return m.Status
}

func (m *MusicFinetuneResponseModel) GetTrainingProgress() (value float64) {
	if m == nil {
		return
	}
	return m.TrainingProgress
}

func (m *MusicFinetuneResponseModel) GetFailureReason() (value *MusicFinetuneFailureReason) {
	if m == nil {
		return
	}
	return m.FailureReason
}

func (m *MusicFinetuneResponseModel) SetId(id string) {
	m.Id = id
}

func (m *MusicFinetuneResponseModel) SetName(name string) {
	m.Name = name
}

func (m *MusicFinetuneResponseModel) SetTags(tags []string) {
	m.Tags = tags
}

func (m *MusicFinetuneResponseModel) SetPrimaryGenre(primaryGenre *string) {
	m.PrimaryGenre = primaryGenre
}

func (m *MusicFinetuneResponseModel) SetModelId(modelId string) {
	m.ModelId = modelId
}

func (m *MusicFinetuneResponseModel) SetCreatedAt(createdAt time.Time) {
	m.CreatedAt = createdAt
}

func (m *MusicFinetuneResponseModel) SetVisibility(visibility FinetuneVisibility) {
	m.Visibility = visibility
}

func (m *MusicFinetuneResponseModel) SetCreatedBy(createdBy FinetuneCreatedBy) {
	m.CreatedBy = createdBy
}

func (m *MusicFinetuneResponseModel) SetStatus(status MusicFinetuneStatus) {
	m.Status = status
}

func (m *MusicFinetuneResponseModel) SetTrainingProgress(trainingProgress float64) {
	m.TrainingProgress = trainingProgress
}

func (m *MusicFinetuneResponseModel) SetFailureReason(failureReason *MusicFinetuneFailureReason) {
	m.FailureReason = failureReason
}

func (m *MusicFinetuneResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MusicFinetuneResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MusicFinetuneResponseModel to string"
	}
	return string(jsonData)
}

func (m *MusicFinetuneResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MusicFinetuneResponseModel) MarshalJSON() ([]byte, error) {
	type embed MusicFinetuneResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
