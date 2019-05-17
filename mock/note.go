package mock

import (
	"github.com/200lab-training-1/helper"
	"github.com/200lab-training-1/models"
	"github.com/stretchr/testify/mock"
)

type NoteRepoImpl struct {
	mock.Mock
}

func (noteRepo *NoteRepoImpl) Create(note models.Note) (*models.Note, error) {
	args := noteRepo.Called(note)
	return args.Get(0).(*models.Note), args.Error(1)
}

func (noteRepo *NoteRepoImpl) Find(id uint) (*models.Note, error) {
	args := noteRepo.Called(id)
	return args.Get(0).(*models.Note), args.Error(1)
}

func (noteRepo *NoteRepoImpl) List(pagination helper.Pagination) ([]models.Note, error) {
	args := noteRepo.Called(pagination)
	return args.Get(0).([]models.Note), args.Error(1)
}

func (noteRepo *NoteRepoImpl) Update(id uint, note models.Note) error {
	args := noteRepo.Called(uint(id), note)
	return args.Error(0)
}

func (noteRepo *NoteRepoImpl) Delete(id uint) error {
	args := noteRepo.Called(uint(id))
	return args.Error(0)
}
