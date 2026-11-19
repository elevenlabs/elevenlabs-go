package types

import (
	"testing"
)

func TestWorkspaceAuditLogEntryResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "metadata": {
    "key": {}
  },
  "time": 123,
  "activity_id": 123,
  "activity_name": "value",
  "category_name": "value",
  "category_uid": 123,
  "class_name": "value",
  "class_uid": 123,
  "severity_id": 123,
  "status_id": 123,
  "actor": {
    "user": {
      "name": "value",
      "uid": "value",
      "type_id": 123,
      "type": "value",
      "email_addr": "value",
      "full_name": "value",
      "domain": "value"
    },
    "app_name": "value",
    "app_uid": "value",
    "session": {
      "key": {}
    }
  },
  "device": {
    "ip": "value",
    "hostname": "value",
    "type_id": 123
  },
  "http_request": {
    "http_method": "value",
    "url": {
      "url_string": "value",
      "scheme": "value",
      "hostname": "value",
      "port": 123,
      "path": "value",
      "query_string": "value"
    },
    "user_agent": "value",
    "x_forwarded_for": [
      "value"
    ]
  },
  "message": "value",
  "unmapped": {
    "key": {}
  },
  "id": "value",
  "time_dt": "value",
  "type_uid": 123,
  "type_name": "value"
}`
	assertRoundTrip(t, wire, new(WorkspaceAuditLogEntryResponse))
}
