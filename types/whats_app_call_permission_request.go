package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppCallPermissionRequest struct {
	State             WhatsAppCallPermissionRequestState `json:"state" xml:"state" required:"true"`
	RequestedAtUnix   *int                               `json:"requested_at_unix,omitempty" xml:"requested_at_unix,omitempty"`
	RespondedAtUnix   *int                               `json:"responded_at_unix,omitempty" xml:"responded_at_unix,omitempty"`
	Permanent         *bool                              `json:"permanent,omitempty" xml:"permanent,omitempty"`
	GrantedByUserCall *bool                              `json:"granted_by_user_call,omitempty" xml:"granted_by_user_call,omitempty"`
	extraProperties   map[string]interface{}
}

func (w *WhatsAppCallPermissionRequest) GetState() (value WhatsAppCallPermissionRequestState) {
	if w == nil {
		return
	}
	return w.State
}

func (w *WhatsAppCallPermissionRequest) GetRequestedAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.RequestedAtUnix
}

func (w *WhatsAppCallPermissionRequest) GetRespondedAtUnix() (value *int) {
	if w == nil {
		return
	}
	return w.RespondedAtUnix
}

func (w *WhatsAppCallPermissionRequest) GetPermanent() (value *bool) {
	if w == nil {
		return
	}
	return w.Permanent
}

func (w *WhatsAppCallPermissionRequest) GetGrantedByUserCall() (value *bool) {
	if w == nil {
		return
	}
	return w.GrantedByUserCall
}

func (w *WhatsAppCallPermissionRequest) SetState(state WhatsAppCallPermissionRequestState) {
	w.State = state
}

func (w *WhatsAppCallPermissionRequest) SetRequestedAtUnix(requestedAtUnix *int) {
	w.RequestedAtUnix = requestedAtUnix
}

func (w *WhatsAppCallPermissionRequest) SetRespondedAtUnix(respondedAtUnix *int) {
	w.RespondedAtUnix = respondedAtUnix
}

func (w *WhatsAppCallPermissionRequest) SetPermanent(permanent *bool) {
	w.Permanent = permanent
}

func (w *WhatsAppCallPermissionRequest) SetGrantedByUserCall(grantedByUserCall *bool) {
	w.GrantedByUserCall = grantedByUserCall
}

func (w *WhatsAppCallPermissionRequest) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppCallPermissionRequest) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppCallPermissionRequest to string"
	}
	return string(jsonData)
}

func (w *WhatsAppCallPermissionRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppCallPermissionRequest) MarshalJSON() ([]byte, error) {
	type embed WhatsAppCallPermissionRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
