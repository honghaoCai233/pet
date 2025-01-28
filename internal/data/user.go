package data

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"pet/internal/data/ent"
	"pet/internal/data/ent/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

// 生成随机盐值
func generateSalt() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// 密码加密
func hashPassword(password, salt string) (string, error) {
	// 将密码和盐值组合
	combined := password + salt
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Register 用户注册
func (r *UserRepo) Register(ctx context.Context, u *ent.User) (*ent.User, error) {
	// 检查用户名是否已存在
	exists, err := r.data.db.User.Query().
		Where(user.Username(u.Username)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("username already exists")
	}

	// 如果提供了邮箱，检查邮箱是否已存在
	if u.Email != "" {
		exists, err = r.data.db.User.Query().
			Where(user.Email(u.Email)).
			Exist(ctx)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, fmt.Errorf("email already exists")
		}
	}

	// 生成盐值
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}

	// 密码加密
	hashedPassword, err := hashPassword(u.Password, salt)
	if err != nil {
		return nil, err
	}

	// 创建用户
	return r.data.db.User.Create().
		SetUsername(u.Username).
		SetPassword(hashedPassword).
		SetSalt(salt).
		SetEmail(u.Email).
		SetName(u.Name).
		SetPhone(u.Phone).
		SetAddress(u.Address).
		SetAge(u.Age).
		SetRole(u.Role).
		SetStatus("active").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}

// Login 用户登录
func (r *UserRepo) Login(ctx context.Context, username, password string) (*ent.User, error) {
	// 查找用户
	u, err := r.data.db.User.Query().
		Where(
			user.Or(
				user.Username(username),
				user.Email(username),
			),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	// 检查用户状态
	if u.Status != "active" {
		return nil, fmt.Errorf("account is %s", u.Status)
	}

	// 验证密码
	combined := password + u.Salt
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(combined))
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	// 更新最后登录时间
	return r.data.db.User.UpdateOne(u).
		SetLastLoginAt(time.Now()).
		Save(ctx)
}

// ChangePassword 修改密码
func (r *UserRepo) ChangePassword(ctx context.Context, userID int, oldPassword, newPassword string) error {
	u, err := r.data.db.User.Get(ctx, userID)
	if err != nil {
		return err
	}

	// 验证旧密码
	combined := oldPassword + u.Salt
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(combined))
	if err != nil {
		return fmt.Errorf("invalid old password")
	}

	// 生成新的盐值和密码哈希
	salt, err := generateSalt()
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(newPassword, salt)
	if err != nil {
		return err
	}

	// 更新密码
	return r.data.db.User.UpdateOne(u).
		SetPassword(hashedPassword).
		SetSalt(salt).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
}

// UpdateStatus 更新用户状态
func (r *UserRepo) UpdateStatus(ctx context.Context, userID int, status string) error {
	return r.data.db.User.UpdateOneID(userID).
		SetStatus(status).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
}

// Update 更新用户信息
func (r *UserRepo) Update(ctx context.Context, u *ent.User) (*ent.User, error) {
	builder := r.data.db.User.UpdateOneID(u.ID).
		SetName(u.Name).
		SetAddress(u.Address).
		SetPhone(u.Phone).
		SetAge(u.Age).
		SetRole(u.Role).
		SetUpdatedAt(time.Now())

	if u.Description != "" {
		builder.SetDescription(u.Description)
	}
	if u.Rating != 0 {
		builder.SetRating(u.Rating)
	}
	if u.CompletedTasks != 0 {
		builder.SetCompletedTasks(u.CompletedTasks)
	}
	if u.Email != "" {
		builder.SetEmail(u.Email)
	}

	return builder.Save(ctx)
}

// GetByUsername 通过用户名查找用户
func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	return r.data.db.User.Query().
		Where(user.Username(username)).
		Only(ctx)
}

// GetByEmail 通过邮箱查找用户
func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.data.db.User.Query().
		Where(user.Email(email)).
		Only(ctx)
}

// 保留原有的其他方法...
func (r *UserRepo) Get(ctx context.Context, id int) (*ent.User, error) {
	return r.data.db.User.Get(ctx, id)
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.User.DeleteOneID(id).Exec(ctx)
}

func (r *UserRepo) List(ctx context.Context, page, pageSize int) ([]*ent.User, error) {
	offset := (page - 1) * pageSize
	return r.data.db.User.Query().
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

func (r *UserRepo) GetByPhone(ctx context.Context, phone string) (*ent.User, error) {
	return r.data.db.User.Query().
		Where(user.PhoneEQ(phone)).
		Only(ctx)
}
