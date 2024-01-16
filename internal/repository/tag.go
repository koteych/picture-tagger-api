package repository

import (
	"database/sql"
	"picture_tagger_api/internal/model"
)

type SQLTagRepository struct {
	db *sql.DB
}

type TagRepository interface {
	CreateTag(tag *model.Tag) error
}

func NewSQLTagRepository(db *sql.DB) *SQLTagRepository {
	return &SQLTagRepository{db: db}
}

func (r *SQLTagRepository) CreateTag(tag *model.Tag) error {
	return nil
}
