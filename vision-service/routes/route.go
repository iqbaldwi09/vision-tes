package routes

import (
	"vision-service/handler"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine, postHandler *handler.PostHandler) {
	r.POST("/article", postHandler.CreateArticle)
	r.GET("/articles", postHandler.GetAllArticle)
	r.GET("/article/:id", postHandler.GetArticleByID)
	r.PUT("/article/:id", postHandler.UpdateArticle)
	r.DELETE("/article/:id", postHandler.DeleteArticle)
}
