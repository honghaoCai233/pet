package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SitterApplication 宠物照护申请表
type SitterApplication struct {
	ent.Schema
}

// Fields 字段定义
func (SitterApplication) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sitter_id").
			Comment("申请人ID").
			Positive(),
		field.Int("owner_id").
			Comment("宠物主人ID").
			Positive(),
		field.Int("pet_id").
			Comment("宠物ID").
			Positive(),
		field.String("status").
			Comment("申请状态: pending-待处理 approved-已通过 rejected-已拒绝").
			Default("pending"),
		field.String("introduction").
			Comment("自我介绍").
			Optional(),
		field.String("experience").
			Comment("照护经验").
			Optional(),
		field.String("availability").
			Comment("可用时间").
			Optional(),
		field.Float("expected_salary").
			Comment("期望薪资").
			Optional(),
		field.String("reject_reason").
			Comment("拒绝原因").
			Optional(),
		field.Time("created_at").
			Comment("创建时间").
			Immutable().
			Optional(),
		field.Time("updated_at").
			Comment("更新时间").
			Optional(),
	}
}
