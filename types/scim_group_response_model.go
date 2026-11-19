package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScimGroupResponseModel struct {
	ScimExternalId  *string   `json:"scim_external_id" xml:"scim_external_id" required:"true"`
	DisplayName     string    `json:"display_name" xml:"display_name" required:"true"`
	CreatedAtUnix   *int      `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	UpdatedAtUnix   *int      `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	SeatType        *SeatType `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	extraProperties map[string]interface{}
}

func (s *ScimGroupResponseModel) GetScimExternalId() (value *string) {
	if s == nil {
		return
	}
	return s.ScimExternalId
}

func (s *ScimGroupResponseModel) GetDisplayName() (value string) {
	if s == nil {
		return
	}
	return s.DisplayName
}

func (s *ScimGroupResponseModel) GetCreatedAtUnix() (value *int) {
	if s == nil {
		return
	}
	return s.CreatedAtUnix
}

func (s *ScimGroupResponseModel) GetUpdatedAtUnix() (value *int) {
	if s == nil {
		return
	}
	return s.UpdatedAtUnix
}

func (s *ScimGroupResponseModel) GetSeatType() (value *SeatType) {
	if s == nil {
		return
	}
	return s.SeatType
}

func (s *ScimGroupResponseModel) SetScimExternalId(scimExternalId *string) {
	s.ScimExternalId = scimExternalId
}

func (s *ScimGroupResponseModel) SetDisplayName(displayName string) {
	s.DisplayName = displayName
}

func (s *ScimGroupResponseModel) SetCreatedAtUnix(createdAtUnix *int) {
	s.CreatedAtUnix = createdAtUnix
}

func (s *ScimGroupResponseModel) SetUpdatedAtUnix(updatedAtUnix *int) {
	s.UpdatedAtUnix = updatedAtUnix
}

func (s *ScimGroupResponseModel) SetSeatType(seatType *SeatType) {
	s.SeatType = seatType
}

func (s *ScimGroupResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScimGroupResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScimGroupResponseModel to string"
	}
	return string(jsonData)
}

func (s *ScimGroupResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScimGroupResponseModel) MarshalJSON() ([]byte, error) {
	type embed ScimGroupResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
