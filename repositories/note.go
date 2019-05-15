package repositories

import (
	"github.com/200lab-training-1/models"
	"github.com/jinzhu/gorm"
)

// NoteRepoImpl is struct that define which type DB we use
type NoteRepoImpl struct {
	DB *gorm.DB
}

// NoteRepo is interface that contains some method we must implement
type NoteRepo interface {
	Delete(uint) error
	Update(uint, models.Note) error
	Find(uint) (*models.Note, error)
	Create(models.Note) (*models.Note, error)
}

// Delete is func delete a note
func (noteRepo *NoteRepoImpl) Delete(id uint) error {
	err := noteRepo.DB.Where("id = ?", id).Delete(&models.Note{}).Error
	return err
}

// Update is func update a note
func (noteRepo *NoteRepoImpl) Update(id uint, note models.Note) error {
	err := noteRepo.DB.Where("id = ?", id).Update(&note).Error
	return err
}

// Find is func find a note based on id
func (noteRepo *NoteRepoImpl) Find(id uint) (*models.Note, error) {
	note := &models.Note{}
	err := noteRepo.DB.Where("id = ?", id).First(note).Error
	return note, err
}

// Create a func create a note
func (noteRepo *NoteRepoImpl) Create(note models.Note) (*models.Note, error) {
	err := noteRepo.DB.Create(&note).Error
	return &note, err
}
