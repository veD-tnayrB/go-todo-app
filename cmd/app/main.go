package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/veD-tnayrB/todo-app/cmd/docs"
	"github.com/veD-tnayrB/todo-app/common/db"
	"github.com/veD-tnayrB/todo-app/common/logger"
	"github.com/veD-tnayrB/todo-app/common/middlewares"
	"github.com/veD-tnayrB/todo-app/common/models"
	rate_limiter "github.com/veD-tnayrB/todo-app/common/rate-limiter"
	todoHandler "github.com/veD-tnayrB/todo-app/internal/todo/handlers"
	todoRepository "github.com/veD-tnayrB/todo-app/internal/todo/repositories"
	todoService "github.com/veD-tnayrB/todo-app/internal/todo/services"
)

// @TODO: Bryant, take care:
// - Unit Testing
// - Dockerization??? Maybe
// - Autodeployment when pushing to the repo

func main() {
	err := gotenv.Load()
	if err != nil {
		panic("error while trying to load the environment variables")
	}

	db := db.NewDB()
	logger, err := logger.NewLogger("logs")
	if err != nil {
		panic(err)
	}

	maxRequestStr := os.Getenv("MAX_REQUEST")
	maxRequest, err := strconv.Atoi(maxRequestStr)
	if err != nil {
		panic("something went wrong while getting the max request environment variable")
	}

	refreshTimeStr := os.Getenv("LIMIT_REFRESH_TIME")
	refreshTime, err := strconv.Atoi(refreshTimeStr)
	if err != nil {
		panic("something went wrong while getting the refresh time environment variable")
	}

	rateLimiter := rate_limiter.NewRateLimiter(int(maxRequest), time.Duration(refreshTime)*time.Second)

	// Simulates the existing data in DB
	db["1"] = models.Todo{Id: "1", Title: "Code", Completed: false}
	db["2"] = models.Todo{Id: "2", Title: "Eat", Completed: true}

	// Dependency injection :)
	todoRepository := todoRepository.TodoRepository{DB: db, Logger: logger}
	todoService := todoService.TodoService{TodoRepository: &todoRepository, Logger: logger}
	todoHandler := todoHandler.TodoHandler{TodoService: &todoService, Logger: logger}

	router := gin.Default()
	router.Use(middlewares.RateLimiterMiddleware(rateLimiter))
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
