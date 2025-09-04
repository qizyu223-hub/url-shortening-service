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
	var out *model.ShortURL
	err := urs.repo.WithTx(func(txRepo *repository.URLShorteningRepository) error {
		if err := txRepo.Create(&shortURL); err != nil {
			return err
		}
		out = &shortURL
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (urs *URLShorteningService) GetByShortCode(shortCode string) (*model.ShortURL, error) {
	return urs.repo.GetByShortCode(shortCode)
}

func (urs *URLShorteningService) UpdateURL(shortCode, url string) (*model.ShortURL, error) {
	var out *model.ShortURL
	err := urs.repo.WithTx(func(txRepo *repository.URLShorteningRepository) error {
		shortURL, err := txRepo.GetByShortCode(shortCode)
		if err != nil {
			return err
		}
		shortURL.URL = url
		err = urs.repo.Update(shortURL)
		if err != nil {
			return err
		}
		out = shortURL
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (urs *URLShorteningService) IncrAccessCount(shortCode string) error {
	return urs.repo.IncrAccessCount(shortCode)
}

func (urs *URLShorteningService) Delete(shortCode string) error {
	return urs.repo.DeleteByShortCode(shortCode)
}
