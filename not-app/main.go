package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	notes  = make(map[int]Note)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/notes", notesHandler)
	http.HandleFunc("/notes/", noteHandler)

	http.ListenAndServe(":8080", nil)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNotes(w)
	case http.MethodPost:
		createNote(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/notes/"):]

	switch r.Method {
	case http.MethodGet:
		getNote(w, id)
	case http.MethodPut:
		updateNote(w, r, id)
	case http.MethodDelete:
		deleteNote(w, id)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func getNotes(w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()

	var notesList []Note
	for _, note := range notes {
		notesList = append(notesList, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notesList)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note.ID = nextID
	nextID++
	notes[note.ID] = note

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func getNote(w http.ResponseWriter, id string) {
	mu.Lock()
	defer mu.Unlock()

	for _, note := range notes {
		if string(note.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

func updateNote(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	defer mu.Unlock()

	var updatedNote Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, note := range notes {
		if string(note.ID) == id {
			updatedNote.ID = note.ID
			notes[i] = updatedNote
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedNote)
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}

func deleteNote(w http.ResponseWriter, id string) {
	mu.Lock()
	defer mu.Unlock()

	for i, note := range notes {
		if string(note.ID) == id {
			delete(notes, note.ID)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Note not found", http.StatusNotFound)
}