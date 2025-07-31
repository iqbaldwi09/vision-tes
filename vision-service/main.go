package main

import (
	"time"
	"vision-service/config"
	"vision-service/handler"
	"vision-service/repository"
	"vision-service/routes"
	"vision-service/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ✅ Tambahkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8888"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 🔌 Inisialisasi database
	db := config.InitDB()

	// 🔄 Inisialisasi repository, usecase, handler
	postRepo := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepo)
	postHandler := handler.NewPostHandler(postUsecase)

	// 📦 Daftarkan route
	routes.ArticleRoutes(r, postHandler)

	// 🔐 Proxy trust
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 🚀 Jalankan server
	r.Run(":8080")
}
