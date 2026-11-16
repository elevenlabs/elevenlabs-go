package types

import (
	"testing"
)

func TestCreateCrawlJobResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "type": "discovery",
  "root_folder_id": "value",
  "status": "value",
  "created_at": 123,
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(CreateCrawlJobResponseModel))
}
