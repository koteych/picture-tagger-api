package main

import (
	"log"

	"picture_tagger_api/internal/handler"
	"picture_tagger_api/internal/repository"
	"picture_tagger_api/internal/service"
	"picture_tagger_api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	pictureRepo := repository.NewSQLPictureRepository(db)
	tagRepo := repository.NewSQLTagRepository(db)

	pictureService := service.NewPictureService(pictureRepo, tagRepo)

	//somePicture := model.Picture{ID: 100, Title: "some picture"}

	//ictureRepo.CreatePicture(&somePicture)

	//someTag := model.Tag{ID: 100, Name: "New tag", Alias: "new_tag_alias", Desc: "", IsHidden: false}
	//_, err = tagRepo.CreateTag(&someTag)

	pictureHandler := &handler.PictureHandler{
		PictureService: *pictureService,
	}

	r.POST("/api/pictures/:picture_id/assign-tag/:tag_id", pictureHandler.AssignTagById)

	r.Run()
}
