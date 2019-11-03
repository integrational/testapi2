package impl

import (
	"github.com/integrational/apitests/testapi2/api/gen"
	"github.com/labstack/echo/v4"
	"log"
)

type APIImpl struct {
	pets map[int64]gen.Pet
}

func New() *APIImpl {
	pets := make(map[int64]gen.Pet)
	return &APIImpl{pets: pets}
}

func (impl *APIImpl) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	defer un(trace("FindPets"))
	return ctx.JSON(200, nil)
}

func (impl *APIImpl) AddPet(ctx echo.Context) error {
	defer un(trace("AddPet"))
	return ctx.JSON(200, nil)
}

func (impl *APIImpl) DeletePet(ctx echo.Context, id int64) error {
	defer un(trace("DeletePet"))
	return ctx.JSON(200, nil)
}

func (impl *APIImpl) FindPetById(ctx echo.Context, id int64) error {
	defer un(trace("FindPetById"))
	return ctx.JSON(200, nil)
}

func trace(s string) string {
	log.Println("=>", s)
	return s
}

func un(s string) {
	log.Println("<=", s)
}
