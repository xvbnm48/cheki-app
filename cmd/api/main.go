// @title           Cheki Service API
// @version         1.0
// @description     A Clean Architecture Go API with JWT Authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	_ "chekisvc/docs"
	"chekisvc/internal/delivery/http/handler"
	"chekisvc/internal/delivery/http/middleware"
	"chekisvc/internal/infrastructure/config"
	"chekisvc/internal/infrastructure/database"

	"chekisvc/internal/usecase/product"
	usecase "chekisvc/internal/usecase/user"
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
	userRepo := database.NewUserRepository(db.DB)
	productRepo := database.NewProductRepository(db.DB)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(userRepo, log)
	productUsecase := product.NewProductUsecase(productRepo, log)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(userUsecase)
	productHandler := handler.NewProductHandler(productUsecase)

	// Setup router
	router := gin.Default()
	router.Use(middleware.LoadUpperLower())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public endpoints (no authentication required)
	public := router.Group("/api/v1")
	{
		public.POST("/login", authHandler.Login)
		public.POST("/refresh-token", authHandler.RefreshToken)
		public.POST("/users", userHandler.CreateUser) // Register user
	}

	// Protected endpoints (JWT authentication required)
	protected := router.Group("/api/v1")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/users/:id", userHandler.GetUser)
		protected.PUT("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
		protected.GET("/users", userHandler.ListUsers)
		protected.POST("/products", productHandler.CreateProduct)
		protected.GET("/products", productHandler.GetAllProducts)
		// protected.PUT("/products/:id", productHandler.UpdateProduct)
		// protected.DELETE("/products/:id", productHandler.DeleteProduct)
		// TODO: Implement ListProducts endpoint in the future
	}

	// Start server
	log.Info("Server starting on port: ", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
