package types

import (
	"testing"
)

func TestMusicExploreSongSourceContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "music_explore_song",
  "music_explore_song_id": "value",
  "title": "value",
  "description": "value",
  "bpm": 123,
  "vocals": "value",
  "lyrics": "value"
}`
	assertRoundTrip(t, wire, new(MusicExploreSongSourceContext))
}
