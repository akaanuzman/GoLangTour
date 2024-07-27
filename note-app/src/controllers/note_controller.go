package controllers

import (
	"net/http"
	"note-app/src/models"
	"note-app/src/services"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteController struct {
	noteService services.INoteService
}

func NewNoteController(noteService services.INoteService) *NoteController {
	return &NoteController{noteService: noteService}
}

func (controller *NoteController) GetAllNotes(c echo.Context) error {
	notes, err := controller.noteService.GetAllNotes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, notes)
}

func (controller *NoteController) GetNoteById(c echo.Context) error {
	id := c.Param("id")
	note, err := controller.noteService.GetNoteById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, note)
}

func (controller *NoteController) GetNoteByUserId(c echo.Context) error {
	userId := c.Param("userId")
	notes, err := controller.noteService.GetNoteByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, notes)
}

func (controller *NoteController) GetNotesByTags(c echo.Context) error {
	tagsParam := c.QueryParam("tags")
	tags := strings.Split(tagsParam, ",")
	notes, err := controller.noteService.GetNotesByTags(tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, notes)
}

func (controller *NoteController) CreateNote(c echo.Context) error {
	var note models.Note
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Extract user ID from context (JWT token)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": "Invalid user ID"})
	}

	if err := controller.noteService.CreateNote(&note, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, note)
}

func (controller *NoteController) UpdateNote(c echo.Context) error {
	var note models.Note
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := c.Param("id")
	note.Id, _ = primitive.ObjectIDFromHex(id)

	updatedNote, err := controller.noteService.UpdateNote(&note)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedNote)
}

func (controller *NoteController) DeleteNote(c echo.Context) error {
	id := c.Param("id")
	if err := controller.noteService.DeleteNote(id); err != nil {
		return c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, bson.M{"message": "Note deleted"})
}
