package router

import (
	"github.com/gin-gonic/gin"
	"url-shortening-service/internal/db"
	"url-shortening-service/internal/handler"
	"url-shortening-service/internal/repository"
	"url-shortening-service/internal/service"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	db.InitDB()
	r.LoadHTMLGlob("internal/templates/*")
	sh := handler.NewShortURLHandler(service.NewURLShorteningService(repository.NewURLShorteningRepository(db.DB)))
	r.POST("/shorten", sh.CreateShortURL)
	r.GET("/shorten/:shortCode", sh.Retrieve)
	r.PUT("/shorten/:shortCode", sh.UpdateShortURL)
	r.DELETE("/shorten/:shortCode", sh.DeleteShortURL)
	r.GET("/shorten/:shortCode/stats", sh.Statistics)
	return r
}
