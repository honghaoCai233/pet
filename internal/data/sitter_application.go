package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/data/ent/sitterapplication"
)

type SitterApplicationRepo struct {
	data *Data
}

func NewSitterApplicationRepo(data *Data) *SitterApplicationRepo {
	return &SitterApplicationRepo{
		data: data,
	}
}

// Create 创建申请
func (r *SitterApplicationRepo) Create(ctx context.Context, application *ent.SitterApplication) (*ent.SitterApplication, error) {
	return r.data.db.SitterApplication.
		Create().
		SetSitterID(application.SitterID).
		SetOwnerID(application.OwnerID).
		SetPetID(application.PetID).
		SetStatus(application.Status).
		SetIntroduction(application.Introduction).
		SetExperience(application.Experience).
		SetAvailability(application.Availability).
		SetExpectedSalary(application.ExpectedSalary).
		Save(ctx)
}

// Update 更新申请
func (r *SitterApplicationRepo) Update(ctx context.Context, application *ent.SitterApplication) (*ent.SitterApplication, error) {
	return r.data.db.SitterApplication.
		UpdateOneID(application.ID).
		SetStatus(application.Status).
		SetIntroduction(application.Introduction).
		SetExperience(application.Experience).
		SetAvailability(application.Availability).
		SetExpectedSalary(application.ExpectedSalary).
		SetNillableRejectReason(&application.RejectReason).
		Save(ctx)
}

// Get 获取申请
func (r *SitterApplicationRepo) Get(ctx context.Context, id int) (*ent.SitterApplication, error) {
	return r.data.db.SitterApplication.Get(ctx, id)
}

// Delete 删除申请
func (r *SitterApplicationRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.SitterApplication.DeleteOneID(id).Exec(ctx)
}

// List 获取申请列表
func (r *SitterApplicationRepo) List(ctx context.Context, page, pageSize int) ([]*ent.SitterApplication, error) {
	offset := (page - 1) * pageSize
	return r.data.db.SitterApplication.
		Query().
		Offset(offset).
		Limit(pageSize).
		Order(ent.Desc(sitterapplication.FieldCreatedAt)).
		All(ctx)
}

// ListBySitter 获取照护者的申请列表
func (r *SitterApplicationRepo) ListBySitter(ctx context.Context, sitterID int) ([]*ent.SitterApplication, error) {
	return r.data.db.SitterApplication.
		Query().
		Where(sitterapplication.SitterID(sitterID)).
		Order(ent.Desc(sitterapplication.FieldCreatedAt)).
		All(ctx)
}

// ListByOwner 获取宠物主人收到的申请列表
func (r *SitterApplicationRepo) ListByOwner(ctx context.Context, ownerID int) ([]*ent.SitterApplication, error) {
	return r.data.db.SitterApplication.
		Query().
		Where(sitterapplication.OwnerID(ownerID)).
		Order(ent.Desc(sitterapplication.FieldCreatedAt)).
		All(ctx)
}

// ListByPet 获取宠物的申请列表
func (r *SitterApplicationRepo) ListByPet(ctx context.Context, petID int) ([]*ent.SitterApplication, error) {
	return r.data.db.SitterApplication.
		Query().
		Where(sitterapplication.PetID(petID)).
		Order(ent.Desc(sitterapplication.FieldCreatedAt)).
		All(ctx)
}

// ListByStatus 获取指定状态的申请列表
func (r *SitterApplicationRepo) ListByStatus(ctx context.Context, status string, page, pageSize int) ([]*ent.SitterApplication, error) {
	offset := (page - 1) * pageSize
	return r.data.db.SitterApplication.
		Query().
		Where(sitterapplication.Status(status)).
		Offset(offset).
		Limit(pageSize).
		Order(ent.Desc(sitterapplication.FieldCreatedAt)).
		All(ctx)
}

// UpdateStatus 更新申请状态
func (r *SitterApplicationRepo) UpdateStatus(ctx context.Context, id int, status string, rejectReason *string) error {
	update := r.data.db.SitterApplication.
		UpdateOneID(id).
		SetStatus(status)

	if rejectReason != nil {
		update.SetRejectReason(*rejectReason)
	}

	_, err := update.Save(ctx)
	return err
}
