package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/data/ent/community"
)

type CommunityRepo struct {
	data *Data
}

// NewCommunityRepo .
func NewCommunityRepo(data *Data) *CommunityRepo {
	return &CommunityRepo{
		data: data,
	}
}

func (r *CommunityRepo) Create(ctx context.Context, c *ent.Community) (*ent.Community, error) {
	builder := r.data.db.Community.Create().
		SetAuthorID(c.AuthorID).
		SetTitle(c.Title).
		SetType(c.Type).
		SetContent(c.Content).
		SetIsPinned(c.IsPinned)

	if c.PetID != 0 {
		builder.SetPetID(c.PetID)
	}
	if len(c.Images) > 0 {
		builder.SetImages(c.Images)
	}
	if len(c.Tags) > 0 {
		builder.SetTags(c.Tags)
	}

	return builder.Save(ctx)
}

func (r *CommunityRepo) Update(ctx context.Context, c *ent.Community) (*ent.Community, error) {
	builder := r.data.db.Community.UpdateOneID(c.ID).
		SetAuthorID(c.AuthorID).
		SetTitle(c.Title).
		SetType(c.Type).
		SetContent(c.Content).
		SetIsPinned(c.IsPinned)

	if c.PetID != 0 {
		builder.SetPetID(c.PetID)
	}
	if len(c.Images) > 0 {
		builder.SetImages(c.Images)
	}
	if len(c.Tags) > 0 {
		builder.SetTags(c.Tags)
	}

	return builder.Save(ctx)
}

func (r *CommunityRepo) Get(ctx context.Context, id int) (*ent.Community, error) {
	return r.data.db.Community.Get(ctx, id)
}

func (r *CommunityRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Community.DeleteOneID(id).Exec(ctx)
}

func (r *CommunityRepo) List(ctx context.Context, page, pageSize int) ([]*ent.Community, error) {
	offset := (page - 1) * pageSize
	return r.data.db.Community.Query().
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

func (r *CommunityRepo) ListByAuthor(ctx context.Context, authorID int) ([]*ent.Community, error) {
	return r.data.db.Community.Query().
		Where(community.AuthorIDEQ(authorID)).
		All(ctx)
}

func (r *CommunityRepo) ListByPet(ctx context.Context, petID int) ([]*ent.Community, error) {
	return r.data.db.Community.Query().
		Where(community.PetIDEQ(petID)).
		All(ctx)
}

func (r *CommunityRepo) ListByType(ctx context.Context, postType string, page, pageSize int) ([]*ent.Community, error) {
	offset := (page - 1) * pageSize
	return r.data.db.Community.Query().
		Where(community.TypeEQ(postType)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

func (r *CommunityRepo) UpdateLikes(ctx context.Context, id int, increment bool) error {
	post, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	likes := post.Likes
	if increment {
		likes++
	} else {
		likes--
	}

	return r.data.db.Community.UpdateOneID(id).
		SetLikes(likes).
		Exec(ctx)
}

func (r *CommunityRepo) UpdateViews(ctx context.Context, id int) error {
	post, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	return r.data.db.Community.UpdateOneID(id).
		SetViews(post.Views + 1).
		Exec(ctx)
}
