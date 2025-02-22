package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/veD-tnayrB/todo-app/cmd/docs"
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/models"
	todoHandler "github.com/veD-tnayrB/todo-app/internal/handlers/todo"
	todoRepository "github.com/veD-tnayrB/todo-app/internal/repositories/todo"
	todoService "github.com/veD-tnayrB/todo-app/internal/services/todo"
)

// Lets gonna stop here for now, tomorrow i need to work

// @TODO: Bryant motherfucker, heres the todos (tdah-prof aclarations):
// - Standarized Responses [READY]
// - Standarized Errors FOR LAYER AND FOR MODULE [READY]
// - Generate the anotations for OPENAPI and fix the routes so it can be OPENAPI with the standar you motherfucker
// - Implement a logger foor error handling (Could consider centralized logging for scalability.)
// - Unit Testing and mocking
// - Dockerization??? Maybe
// - Autodeployment when pushing to the repo

func main() {
	db := db.NewDB()

	// Simulates the existing data in DB
	db["1"] = models.Todo{Id: "1", Title: "Code", Completed: false}
	db["2"] = models.Todo{Id: "2", Title: "Eat", Completed: true}

	// Dependency injection :)
	todoRepository := todoRepository.TodoRepository{DB: db}
	todoService := todoService.TodoService{TodoRepository: &todoRepository}
	todoHandler := todoHandler.TodoHandler{TodoService: &todoService}

	router := gin.Default()
	todoGroup := router.Group("/todo")
	todoGroup.GET("/list", todoHandler.List)
	todoGroup.GET("/get/:id", todoHandler.Get)
	todoGroup.POST("/create", todoHandler.Create)
	todoGroup.PUT("/update/:id", todoHandler.Update)
	todoGroup.DELETE("/delete/:id", todoHandler.Remove)

	// Documentation
	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5002")
}
