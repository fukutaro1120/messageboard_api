package router

import (
	"message-board-api/controller"
	"message-board-api/repository"
	"message-board-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Config ...
type Config struct {
	MySQLArticleRepo repository.ArticleRepository
}

// New ...
func New(config *Config) http.Handler {
	router := gin.Default()
	// router.LoadHTMLGlob("template/*")
	{
		article := &controller.ArticleController{
			UseCase: usecase.NewArticleUseCase(config.MySQLArticleRepo),
		}
		router.GET("/articles", article.List)
		router.POST("/articles", article.Create)
		router.DELETE("/articles", article.Delete)
		router.PUT("/articles", article.Update)

		router.Run(":3000")
	}
	return router
}
