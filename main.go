package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Sprintln(db)

	pictureRepo := repository.NewSQLPictureRepository(db)
	pictureService := service.NewPictureService(pictureRepo)

	fmt.Println(pictureService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
