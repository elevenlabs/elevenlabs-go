package types

import (
	"testing"
)

func TestDeviceModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "ip": "value",
  "hostname": "value",
  "type_id": 123
}`
	assertRoundTrip(t, wire, new(DeviceModel))
}
