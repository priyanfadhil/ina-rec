package application

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/priyanfadhil/ina-rec/common/database"
	v "github.com/priyanfadhil/ina-rec/helper"
	taskController "github.com/priyanfadhil/ina-rec/service/api/task/controller"
	taskHandler "github.com/priyanfadhil/ina-rec/service/api/task/handler"
	taskRepository "github.com/priyanfadhil/ina-rec/service/api/task/repository"
	userController "github.com/priyanfadhil/ina-rec/service/api/user/controller"
	userHandler "github.com/priyanfadhil/ina-rec/service/api/user/handler"
	userRepository "github.com/priyanfadhil/ina-rec/service/api/user/repository"
	"github.com/priyanfadhil/ina-rec/service/middleware"
	"github.com/priyanfadhil/ina-rec/service/model/api_response"
)

type App struct {
	Engine         *gin.Engine
	DBManager      database.DatabaseManager
	UserController *userController.UserController
	TaskController *taskController.TaskController
	Validator      v.CustomValidator
}

func (app *App) Start(addr string) {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func(sig chan os.Signal) {
		if err := app.Engine.Run(addr); err != nil {
			sig <- syscall.SIGTERM
		}
	}(sig)

	<-sig
}

func (app *App) Initialize() {
	app.InitializeDomain()
	app.initializeRoutes()
}

func (app *App) InitializeDomain() {
	userRepository := userRepository.NewUserRepository()
	userHandler := userHandler.NewHandler(app.DBManager, userRepository)
	userController := userController.NewController(userHandler)

	taskRepository := taskRepository.NewTaskRepository()
	taskHandler := taskHandler.NewHandler(app.DBManager, taskRepository)
	taskController := taskController.NewController(taskHandler)

	app.UserController = userController
	app.TaskController = taskController
}

func (app *App) initializeRoutes() {
	app.Engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, api_response.BuildSuccessResponse("ina recruitment", 200, nil))
	})

	v1 := app.Engine.Group("/api/v1")
	authRoute := app.Engine.Group("/api/v1")
	authRoute.Use(middleware.NewMiddleware().JWTMiddleware()) // Apply the JWT middleware to v1 routes

	// user
	user := v1.Group("/users")
	user.POST("/login", app.UserController.LoginUserController)
	user.POST("", app.UserController.RegisterUser)
	user.GET("", app.UserController.GetUsers)
	user.GET("/:id", app.UserController.GetUserById)
	user.PUT("/:id", app.UserController.UpdateUser)
	user.DELETE("/:id", app.UserController.DeleteUser)

	v1.GET("/auth/google/login", app.UserController.OauthGoogleLogin)
	v1.GET("/auth/google/callback", app.UserController.OauthGoogleCallback)

	// task
	task := authRoute.Group("/task")
	task.POST("", app.TaskController.CreateTask)
	task.GET("", app.TaskController.GetTasks)
	task.GET("/:id", app.TaskController.GetTaskById)
	task.PUT("/:id", app.TaskController.UpdateTask)
	task.DELETE("/:id", app.TaskController.DeleteTask)
}

func (app *App) InitializeDatabase(dsn string, maxIdleConns int, maxOpenConns int) {
	err := app.DBManager.Initialize(dsn, maxIdleConns, maxOpenConns)

	if err != nil {
		panic(err)
	}
}
