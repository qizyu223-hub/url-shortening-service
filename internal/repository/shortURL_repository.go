package repository

import (
	"gorm.io/gorm"
	"url-shortening-service/internal/model"
)

type URLShorteningRepository struct {
	db *gorm.DB
}

func NewURLShorteningRepository(db *gorm.DB) *URLShorteningRepository {
	return &URLShorteningRepository{db: db}
}

func (ur *URLShorteningRepository) Create(shortURL *model.ShortURL) (*model.ShortURL, error) {
	err := ur.db.Create(shortURL).Error
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (ur *URLShorteningRepository) GetByShortCode(shortCode string) (*model.ShortURL, error) {
	shortURL := &model.ShortURL{}
	err := ur.db.Where("short_code = ?", shortCode).First(shortURL).Error
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (ur *URLShorteningRepository) Update(shortURL *model.ShortURL) (*model.ShortURL, error) {
	err := ur.db.Save(shortURL).Error
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (ur *URLShorteningRepository) Delete(shortURL *model.ShortURL) error {
	return ur.db.Delete(shortURL).Error
}
