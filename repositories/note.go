package repositories

import (
	"github.com/200lab-training-1/helper"
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
	Find(uint) (*[]models.Note, error)
	Create(models.Note) (*models.Note, error)
	List(helper.Pagination) ([]models.Note, error)
}

// Delete is func delete a note
func (noteRepo *NoteRepoImpl) Delete(id uint) error {
	err := noteRepo.DB.Where("id = ?", id).Delete(&models.Note{}).Error
	return err
}

// Update is func update a note
func (noteRepo *NoteRepoImpl) Update(id uint, note models.Note) error {
	err := noteRepo.DB.Model(&models.Note{}).Where("id = ?", id).Update(note).Error
	return err
}

// Find is func find a note based on id
func (noteRepo *NoteRepoImpl) Find(id uint) (*[]models.Note, error) {
	notes := &[]models.Note{}
	err := noteRepo.DB.Find(notes).Where("user_id = ?", id).Error
	return notes, err
}

// Create a func create a note
func (noteRepo *NoteRepoImpl) Create(note models.Note) (*models.Note, error) {
	err := noteRepo.DB.Create(&note).Error
	return &note, err
}

func (noteRepo *NoteRepoImpl) List(pagination helper.Pagination) ([]models.Note, error) {
	notes := []models.Note{}
	offset := pagination.GetOffSet()
	limit := pagination.GetLimit()
	err := noteRepo.DB.Offset(offset).Limit(limit).Find(&notes).Error
	return notes, err
}
