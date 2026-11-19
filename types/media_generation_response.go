package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type MediaGenerationResponse struct {
	Status     string
	Completed  *MediaGenerationCompletedResponse
	Failed     *MediaGenerationFailedResponse
	Generating *MediaGenerationInProgressResponse
	Pending    *MediaGenerationInProgressResponse
}

func (m *MediaGenerationResponse) GetCompleted() *MediaGenerationCompletedResponse {
	if m == nil {
		return nil
	}
	if m.Completed == nil {
		return nil
	}
	return m.Completed
}

func (m *MediaGenerationResponse) GetFailed() *MediaGenerationFailedResponse {
	if m == nil {
		return nil
	}
	if m.Failed == nil {
		return nil
	}
	return m.Failed
}

func (m *MediaGenerationResponse) GetGenerating() *MediaGenerationInProgressResponse {
	if m == nil {
		return nil
	}
	if m.Generating == nil {
		return nil
	}
	return m.Generating
}

func (m *MediaGenerationResponse) GetPending() *MediaGenerationInProgressResponse {
	if m == nil {
		return nil
	}
	if m.Pending == nil {
		return nil
	}
	return m.Pending
}

func (m *MediaGenerationResponse) GetStatus() string {
	if m == nil {
		return ""
	}
	return m.Status
}

func (m *MediaGenerationResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", m)
	}
	switch unmarshaler.Status {
	case "completed":
		value := new(MediaGenerationCompletedResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Completed = value
	case "failed":
		value := new(MediaGenerationFailedResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Failed = value
	case "generating":
		value := new(MediaGenerationInProgressResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Generating = value
	case "pending":
		value := new(MediaGenerationInProgressResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Pending = value
	}
	return nil
}

func (m MediaGenerationResponse) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Completed != nil {
		return internal.MarshalJSONWithDiscriminant(m.Completed, "status", "completed")
	}
	if m.Failed != nil {
		return internal.MarshalJSONWithDiscriminant(m.Failed, "status", "failed")
	}
	if m.Generating != nil {
		return internal.MarshalJSONWithDiscriminant(m.Generating, "status", "generating")
	}
	if m.Pending != nil {
		return internal.MarshalJSONWithDiscriminant(m.Pending, "status", "pending")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type MediaGenerationResponseVisitor interface {
	VisitCompleted(*MediaGenerationCompletedResponse) error
	VisitFailed(*MediaGenerationFailedResponse) error
	VisitGenerating(*MediaGenerationInProgressResponse) error
	VisitPending(*MediaGenerationInProgressResponse) error
}

func (m *MediaGenerationResponse) Accept(visitor MediaGenerationResponseVisitor) error {
	if m.Completed != nil {
		return visitor.VisitCompleted(m.Completed)
	}
	if m.Failed != nil {
		return visitor.VisitFailed(m.Failed)
	}
	if m.Generating != nil {
		return visitor.VisitGenerating(m.Generating)
	}
	if m.Pending != nil {
		return visitor.VisitPending(m.Pending)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *MediaGenerationResponse) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Completed != nil {
		fields = append(fields, "completed")
	}
	if m.Failed != nil {
		fields = append(fields, "failed")
	}
	if m.Generating != nil {
		fields = append(fields, "generating")
	}
	if m.Pending != nil {
		fields = append(fields, "pending")
	}
	if len(fields) == 0 {
		if m.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Status)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Status != "" {
		field := fields[0]
		if m.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Status,
				m,
			)
		}
	}
	return nil
}
