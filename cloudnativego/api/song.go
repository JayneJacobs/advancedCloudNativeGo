package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Song type with ID Title Author and url
type Song struct {
	//define song
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	MusicURL string `json:"musicurl"`
}

var songs = map[string]Song{
	"2": Song{ID: "2", Title: "Title1", Author: "Jayne Jacobs", MusicURL: "http://jaynejacobs.com"},
	"3": Song{ID: "3", Title: "Title2", Author: "Jayne Jacobs", MusicURL: "http://jaynejacobs.com/2"},
}

// ToJSON Is here
func (s Song) ToJSON() []byte {

	ToJSON, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON is here
func FromJSON(data []byte) Song {
	song := Song{}
	err := json.Unmarshal(data, &song)
	if err != nil {
		panic(err)
	}
	return song
}

// AllSongs is here
func AllSongs() []Song {
	values := make([]Song, len(songs))
	idx := 0
	for _, song := range songs {
		values[idx] = song
		idx++
	}
	return values
}

// SongsHandleFunc here
func SongsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		songs := AllSongs()
		writeJSON(w, songs)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		song := FromJSON(body)
		id, created := CreateSong(song)
		if created {
			w.Header().Add("Location", "api/songs/"+id)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsuported"))

	}
}

// SongHandleFunc this is...
func SongHandleFunc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("api/songs/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		song, found := GetSong(id)
		if found {
			writeJSON(w, song)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		song := FromJSON(body)
		exists := UpdateSong(id, song)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteSong(id)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Request Method."))
	}

}

// GetSong this is...
func GetSong(id string) (Song, bool) {
	song, found := songs[id]
	return song, found

}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json;charset=utf8")
	w.Write(b)
}

// UpdateSong ....
func UpdateSong(id string, song Song) bool {
	_, exists := songs[id]
	if exists {
		songs[id] = song
	}
	return exists
}

// CreateSong ....
func CreateSong(song Song) (string, bool) {
	_, exists := songs[song.ID]
	if exists {
		return "", false
	}
	songs[song.ID] = song
	return song.ID, true
}

// DeleteSong ...
func DeleteSong(id string) {
	delete(songs, id)
}
