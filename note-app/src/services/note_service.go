package services

import (
	"context"
	"errors"
	"note-app/src/core/db"
	"note-app/src/models"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type INoteService interface {
	GetAllNotes() ([]models.Note, error)
	GetNoteById(id string) (*models.Note, error)
	GetNoteByUserId(userId string) ([]models.Note, error)
	GetNotesByTags(tags []string) ([]models.Note, error)
	CreateNote(note *models.Note, id primitive.ObjectID) error
	UpdateNote(note *models.Note) (*models.Note, error)
	DeleteNote(id string) error
}

type NoteService struct {
	noteCollection *mongo.Collection
	context        context.Context
}

func NewNoteService() INoteService {
	db := db.Database{}
	db.ConnectDB()
	collection := db.GetCollection("notes")
	return &NoteService{
		noteCollection: collection,
		context:        context.Background(),
	}
}

func (noteService *NoteService) GetAllNotes() ([]models.Note, error) {
	var notes []models.Note
	cursor, err := noteService.noteCollection.Find(noteService.context, bson.M{
		"isActive": true,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(noteService.context)
	if err := cursor.All(noteService.context, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (noteService *NoteService) GetNoteById(id string) (*models.Note, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid note ID")
	}

	var note models.Note
	err = noteService.noteCollection.FindOne(noteService.context, bson.M{"_id": objectId, "isActive": true}).Decode(&note)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Note not found
		}
		return nil, err
	}
	return &note, nil
}

func (noteService *NoteService) GetNoteByUserId(userId string) ([]models.Note, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var notes []models.Note
	cursor, err := noteService.noteCollection.Find(noteService.context, bson.M{
		"userId":   objectId,
		"isActive": true,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(noteService.context)

	if err := cursor.All(noteService.context, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (noteService *NoteService) GetNotesByTags(tags []string) ([]models.Note, error) {
	var notes []models.Note
	cursor, err := noteService.noteCollection.Find(noteService.context, bson.M{
		"tags":     bson.M{"$in": tags},
		"isActive": true,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(noteService.context)
	if err := cursor.All(noteService.context, &notes); err != nil {
		return nil, err
	}
	return notes, nil
}

func (noteService *NoteService) CreateNote(note *models.Note, id primitive.ObjectID) error {
	note.Id = primitive.NewObjectID()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	note.IsActive = true
	note.UserId = id
	_, err := noteService.noteCollection.InsertOne(noteService.context, note)
	if err != nil {
		return err
	}
	return nil
}

func (noteService *NoteService) UpdateNote(note *models.Note) (*models.Note, error) {
	note.UpdatedAt = time.Now()

	updateFields := bson.M{}
	val := reflect.ValueOf(note).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()
		tag := field.Tag.Get("bson")

		// If the field is set and it's not zero value, add to updateFields
		if tag != "" && !reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			updateFields[tag] = value
		}
	}

	// Ensure UpdatedAt is always updated
	updateFields["updatedAt"] = note.UpdatedAt

	update := bson.M{"$set": updateFields}
	_, err := noteService.noteCollection.UpdateOne(noteService.context, bson.M{"_id": note.Id}, update)
	if err != nil {
		return nil, err
	}

	// Retrieve updated document
	var updatedNote models.Note
	err = noteService.noteCollection.FindOne(noteService.context, bson.M{"_id": note.Id}).Decode(&updatedNote)
	if err != nil {
		return nil, err
	}

	return &updatedNote, nil
}

func (noteService *NoteService) DeleteNote(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{"isActive": false}}
	_, err = noteService.noteCollection.UpdateOne(noteService.context, bson.M{"_id": objectId}, update)
	if err != nil {
		return err
	}
	return nil
}
