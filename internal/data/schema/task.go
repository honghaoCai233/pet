package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 任务状态枚举
const (
	TaskStatusOpen      = "open"        // 开放中
	TaskStatusProgress  = "in_progress" // 进行中
	TaskStatusCompleted = "completed"   // 已完成
	TaskStatusCanceled  = "canceled"    // 已取消
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("publisher_id").
			Comment("发布者ID"),
		field.Int("pet_id").
			Comment("宠物ID"),
		field.Int("sitter_id").
			Optional().
			Comment("接单者ID"),
		field.String("title").
			NotEmpty().
			Comment("任务标题"),
		field.Text("description").
			NotEmpty().
			Comment("任务描述"),
		field.Float("reward").
			Min(0).
			Comment("任务酬劳"),
		field.Time("start_time").
			Comment("任务开始时间"),
		field.Time("end_time").
			Comment("任务结束时间"),
		field.String("status").
			Default(TaskStatusOpen).
			Comment("任务状态").
			Validate(func(s string) error {
				switch s {
				case TaskStatusOpen, TaskStatusProgress, TaskStatusCompleted, TaskStatusCanceled:
					return nil
				default:
					return fmt.Errorf("invalid task status %q", s)
				}
			}),
		field.String("location").
			NotEmpty().
			Comment("任务地点"),
		field.Text("requirements").
			Optional().
			Comment("特殊要求"),
		field.Int("visits_count").
			Default(0).
			Min(1).
			Comment("需要上门次数"),
		field.Text("care_instructions").
			Optional().
			Comment("照护说明"),
		field.Time("created_at").
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}
