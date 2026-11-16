package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A reference image guiding a Veo generation, with its role.
type VeoImageReference struct {
	Image *ImageReference `json:"image" xml:"image" required:"true"`
	// How the model uses the image: `subject` places its subject or scene elements into the video; `style` transfers its visual style.
	Role            VeoImageReferenceRole `json:"role" xml:"role" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VeoImageReference) GetImage() (value *ImageReference) {
	if v == nil {
		return
	}
	return v.Image
}

func (v *VeoImageReference) GetRole() (value VeoImageReferenceRole) {
	if v == nil {
		return
	}
	return v.Role
}

func (v *VeoImageReference) SetImage(image *ImageReference) {
	v.Image = image
}

func (v *VeoImageReference) SetRole(role VeoImageReferenceRole) {
	v.Role = role
}

func (v *VeoImageReference) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VeoImageReference) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VeoImageReference to string"
	}
	return string(jsonData)
}

func (v *VeoImageReference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VeoImageReference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VeoImageReference(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VeoImageReference) MarshalJSON() ([]byte, error) {
	type embed VeoImageReference
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
