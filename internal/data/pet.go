package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/data/ent/pet"
)

type PetRepo struct {
	data *Data
}

// NewPetRepo .
func NewPetRepo(data *Data) *PetRepo {
	return &PetRepo{
		data: data,
	}
}

func (r *PetRepo) Create(ctx context.Context, p *ent.Pet) (*ent.Pet, error) {
	builder := r.data.db.Pet.Create().
		SetOwnerID(p.OwnerID).
		SetName(p.Name).
		SetType(p.Type)

	if p.Breed != "" {
		builder.SetBreed(p.Breed)
	}
	if p.Age != 0 {
		builder.SetAge(p.Age)
	}
	if p.Weight != 0 {
		builder.SetWeight(p.Weight)
	}
	if p.Gender != "" {
		builder.SetGender(p.Gender)
	}
	if p.Description != "" {
		builder.SetDescription(p.Description)
	}
	if p.CareInstructions != "" {
		builder.SetCareInstructions(p.CareInstructions)
	}
	if len(p.Photos) > 0 {
		builder.SetPhotos(p.Photos)
	}
	builder.SetVaccinated(p.Vaccinated)

	return builder.Save(ctx)
}

func (r *PetRepo) Update(ctx context.Context, p *ent.Pet) (*ent.Pet, error) {
	builder := r.data.db.Pet.UpdateOneID(p.ID).
		SetOwnerID(p.OwnerID).
		SetName(p.Name).
		SetType(p.Type)

	if p.Breed != "" {
		builder.SetBreed(p.Breed)
	}
	if p.Age != 0 {
		builder.SetAge(p.Age)
	}
	if p.Weight != 0 {
		builder.SetWeight(p.Weight)
	}
	if p.Gender != "" {
		builder.SetGender(p.Gender)
	}
	if p.Description != "" {
		builder.SetDescription(p.Description)
	}
	if p.CareInstructions != "" {
		builder.SetCareInstructions(p.CareInstructions)
	}
	if len(p.Photos) > 0 {
		builder.SetPhotos(p.Photos)
	}
	builder.SetVaccinated(p.Vaccinated)

	return builder.Save(ctx)
}

func (r *PetRepo) Get(ctx context.Context, id int) (*ent.Pet, error) {
	return r.data.db.Pet.Get(ctx, id)
}

func (r *PetRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Pet.DeleteOneID(id).Exec(ctx)
}

func (r *PetRepo) List(ctx context.Context, page, pageSize int) ([]*ent.Pet, error) {
	offset := (page - 1) * pageSize
	return r.data.db.Pet.Query().
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

func (r *PetRepo) ListByOwner(ctx context.Context, ownerID int) ([]*ent.Pet, error) {
	return r.data.db.Pet.Query().
		Where(pet.OwnerIDEQ(ownerID)).
		All(ctx)
}
