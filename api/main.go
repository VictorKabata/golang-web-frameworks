package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/home", Home)
	router.GET("/notes", GetNotes)
	router.GET("/note/:id", GetNote)
	router.POST("/create/note", CreateNote)
	router.PUT("/update/note/:id", UpdateNote)
	router.DELETE("/delete/notes", DeleteAllNotes)
	router.DELETE("/delete/note", DeleteNote)

	log.Fatal(router.Run("localhost:8080"))
}

func Home(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, map[string]string{"message": "Home page"})
}

func GetNotes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, notes)
}

func GetNote(ctx *gin.Context) {
	id := ctx.Query("id")

	for _, v := range notes {
		if strconv.Itoa(v.ID) == id {
			ctx.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}

func CreateNote(ctx *gin.Context) {
	var newNote models.Note

	err := ctx.BindJSON(&newNote)
	if err != nil {
		return
	}

	notes = append(notes, newNote)

	ctx.IndentedJSON(http.StatusCreated, notes)
}

func UpdateNote(ctx *gin.Context) {
	var updatedNote models.Note

	id := ctx.Query("id")

	err := ctx.BindJSON(&updatedNote)
	if err != nil {
		return
	}

	for i, v := range notes {
		if strconv.Itoa(v.ID) == id {
			v = models.Note{
				ID:          updatedNote.ID,
				Title:       updatedNote.Title,
				Description: updatedNote.Description,
			}

			notes = append(notes[:i], v)
		}
	}

	ctx.IndentedJSON(http.StatusCreated, notes)
}

func DeleteAllNotes(ctx *gin.Context) {
	notes = notes[:0]

	err := ctx.BindJSON(&notes)
	if err != nil {
		ctx.IndentedJSON(http.StatusNoContent, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, notes)
}

func DeleteNote(ctx *gin.Context) {
	id := ctx.Query("id")

	for i, v := range notes {
		if strconv.Itoa(v.ID) == id {
			notes = append(notes[:i], notes[i+1:]...)
			return
		}
	}

	err := ctx.BindJSON(&notes)
	if err != nil {
		ctx.IndentedJSON(http.StatusNoContent, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, notes)
}
