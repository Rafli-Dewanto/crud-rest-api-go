package main

import (
	"log"
	_ "net/http"

	"github.com/Rafli-Dewanto/go-rest/config"
	"github.com/Rafli-Dewanto/go-rest/handlers"
	_ "github.com/Rafli-Dewanto/go-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()
	r := gin.Default()

	studentHandler := handlers.NewStudentHandler(db)
	r.GET("/students", studentHandler.GetStudents)
	r.GET("/students/:id", studentHandler.GetStudentById)
	r.POST("/students", studentHandler.CreateStudent)
	r.PATCH("/students/:id", studentHandler.UpdateStudent)
	r.DELETE("/students/:id", studentHandler.DeleteStudent)

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}