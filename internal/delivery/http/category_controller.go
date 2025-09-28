package http

import (
	"pakyus_commerce/internal/model"
	"pakyus_commerce/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryController struct {
	Log     *logrus.Logger
	UseCase *usecase.CategoryUsecase
}

func NewCategoryController(useCase *usecase.CategoryUsecase, logger *logrus.Logger) *CategoryController {
	return &CategoryController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CategoryController) List(ctx *fiber.Ctx) error {
	response, err := c.UseCase.FindAll(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("error fetching categories")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return ctx.JSON(model.WebResponse[[]model.CategoryResponse]{Data: response})
}

func SetupCategoryRoute(app *fiber.App, controller *CategoryController) {
	app.Get("/api/categories", controller.List)
}
