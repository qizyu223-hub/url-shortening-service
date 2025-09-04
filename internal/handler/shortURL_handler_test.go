package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"url-shortening-service/internal/dto"
	"url-shortening-service/internal/model"
	"url-shortening-service/internal/repository"
	"url-shortening-service/internal/service"
)

func startSqlite() *gorm.DB {
	sqliteDB, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = sqliteDB.AutoMigrate(&model.ShortURL{})
	return sqliteDB
}

// 用 SQLite 内存版，更快；如需真实 Postgres 可换 startPostgres
func setupRouter(DB *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := repository.NewURLShorteningRepository(DB)
	svc := service.NewURLShorteningService(repo)
	h := NewShortURLHandler(svc)

	r := gin.New()
	r.POST("/shorten", h.CreateShortURL)
	r.GET("/shorten/:shortCode", h.Retrieve)
	return r
}

func TestCreateShortURL(t *testing.T) {
	r := setupRouter(startSqlite())

	tests := []struct {
		name       string
		body       interface{}
		wantStatus int
	}{
		{"200 ok", dto.Request{URL: "https://a.com"}, http.StatusCreated},
		{"400 bad json", "not json", http.StatusBadRequest},
		{"400 empty url", dto.Request{URL: ""}, http.StatusBadRequest}, // service 返回 err
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			_ = json.NewEncoder(&buf).Encode(tt.body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/shorten", &buf)
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestRetrieve(t *testing.T) {
	DB := startSqlite()
	r := setupRouter(DB)
	// 预插入一条
	repo := repository.NewURLShorteningRepository(DB)
	svc := service.NewURLShorteningService(repo)
	su, _ := svc.Create("https://b.com")

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{"200 redirect", "/shorten/" + su.ShortCode, http.StatusFound},
		{"404 not found", "/shorten/nosuch", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.path, nil)
			r.ServeHTTP(w, req)
			require.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
