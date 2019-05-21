package mock

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/200lab-training-1/handler"
	"github.com/200lab-training-1/models"
	"github.com/gin-gonic/gin"
)

func buildMockContext(method string, path string, data string) *gin.Context {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(method, path, strings.NewReader(data))
	ctx.Request.Header.Set("Content-Type", "application/json")
	return ctx
}

func Test_NoteCreate_AllFieldAreValid(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	data := `{"title": "Should do homework","completed": false, "user_id": 1}`
	note := models.Note{}
	json.Unmarshal([]byte(data), &note)
	ctx := buildMockContext("POST", "/note", data)
	noteRepo := new(NoteRepoImpl)
	expected := models.Note{
		Title:     "Should do homework",
		Completed: false,
		UserID:    1,
	}
	noteRepo.On("Create", note).Return(&expected, nil)
	actual, err := handler.NoteCreate(ctx, noteRepo)
	if err != nil {
		t.Error("Error should not be nil")
	}
	if actual.ID != expected.ID {
		t.Error("Actual note should be same expected note")
	}
}
