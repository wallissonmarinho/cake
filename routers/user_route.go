package routers

import (
	"github.com/wallissonmarinho/cake/controllers"
	"github.com/wallissonmarinho/cake/database"
	"github.com/wallissonmarinho/cake/middlewares"
	"github.com/wallissonmarinho/cake/repositories"
	"github.com/wallissonmarinho/cake/services"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	userController controllers.UserController
}

func NewUserRouter() *UserRouter {
	// Initialize repositories
	userRepository := repositories.NewUserRepository(database.DB)

	// Initialize services with repositories
	userService := services.NewUserService(userRepository)

	// Initialize controllers with services
	userController := controllers.NewUserController(userService)

	return &UserRouter{
		userController: userController,
	}
}

func (r *UserRouter) InstallRouters(app *fiber.App) {
	userGroup := app.Group("/users").Use(middlewares.LoginRequired())
	userGroup.Get("/user/:id", r.userController.RenderUser)
	userGroup.Post("/user/update/:id", r.userController.Update)

	adminGroup := app.Group("/users").Use(middlewares.LoginAndStaffRequired())
	adminGroup.Get("/create", r.userController.RenderCreate)
	adminGroup.Post("/create", r.userController.Create)
	adminGroup.Get("/delete/:id", r.userController.RenderDelete)
	adminGroup.Post("/delete/:id", r.userController.Delete)
	adminGroup.Get("/", r.userController.RenderUsers)

}
