package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Device information.
//
// Spec: https://schema.ocsf.io/1.6.0/objects/device
type DeviceModel struct {
	// IP address
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// Device hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// Device type ID (99 = Unknown)
	TypeId          *int `json:"type_id,omitempty" xml:"type_id,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DeviceModel) GetIp() (value *string) {
	if d == nil {
		return
	}
	return d.Ip
}

func (d *DeviceModel) GetHostname() (value *string) {
	if d == nil {
		return
	}
	return d.Hostname
}

func (d *DeviceModel) GetTypeId() (value *int) {
	if d == nil {
		return
	}
	return d.TypeId
}

func (d *DeviceModel) SetIp(ip *string) {
	d.Ip = ip
}

func (d *DeviceModel) SetHostname(hostname *string) {
	d.Hostname = hostname
}

func (d *DeviceModel) SetTypeId(typeId *int) {
	d.TypeId = typeId
}

func (d *DeviceModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeviceModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeviceModel to string"
	}
	return string(jsonData)
}

func (d *DeviceModel) UnmarshalJSON(data []byte) error {
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

func (d *DeviceModel) MarshalJSON() ([]byte, error) {
	type embed DeviceModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
