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

// @TODO: Bryant, take care:
// Im gonna stop with the logger because looks like a deeper topic and ill love to have the time
// to do research before doing something stupid
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
	todoGroup := router.Group("/todos")
	todoGroup.GET("", todoHandler.List)
	todoGroup.GET("/:id", todoHandler.Get)
	todoGroup.POST("", todoHandler.Create)
	todoGroup.PUT("/:id", todoHandler.Update)
	todoGroup.DELETE("/:id", todoHandler.Remove)

	// Documentation
	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5002")
}
