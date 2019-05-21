package validators

import (
	"log"

	"github.com/200lab-training-1/models"
	validator "gopkg.in/go-playground/validator.v9"
)

type UsersValidate struct {
	Title     string `validate:"required"`
	Completed bool
	UserID    uint
}

var (
	validate *validator.Validate
)

func TitleIsEmpty(note models.Note) error {
	validate = validator.New()
	users := &UsersValidate{
		Title:     note.Title,
		Completed: note.Completed,
		UserID:    note.UserID,
	}

	errs := validate.Struct(users)
	if errs != nil {
		log.Fatal("hello")
		return errs
	}

	return nil
}
