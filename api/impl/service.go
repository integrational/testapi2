package impl

import (
	"errors"
	"github.com/integrational/apitests/testapi2/api/gen"
	"log"
	"math"
	"sort"
)

type Petstore struct {
	// FIXME mutex
	petsById map[int64]gen.Pet
	nextId   int64
}

func NewPetstore() *Petstore {
	cat := "cat"
	dog := "dog"
	petsById := map[int64]gen.Pet{
		1: {Id: 1, NewPet: gen.NewPet{"Max", &cat}},
		2: {Id: 2, NewPet: gen.NewPet{"Moritz", &cat}},
		3: {Id: 3, NewPet: gen.NewPet{"Muppi", &dog}},
	}
	return &Petstore{petsById, 4}
}

func (svc *Petstore) FindPets(params gen.FindPetsParams) (pets []gen.Pet, err error) {
	defer un(trace("Petstore.FindPets"))

	tags := sortTags(params)
	limit := limit(params)
	log.Println("Filtering by: tags", tags, ", limit", limit)

	pets = make([]gen.Pet, 0, len(svc.petsById)) // FIXME can get too large
	for _, pet := range svc.petsById {
		if matchesAnyTag(pet, tags) {
			pets = append(pets, pet)
			if len(pets) >= limit {
				return
			}
		}
	}
	return
}

func limit(params gen.FindPetsParams) int {
	if params.Limit == nil {
		return math.MaxInt32
	}
	return int(*params.Limit)
}

func sortTags(params gen.FindPetsParams) sort.StringSlice {
	if params.Tags == nil {
		return []string{}
	}
	tags := make([]string, len(*params.Tags))
	copy(tags, *params.Tags)
	sort.Strings(tags)
	return tags
}

func matchesAnyTag(pet gen.Pet, tags sort.StringSlice) bool {
	switch {
	case tags.Len() == 0:
		return true
	case pet.Tag == nil:
		return false
	default:
		i := tags.Search(*pet.Tag)
		return i < tags.Len() && tags[i] == *pet.Tag
	}
}

func (svc *Petstore) AddPet(newPet *gen.NewPet) (pet *gen.Pet, err error) {
	defer un(trace("Petstore.AddPet"))
	id := svc.nextId
	svc.nextId++
	pet = &gen.Pet{Id: id, NewPet: *newPet}
	svc.petsById[id] = *pet
	return
}

func (svc *Petstore) FindPetById(id int64) (*gen.Pet, error) {
	defer un(trace("Petstore.FindPetById"))
	if pet, ok := svc.petsById[id]; ok {
		return &pet, nil
	} else {
		return nil, errors.New("no such ID") // FIXME
	}
}

func (svc *Petstore) DeletePet(id int64) error {
	defer un(trace("Petstore.DeletePet"))
	delete(svc.petsById, id)
	return nil
}
