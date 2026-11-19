package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// OCSF Actor object - describes the entity that performed the action.
//
// Spec: https://schema.ocsf.io/1.6.0/objects/actor
type ActorModel struct {
	// OCSF User object.
	//
	// Spec: https://schema.ocsf.io/1.6.0/objects/user
	User *UserModel `json:"user" xml:"user" required:"true"`
	// Client application or service name
	AppName *string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// Client application unique identifier
	AppUid *string `json:"app_uid,omitempty" xml:"app_uid,omitempty"`
	// Session information
	Session         map[string]any `json:"session,omitempty"`
	extraProperties map[string]interface{}
}

func (a *ActorModel) GetUser() (value *UserModel) {
	if a == nil {
		return
	}
	return a.User
}

func (a *ActorModel) GetAppName() (value *string) {
	if a == nil {
		return
	}
	return a.AppName
}

func (a *ActorModel) GetAppUid() (value *string) {
	if a == nil {
		return
	}
	return a.AppUid
}

func (a *ActorModel) GetSession() (value map[string]any) {
	if a == nil {
		return
	}
	return a.Session
}

func (a *ActorModel) SetUser(user *UserModel) {
	a.User = user
}

func (a *ActorModel) SetAppName(appName *string) {
	a.AppName = appName
}

func (a *ActorModel) SetAppUid(appUid *string) {
	a.AppUid = appUid
}

func (a *ActorModel) SetSession(session map[string]any) {
	a.Session = session
}

func (a *ActorModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ActorModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ActorModel to string"
	}
	return string(jsonData)
}

func (a *ActorModel) UnmarshalJSON(data []byte) error {
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

func (a *ActorModel) MarshalJSON() ([]byte, error) {
	type embed ActorModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
