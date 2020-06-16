package main

import (
	"TodoApp/middleware"
	Note "TodoApp/notes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/notesappdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Note.Note{})

	return db
}

func main() {
	db := InitDB()
	defer db.Close()
	notesRepo := Note.ProvideNotesRepository(db)
	notesService := Note.ProvideNoteService(notesRepo)
	notesApi := Note.ProvideNoteAPI(notesService)

	r := gin.Default()
	r.Use(middleware.Auth())
	r.POST("/v1/notes/create", notesApi.Create)
	r.GET("/v1/notes/:id", notesApi.FindById)
	r.DELETE("/v1/notes/:id", notesApi.Delete)
	r.GET("/v1/notes", notesApi.FindAll)
	r.PATCH("/v1/notes/:id", notesApi.Update)
	r.DELETE("/v1/notes", notesApi.DeleteAll)
	err := r.Run(":8081")
	if err != nil {
		panic(err)
	}
}
