package handler

import (
	"net/http"
	"url-shortening-service/internal/dto"
	"url-shortening-service/internal/service"
)
import "github.com/gin-gonic/gin"

type ShortURLHandler struct {
	svc *service.URLShorteningService
}

func NewShortURLHandler(svc *service.URLShorteningService) *ShortURLHandler {
	return &ShortURLHandler{svc: svc}
}

func (h *ShortURLHandler) CreateShortURL(c *gin.Context) {
	var req dto.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortURL, err := h.svc.Create(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := dto.Response{
		ID:        shortURL.ID,
		URL:       shortURL.URL,
		ShortCode: shortURL.ShortCode,
		CreatedAt: shortURL.CreatedAt,
		UpdatedAt: shortURL.UpdatedAt,
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *ShortURLHandler) Retrieve(c *gin.Context) {
	shortCode := c.Param("shortCode")
	shortURL, err := h.svc.GetByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	resp := dto.Response{
		ID:        shortURL.ID,
		URL:       shortURL.URL,
		ShortCode: shortCode,
		CreatedAt: shortURL.CreatedAt,
		UpdatedAt: shortURL.UpdatedAt,
	}
	err = h.svc.UpdateAC(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "shorten.html", gin.H{"resp": resp})
}

func (h *ShortURLHandler) UpdateShortURL(c *gin.Context) {
	var req dto.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortCode := c.Param("shortCode")
	shortURL, err := h.svc.UpdateURL(shortCode, req.URL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	resp := dto.Response{
		ID:        shortURL.ID,
		URL:       shortURL.URL,
		ShortCode: shortCode,
		CreatedAt: shortURL.CreatedAt,
		UpdatedAt: shortURL.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *ShortURLHandler) DeleteShortURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	err := h.svc.Delete(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"success": true,
	})
}

func (h *ShortURLHandler) Statistics(c *gin.Context) {
	shortCode := c.Param("shortCode")
	shortURL, err := h.svc.GetByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shortURL)
}
