package types

import (
	"testing"
)

func TestListHolidaysParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_holidays"
}`
	assertRoundTrip(t, wire, new(ListHolidaysParams))
}
