package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 用户角色枚举
const (
	RoleTaskPublisher = "task_publisher" // 发布任务的用户
	RolePetSitter     = "pet_sitter"     // 宠物照看者
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("address"),
		field.String("phone"),
		field.Int("age"),
		field.String("role").
			Default(RoleTaskPublisher).
			Optional().
			Comment("用户角色: task_publisher-发布任务者, pet_sitter-宠物照看者").
			Validate(func(s string) error {
				switch s {
				case RoleTaskPublisher, RolePetSitter:
					return nil
				default:
					return fmt.Errorf("invalid role type %q", s)
				}
			}),
		field.Text("description").
			Optional().
			Comment("用户描述，对于宠物照看者可以描述自己的经验和专长"),
		field.Float("rating").
			Optional().
			Default(5.0).
			Comment("用户评分，主要用于评价宠物照看者的服务质量"),
		field.Int("completed_tasks").
			Optional().
			Default(0).
			Comment("完成的任务数量"),
		field.String("username").
			Unique().
			NotEmpty().
			Comment("用户名，用于登录"),
		field.String("email").
			Unique().
			Optional().
			Comment("邮箱，可用于登录和找回密码"),
		field.String("password").
			Sensitive().
			NotEmpty().
			Comment("密码哈希值"),
		field.String("salt").
			Sensitive().
			Comment("密码加密盐值"),
		field.Time("last_login_at").
			Optional().
			Comment("最后登录时间"),
		field.String("status").
			Default("active").
			Comment("用户状态：active-正常, disabled-禁用, locked-锁定").
			Validate(func(s string) error {
				switch s {
				case "active", "disabled", "locked":
					return nil
				default:
					return fmt.Errorf("invalid status type %q", s)
				}
			}),
		field.Time("created_at").
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}
