package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Hello world!")

	http.HandleFunc("/home", Home)
	http.HandleFunc("/notes", GetNotes)
	http.HandleFunc("/note", GetNote)
	http.HandleFunc("/create/note", CreateNote)
	http.HandleFunc("/update/note", UpdateNote)
	http.HandleFunc("/delete/notes", DeleteAllNotes)
	http.HandleFunc("/delete/note", DeleteNote)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Home(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(wr, "Home page")
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
	fmt.Fprint(wr, string(notes))
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

	fmt.Fprint(wr, resp)
}
