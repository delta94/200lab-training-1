package routes

import (
	"github.com/200lab-training-1/handler"
	"github.com/200lab-training-1/middlewares"
	"github.com/200lab-training-1/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	jwtSecretKey = []byte("ThisIsAVerySecretKey")
	identityKey  = "identityKey"
)

func log(c *gin.Context, err error, result interface{}) {
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}

func initUserRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.POST("/signup", func(c *gin.Context) {
		userRepo := &repositories.UserRepoImpl{
			DB: db,
		}
		result, err := handler.SignUp(c, userRepo)
		log(c, err, result)

	})

	engine.POST("/login", func(c *gin.Context) {
		userRepo := &repositories.UserRepoImpl{
			DB: db,
		}
		result, err := handler.Login(c, userRepo)
		log(c, err, result)
	})
}

func initNoteRoutes(engine *gin.Engine, db *gorm.DB) {
	groupRouter := engine.Group("/note")
	groupRouter.Use(middlewares.AuthenMiddleware)
	{
		groupRouter.GET(":/id", func(c *gin.Context) {
			noteRepo := &repositories.NoteRepoImpl{
				DB: db,
			}
			result, err := handler.NoteGet(c, noteRepo)
			log(c, err, result)
		})

		groupRouter.POST("", func(c *gin.Context) {
			noteRepo := &repositories.NoteRepoImpl{
				DB: db,
			}
			result, err := handler.NoteCreate(c, noteRepo)
			log(c, err, result)
		})

		groupRouter.PUT("/:id", func(c *gin.Context) {
			noteRepo := &repositories.NoteRepoImpl{
				DB: db,
			}
			err := handler.NoteUpdate(c, noteRepo)
			log(c, err, nil)
		})

		groupRouter.DELETE("/:id", func(c *gin.Context) {
			noteRepo := &repositories.NoteRepoImpl{
				DB: db,
			}
			err := handler.NoteDelete(c, noteRepo)
			log(c, err, nil)
		})
	}
}

func InitRoutes(engine *gin.Engine, db *gorm.DB) {
	initNoteRoutes(engine, db)
	initUserRoutes(engine, db)
}