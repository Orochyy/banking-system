package main

import (
	"banking-system/config"
	"banking-system/controller"
	"banking-system/middleware"
	"banking-system/repository"
	"banking-system/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	accountRepository repository.AccountRepository = repository.NewAccountRepository(db)
	jwtService        service.JWTService           = service.NewJWTService()
	userService       service.UserService          = service.NewUserService(userRepository)
	accountService    service.AccountService       = service.NewAccountService(accountRepository)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	accountController controller.AccountController = controller.NewAccountController(accountService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	accountRoutes := r.Group("api/account", middleware.AuthorizeJWT(jwtService))
	{

		accountRoutes.GET("/", accountController.All)
		accountRoutes.POST("/", accountController.Insert)
		accountRoutes.GET("/:id", accountController.FindByID)
		accountRoutes.PUT("/:id", accountController.Update)
		accountRoutes.DELETE("/:id", accountController.Delete)
	}

	r.Run(":8080")
}
