package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"noteIt/api/models"
	"strconv"
)

/**Mock database*/
var notes = []models.Note{
	{1, "Title 1", "Description 1"},
	{2, "Title 2", "Description 2"},
	{3, "Title 3", "Description 3"},
}

func main() {
	fmt.Println("Starting server...")

	router := mux.NewRouter()

	router.HandleFunc("/", Home)
	router.HandleFunc("/notes", GetNotes).Methods("GET")
	router.HandleFunc("/note", GetNote).Methods("POST")
	router.HandleFunc("/create/note", CreateNote).Methods("POST")
	router.HandleFunc("/update/note", UpdateNote).Methods("PATCH")
	router.HandleFunc("/delete/notes", DeleteAllNotes).Methods("DELETE")
	router.HandleFunc("/delete/note", DeleteNote).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Home(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprint(wr, "Home endpoint")
}

func GetNotes(wr http.ResponseWriter, r *http.Request) {
	value, err := json.Marshal(&notes)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	wr.WriteHeader(http.StatusOK)
	fmt.Fprint(wr, string(value))
}

func GetNote(wr http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var note models.Note

	for _, v := range notes {
		if strconv.Itoa(v.ID) == id {
			note = v
		}
	}

	value, err := json.Marshal(&note)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	wr.WriteHeader(http.StatusOK)
	fmt.Fprint(wr, string(value))
}

func CreateNote(wr http.ResponseWriter, r *http.Request) {
	var newNote models.Note

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = json.Unmarshal(body, &newNote)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	notes = append(notes, newNote)
	notes, err := json.Marshal(&notes)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(notes)
}

func UpdateNote(wr http.ResponseWriter, r *http.Request) {
	var newNote models.Note

	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = json.Unmarshal(body, &newNote)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	for i, v := range notes {
		if strconv.Itoa(v.ID) == id {
			v = models.Note{
				ID:          newNote.ID,
				Title:       newNote.Title,
				Description: newNote.Description,
			}

			notes = append(notes[:i], v)
		}
	}

	wr.WriteHeader(http.StatusCreated)
	fmt.Fprint(wr, notes)
}

func DeleteAllNotes(wr http.ResponseWriter, r *http.Request) {
	notes = notes[:0]

	body, err := json.Marshal(notes)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Fprint(wr, string(body))
}

func DeleteNote(wr http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for i, v := range notes {
		if strconv.Itoa(v.ID) == id {
			notes = append(notes[:i], notes[i+1:]...)
		}
	}

	wr.WriteHeader(http.StatusNoContent)

	resp, err := json.Marshal(notes)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Fprint(wr, string(resp))
}
