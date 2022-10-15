package handlers

import (
	"github.com/Sup3r-Us3r/go-container-pattern/src/handlers/installment"
	"github.com/Sup3r-Us3r/go-container-pattern/src/services"
	"github.com/gofiber/fiber/v2"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(
		router,
		serviceContainer.InstallmentService,
	).SetRoutes()
}
