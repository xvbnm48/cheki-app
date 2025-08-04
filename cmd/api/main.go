package main

import (
	_ "chekisvc/docs"
	"chekisvc/internal/delivery/http/handler"
	"chekisvc/internal/infrastructure/config"
	"chekisvc/internal/infrastructure/database"
	user_repository "chekisvc/internal/infrastructure/database"
	"chekisvc/internal/usecase"
	"chekisvc/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.NewLogger()

	// Initialize database
	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := user_repository.NewUserRepository(db.DB)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(userRepo, log)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)

	// Setup router
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)
		v1.GET("/users", userHandler.ListUsers)

	}

	// Start server
	log.Info("Server starting on port: ", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
