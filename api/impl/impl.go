package impl

import (
	"github.com/integrational/apitests/testapi2/api/gen"
	"github.com/labstack/echo/v4"
)

type APIImpl struct {
	svc *Petstore
}

func New() *APIImpl {
	return &APIImpl{NewPetstore()}
}

func (impl *APIImpl) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	defer un(trace("APIImpl.FindPets"))
	if pets, err := impl.svc.FindPets(params); err != nil {
		return err
	} else {
		return ctx.JSON(200, pets)
	}
}

func (impl *APIImpl) AddPet(ctx echo.Context) (err error) {
	defer un(trace("APIImpl.AddPet"))
	newPet := gen.NewPet{}
	if err = ctx.Bind(&newPet); err != nil {
		return err
	}
	if pet, err := impl.svc.AddPet(&newPet); err != nil {
		return err
	} else {
		return ctx.JSON(200, pet)
	}
}

func (impl *APIImpl) FindPetById(ctx echo.Context, id int64) error {
	defer un(trace("APIImpl.FindPetById"))
	if pet, err := impl.svc.FindPetById(id); err != nil {
		return err
	} else {
		return ctx.JSON(200, pet)
	}
}

func (impl *APIImpl) DeletePet(ctx echo.Context, id int64) error {
	defer un(trace("APIImpl.DeletePet"))
	if err := impl.svc.DeletePet(id); err != nil {
		return err
	} else {
		return ctx.NoContent(204)
	}
}
