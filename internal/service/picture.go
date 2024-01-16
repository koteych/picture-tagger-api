package service

import (
	"picture_tagger_api/internal/repository"
)

type PictureService struct {
	pictureRepo repository.PictureRepository
}

func NewPictureService(pictureRepo repository.PictureRepository) *PictureService {
	return &PictureService{pictureRepo: pictureRepo}
}
