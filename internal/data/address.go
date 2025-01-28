package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/data/ent/address"
)

// AddressRepo 地址仓储接口实现
type AddressRepo struct {
	data *Data
}

// NewAddressRepo 创建地址仓储实例
func NewAddressRepo(data *Data) *AddressRepo {
	return &AddressRepo{
		data: data,
	}
}

// Create 创建新地址
// 参数:
//   - ctx: 上下文
//   - a: 地址信息
//
// 返回:
//   - *ent.Address: 创建的地址实体
//   - error: 错误信息
func (r *AddressRepo) Create(ctx context.Context, a *ent.Address) (*ent.Address, error) {
	builder := r.data.db.Address.Create().
		SetName(a.Name).
		SetProvince(a.Province).
		SetCity(a.City).
		SetDistrict(a.District).
		SetStreet(a.Street).
		SetDetailedInfo(a.DetailedInfo).
		SetContactName(a.ContactName).
		SetContactPhone(a.ContactPhone).
		SetIsDefault(a.IsDefault)

	if a.UserID != 0 {
		builder.SetUserID(a.UserID)
	}

	return builder.Save(ctx)
}

// Update 更新地址信息
// 参数:
//   - ctx: 上下文
//   - a: 更新的地址信息
//
// 返回:
//   - *ent.Address: 更新后的地址实体
//   - error: 错误信息
func (r *AddressRepo) Update(ctx context.Context, a *ent.Address) (*ent.Address, error) {
	builder := r.data.db.Address.UpdateOneID(a.ID).
		SetName(a.Name).
		SetProvince(a.Province).
		SetCity(a.City).
		SetDistrict(a.District).
		SetStreet(a.Street).
		SetDetailedInfo(a.DetailedInfo).
		SetContactName(a.ContactName).
		SetContactPhone(a.ContactPhone).
		SetIsDefault(a.IsDefault)

	if a.UserID != 0 {
		builder.SetUserID(a.UserID)
	}

	return builder.Save(ctx)
}

// Get 获取指定ID的地址
// 参数:
//   - ctx: 上下文
//   - id: 地址ID
//
// 返回:
//   - *ent.Address: 地址实体
//   - error: 错误信息
func (r *AddressRepo) Get(ctx context.Context, id int) (*ent.Address, error) {
	return r.data.db.Address.Get(ctx, id)
}

// Delete 删除指定ID的地址
// 参数:
//   - ctx: 上下文
//   - id: 地址ID
//
// 返回:
//   - error: 错误信息
func (r *AddressRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Address.DeleteOneID(id).Exec(ctx)
}

// List 分页获取地址列表
// 参数:
//   - ctx: 上下文
//   - page: 页码
//   - pageSize: 每页数量
//
// 返回:
//   - []*ent.Address: 地址列表
//   - error: 错误信息
func (r *AddressRepo) List(ctx context.Context, page, pageSize int) ([]*ent.Address, error) {
	offset := (page - 1) * pageSize
	return r.data.db.Address.Query().
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

// ListByUser 获取指定用户的所有地址
// 参数:
//   - ctx: 上下文
//   - userID: 用户ID
//
// 返回:
//   - []*ent.Address: 地址列表
//   - error: 错误信息
func (r *AddressRepo) ListByUser(ctx context.Context, userID int) ([]*ent.Address, error) {
	return r.data.db.Address.Query().
		Where(address.UserIDEQ(userID)).
		All(ctx)
}

// GetDefaultAddress 获取用户的默认地址
// 参数:
//   - ctx: 上下文
//   - userID: 用户ID
//
// 返回:
//   - *ent.Address: 默认地址实体
//   - error: 错误信息
func (r *AddressRepo) GetDefaultAddress(ctx context.Context, userID int) (*ent.Address, error) {
	return r.data.db.Address.Query().
		Where(address.UserIDEQ(userID)).
		Where(address.IsDefaultEQ(true)).
		Only(ctx)
}

// SetDefault 设置指定地址为默认地址
// 该操作会将用户的其他地址设置为非默认
// 参数:
//   - ctx: 上下文
//   - id: 要设置为默认的地址ID
//   - userID: 用户ID
//
// 返回:
//   - error: 错误信息
func (r *AddressRepo) SetDefault(ctx context.Context, id int, userID int) error {
	// 先将该用户的所有地址设置为非默认
	_, err := r.data.db.Address.Update().
		Where(address.UserIDEQ(userID)).
		SetIsDefault(false).
		Save(ctx)
	if err != nil {
		return err
	}

	// 将指定地址设置为默认
	return r.data.db.Address.UpdateOneID(id).
		SetIsDefault(true).
		Exec(ctx)
}
