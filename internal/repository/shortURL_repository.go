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

func (ur *URLShorteningRepository) WithTx(fn func(repository *URLShorteningRepository) error) error {
	return ur.db.Transaction(func(tx *gorm.DB) error {
		return fn(&URLShorteningRepository{db: tx})
	})
}

func (ur *URLShorteningRepository) Create(shortURL *model.ShortURL) error {
	return ur.db.Create(shortURL).Error
}

func (ur *URLShorteningRepository) GetByShortCode(shortCode string) (*model.ShortURL, error) {
	var shortURL model.ShortURL
	err := ur.db.Where("short_code = ?", shortCode).First(&shortURL).Error
	if err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (ur *URLShorteningRepository) Update(shortURL *model.ShortURL) error {
	return ur.db.Save(shortURL).Error
}

func (ur *URLShorteningRepository) DeleteByShortCode(shortCode string) error {
	tx := ur.db.Where("short_code = ?", shortCode).Delete(&model.ShortURL{})
	return tx.Error
}

func (ur *URLShorteningRepository) IncrAccessCount(shortCode string) error {
	return ur.db.Where("short_code = ?", shortCode).
		Update("access_count", gorm.Expr("access_count + ?", 1)).Error
}
