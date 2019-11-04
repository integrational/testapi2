package impl

import (
	"github.com/integrational/apitests/testapi2/api/gen"
	"log"
	"sort"
	"sync"
)

const (
	maxNumResults = 100
)

type Petstore struct {
	mutex    sync.Mutex
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
	return &Petstore{petsById: petsById, nextId: 4}
}

func (svc *Petstore) lock() *Petstore {
	svc.mutex.Lock()
	return svc
}

func (svc *Petstore) unlock() {
	svc.mutex.Unlock()
}

func (svc *Petstore) FindPets(params gen.FindPetsParams) (pets []gen.Pet, err error) {
	defer un(trace("Petstore.FindPets"))

	tags := sortTags(params)
	limit := limit(params)
	log.Println("Filtering by: tags", tags, ", limit", limit)

	pets = make([]gen.Pet, 0, maxNumResults)
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
		return maxNumResults // math.MaxInt32
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
	defer svc.lock().unlock()

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
		return nil, nil // not reasonable to return error if no matching pet
	}
}

func (svc *Petstore) DeletePet(id int64) error {
	defer un(trace("Petstore.DeletePet"))
	defer svc.lock().unlock()

	delete(svc.petsById, id)
	return nil
}
