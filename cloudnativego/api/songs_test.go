package api

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestSongToJSON(t *testing.T) {
	song := Song{ID: "1", Title: "Girls Just Want to Have Fun", Author: "Cindy Lauper", MusicURL: "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196"}
	json := song.ToJSON()
	assert.Equal(t, `{"id": "1", "title": "Girls Just Want to Have Fun", "author": "Cindy Lauper", "musicurl": "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196"}`, json, "Song JSON marshalling wrong")
}

// TestSongFromJSON
func TestSongFromJSON(t *testing.T) {

	json := []byte(`{"id": "1", "title": "Girls Just Want to Have Fun", "author": "Cindy Lauper", "musicurl": "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196"}`)
	song := FromJSON(json)
	assert.Equal(t, Song{ID: "1", Title: "Girls Just Want to Have Fun", Author: "Cindy Lauper", MusicURL: "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196"}, song, "Song JSON-unmarshalling wrong")
}
