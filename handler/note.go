package handler

import (
	"strconv"

	"github.com/200lab-training-1/helper"
	"github.com/200lab-training-1/models"
	"github.com/200lab-training-1/repositories"
	"github.com/gin-gonic/gin"
)

type users struct {
	Title     string `validate:"required"`
	Completed bool
	UserID    uint
}

var (
	pagination helper.Pagination
)

func NoteCreate(c *gin.Context, noteRepo repositories.NoteRepo) (*models.Note, error) {
	note := models.Note{}
	if err := c.ShouldBind(&note); err != nil {
		return nil, err
	}
	return noteRepo.Create(note)
}

func NoteGet(c *gin.Context, noteRepo repositories.NoteRepo) (*[]models.Note, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	return noteRepo.Find(uint(id))
}

func NoteUpdate(c *gin.Context, noteRepo repositories.NoteRepo) error {
	id, _ := strconv.Atoi(c.Param("id"))
	note := models.Note{}
	if err := c.ShouldBind(&note); err != nil {
		return err
	}
	return noteRepo.Update(uint(id), note)
}

func NoteDelete(c *gin.Context, noteRepo repositories.NoteRepo) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return noteRepo.Delete(uint(id))
}

func NoteList(c *gin.Context, noteRepo repositories.NoteRepo) ([]models.Note, error) {
	c.ShouldBindQuery(&pagination)
	return noteRepo.List(pagination)
}
