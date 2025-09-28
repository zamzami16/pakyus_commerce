package http

import (
	"pakyus_commerce/internal/delivery/http/middleware"
	"pakyus_commerce/internal/model"
	"pakyus_commerce/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AddressController struct {
	UseCase *usecase.AddressUseCase
	Log     *logrus.Logger
}

func NewAddressController(useCase *usecase.AddressUseCase, log *logrus.Logger) *AddressController {
	return &AddressController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *AddressController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	contactId, err := uuid.Parse(ctx.Params("contactId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid contact ID format")
		return fiber.ErrBadRequest
	}

	request := new(model.CreateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = contactId

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

func (c *AddressController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	contactId, err := uuid.Parse(ctx.Params("contactId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid contact ID format")
		return fiber.ErrBadRequest
	}

	request := &model.ListAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list addresses")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.AddressResponse]{Data: responses})
}

func (c *AddressController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	contactId, err := uuid.Parse(ctx.Params("contactId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid contact ID format")
		return fiber.ErrBadRequest
	}

	addressId, err := uuid.Parse(ctx.Params("addressId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid address ID format")
		return fiber.ErrBadRequest
	}

	request := &model.GetAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        addressId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

func (c *AddressController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	contactId, err := uuid.Parse(ctx.Params("contactId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid contact ID format")
		return fiber.ErrBadRequest
	}

	addressId, err := uuid.Parse(ctx.Params("addressId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid address ID format")
		return fiber.ErrBadRequest
	}

	request := new(model.UpdateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = contactId
	request.ID = addressId

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	contactId, err := uuid.Parse(ctx.Params("contactId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid contact ID format")
		return fiber.ErrBadRequest
	}

	addressId, err := uuid.Parse(ctx.Params("addressId"))
	if err != nil {
		c.Log.WithError(err).Error("invalid address ID format")
		return fiber.ErrBadRequest
	}

	request := &model.DeleteAddressRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        addressId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
