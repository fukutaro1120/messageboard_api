package main

import (
	"message-board-api/controller"
	"message-board-api/infra/mysql"
	"message-board-api/usecase"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	articleRepo := mysql.NewArticleRepository()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{"*"},
		MaxAge:       24 * time.Hour,
	}))

	// ping := &controller.PingController{}
	// r.GET("/ping", ping.Ping)

	article := &controller.ArticleController{
		UseCase: usecase.NewArticleUseCase(articleRepo),
	}
	r.GET("/articles", article.List)

	r.POST("/articles", article.Create)

	r.DELETE("/articles", article.Delete)

	r.PUT("/articles", article.Update)

	r.Run(":3000")
}

// func createDummyArticles(articleRepo repository.ArticleRepository) {
// 	articles := []model.Article{
// 		{Name: "user1", Message: "hello"},
// 		{Name: "user2", Message: "nice to meet you"},
// 		{Name: "user3", Message: "long time see you"},
// 		{Name: "user4", Message: "glad to see you"},
// 		{Name: "user5", Message: "miss you"},
// 		{Name: "user6", Message: "so good"},
// 		{Name: "user7", Message: "absolutely"},
// 	}

// 	ctx := context.Background()
// 	for _, a := range articles {
// 		err := articleRepo.Create(ctx, &a)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// }
