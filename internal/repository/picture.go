package repository

import (
	"database/sql"
	"picture_tagger_api/internal/model"
)

type SQLPictureRepository struct {
	db *sql.DB
}

type PictureRepository interface {
	GetPictureByID(id int) (*model.Picture, error)
	CreatePicture(picture *model.Picture) error
	UpdatePicture(picture *model.Picture) error
	DeletePicture(id int) error
}

func NewSQLPictureRepository(db *sql.DB) *SQLPictureRepository {
	return &SQLPictureRepository{db: db}
}

func (r *SQLPictureRepository) CreatePicture(picture *model.Picture) error {
	return nil
}

func (r *SQLPictureRepository) DeletePicture(id int) error {
	return nil
}

func (r *SQLPictureRepository) GetPictureByID(id int) (*model.Picture, error) {
	return nil, nil
}

func (r *SQLPictureRepository) UpdatePicture(picture *model.Picture) error {
	return nil
}
