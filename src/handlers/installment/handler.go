package installment

import (
	"fmt"
	"net/http"

	"github.com/Sup3r-Us3r/go-container-pattern/src/interfaces"
	"github.com/Sup3r-Us3r/go-container-pattern/src/structs"
	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	Router             fiber.Router
	InstallmentService interfaces.InstallmentService
}

func NewInstallmentHandler(
	router fiber.Router,
	installmentService interfaces.InstallmentService,
) *InstallmentHandler {
	return &InstallmentHandler{
		Router:             router,
		InstallmentService: installmentService,
	}
}

func (ih *InstallmentHandler) SetRoutes() {
	group := ih.Router.Group("/installment")

	group.Post("/", ih.CreateInstallment)
}

func (ih *InstallmentHandler) CreateInstallment(context *fiber.Ctx) error {
	var installmentBody structs.Installment

	err := context.BodyParser(&installmentBody)

	if err != nil {
		fmt.Println(err.Error())
		return context.Status(http.StatusBadRequest).JSON(
			structs.Response{
				Data: err.Error(),
				Tag:  "BAD_REQUEST",
			},
		)
	}

	err = ih.InstallmentService.Create(installmentBody)

	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(
			structs.Response{
				Data: err.Error(),
				Tag:  "INTERNAL_SERVER_ERROR",
			},
		)
	}

	return context.SendStatus(http.StatusCreated)
}
