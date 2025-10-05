package routers

import (
	"github.com/wallissonmarinho/cake/controllers"
	"github.com/wallissonmarinho/cake/database"
	"github.com/wallissonmarinho/cake/middlewares"
	"github.com/wallissonmarinho/cake/repositories"
	"github.com/wallissonmarinho/cake/services"

	"github.com/gofiber/fiber/v2"
)

type DashboardRouter struct {
	dashboardController controllers.DashboardController
}

func NewDashboardRouter() *DashboardRouter {
	// Initialize repositories
	dashboardRepository := repositories.NewDashboardRepository(database.DB)

	// Initialize services with repositories
	dashboardService := services.NewDashboardService(dashboardRepository)

	// Initialize controllers with services
	dashboardController := controllers.NewDashboardController(dashboardService)

	return &DashboardRouter{
		dashboardController: dashboardController,
	}
}

func (r *DashboardRouter) InstallRouters(app *fiber.App) {
	dashboard := app.Group("/dashboard").Use(middlewares.LoginAndStaffRequired())
	dashboard.Get("/", r.dashboardController.RenderDashboard)
}
