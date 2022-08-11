package main

import (
	"fmt"
	"log"
	"net/http"
	"progate_crud_golang/controllers"
	"progate_crud_golang/models"

	"github.com/julienschmidt/httprouter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// connect ke database data.db
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// koneksikan ke model dengan db.automigrate

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// deklarasi router

	router := httprouter.New()

	todoController := &controllers.TodoController{}

	router.ServeFiles("/img/*filepath", http.Dir("img"))

	// http method GET AND POST
	router.GET("/", todoController.Index)
	router.GET("/create", todoController.Create)
	router.POST("/create", todoController.Create)
	router.GET("/edit/:id", todoController.Edit)
	router.POST("/edit/:id", todoController.Edit)
	router.GET("/done/:id", todoController.Done)

	log.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
