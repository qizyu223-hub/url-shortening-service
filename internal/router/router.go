package router

import (
	"github.com/gin-gonic/gin"
	"url-shortening-service/internal/db"
	"url-shortening-service/internal/handler"
	"url-shortening-service/internal/repository"
	"url-shortening-service/internal/service"
)

func NewRouter() (*gin.Engine, error) {
	DB, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	r := gin.Default()
	sh := handler.NewShortURLHandler(service.NewURLShorteningService(repository.NewURLShorteningRepository(DB)))
	r.POST("/shorten", sh.CreateShortURL)
	r.GET("/shorten/:shortCode", sh.Retrieve)
	r.PUT("/shorten/:shortCode", sh.UpdateShortURL)
	r.DELETE("/shorten/:shortCode", sh.DeleteShortURL)
	r.GET("/shorten/:shortCode/stats", sh.Statistics)
	return r, nil
}
