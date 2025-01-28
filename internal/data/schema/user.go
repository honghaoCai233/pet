package schema

import (
	"fmt"

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
	}
}
