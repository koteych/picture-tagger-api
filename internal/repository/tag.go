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

func (r *SQLTagRepository) CreateTag(tag *model.Tag) (*model.Tag, error) {
	_, err := r.db.Exec(
		"INSERT INTO tags (tag_name, tag_alias, tag_desc, is_hidden) VALUES (?, ?, ?, ?)",
		tag.Name, tag.Alias, tag.Desc, tag.IsHidden)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
