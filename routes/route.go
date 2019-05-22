package routes

import (
	"github.com/200lab-training-1/handler"
	"github.com/200lab-training-1/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	groupRouter := engine.Group("v1/users")

	groupRouter.Use()
	{
		groupRouter.POST("/register", func(c *gin.Context) {
			userRepo := &repositories.UserRepoImpl{
				DB: db,
			}
			result, err := handler.SignUp(c, userRepo)
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(201, result)
		})

		groupRouter.POST("/authen", func(c *gin.Context) {
			userRepo := &repositories.UserRepoImpl{
				DB: db,
			}
			result, err := handler.Login(c, userRepo)
			if err != nil {
				c.AbortWithStatusJSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, result)
		})

	}
}

func initNoteRoutes(engine *gin.Engine, db *gorm.DB) {
	groupRouter := engine.Group("v1/notes")
	groupRouter.Use()
	{
		groupRouter.GET("/:id", func(c *gin.Context) {
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
