package repository

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func NewRepos(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
