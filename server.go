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
	db                    *gorm.DB                         = config.SetupDatabaseConnection()
	userRepository        repository.UserRepository        = repository.NewUserRepository(db)
	accountRepository     repository.AccountRepository     = repository.NewAccountRepository(db)
	managerRepository     repository.ManagerRepository     = repository.NewManagerRepository(db)
	transactionRepository repository.TransactionRepository = repository.NewTransactionRepository(db)
	jwtService            service.JWTService               = service.NewJWTService()
	userService           service.UserService              = service.NewUserService(userRepository)
	accountService        service.AccountService           = service.NewAccountService(accountRepository)
	managerService        service.ManagerService           = service.NewManagerService(managerRepository)
	transactionService    service.TransactionService       = service.NewTransactionService(transactionRepository)
	authService           service.AuthService              = service.NewAuthService(userRepository)
	authController        controller.AuthController        = controller.NewAuthController(authService, jwtService)
	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	accountController     controller.AccountController     = controller.NewAccountController(accountService, jwtService)
	transactionController controller.TransactionController = controller.NewTransactionController(transactionService, accountService, jwtService)
	managerController     controller.ManagerController     = controller.NewManagerController(managerService, jwtService)
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
		authRoutes.POST("/register", authController.Register) // Create Client
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	accountRoutes := r.Group("api/account", middleware.AuthorizeJWT(jwtService))
	{
		accountRoutes.GET("/", accountController.All)         // Get all Accounts
		accountRoutes.POST("/", accountController.Insert)     // Create Account for Client
		accountRoutes.GET("/:id", accountController.FindByID) // Get Account by ID
		accountRoutes.PUT("/:id", accountController.Update)
		accountRoutes.DELETE("/:id", accountController.Delete)
	}

	passwordManagerRoutes := r.Group("api/manager", middleware.AuthorizeJWT(jwtService))
	{
		passwordManagerRoutes.GET("/", managerController.All)
		passwordManagerRoutes.POST("/", managerController.Insert)
		passwordManagerRoutes.PUT("/:id", managerController.Update)
		passwordManagerRoutes.DELETE("/:id", managerController.Delete)
	}

	transactionRoutes := r.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.GET("/:id", transactionController.GetAllTransactionByAccountID) // Get all Transactions by Account ID
		transactionRoutes.POST("/", transactionController.CreateTransaction)              // Create Transaction
	}

	r.Run(":8080")
}
