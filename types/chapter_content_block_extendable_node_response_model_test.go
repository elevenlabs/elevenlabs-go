package types

import (
	"testing"
)

func TestChapterContentBlockExtendableNodeResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "_other"
}`
	assertRoundTrip(t, wire, new(ChapterContentBlockExtendableNodeResponseModel))
}
