package service

import (
	"url-shortening-service/internal/model"
	"url-shortening-service/internal/repository"
	"url-shortening-service/internal/utils"
)

type URLShorteningService struct {
	repo *repository.URLShorteningRepository
}

func NewURLShorteningService(repo *repository.URLShorteningRepository) *URLShorteningService {
	return &URLShorteningService{repo: repo}
}

func (urs *URLShorteningService) Create(url string) (*model.ShortURL, error) {
	var shortURL model.ShortURL
	shortURL.URL = url
	shortCode := utils.GenerateShortCode(url)
	shortURL.ShortCode = shortCode
	return urs.repo.Create(&shortURL)
}

func (urs *URLShorteningService) GetByShortCode(shortCode string) (*model.ShortURL, error) {
	return urs.repo.GetByShortCode(shortCode)
}

func (urs *URLShorteningService) UpdateURL(shortCode, url string) (*model.ShortURL, error) {
	shortURL, err := urs.GetByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	shortURL.URL = url
	return urs.repo.Update(shortURL)
}

func (urs *URLShorteningService) UpdateAC(shortCode string) error {
	shortURL, err := urs.GetByShortCode(shortCode)
	if err != nil {
		return err
	}
	shortURL.AccessCount++
	_, err = urs.repo.Update(shortURL)
	return err
}

func (urs *URLShorteningService) Delete(shortCode string) error {
	su, err := urs.GetByShortCode(shortCode)
	if err != nil {
		return err
	}
	return urs.repo.Delete(su)
}
