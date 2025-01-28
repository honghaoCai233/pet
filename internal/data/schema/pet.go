package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 宠物类型枚举
const (
	PetTypeDog    = "dog"    // 狗
	PetTypeCat    = "cat"    // 猫
	PetTypeBird   = "bird"   // 鸟
	PetTypeFish   = "fish"   // 鱼
	PetTypeOthers = "others" // 其他
)

type Pet struct {
	ent.Schema
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id").
			Comment("主人ID"),
		field.String("name").
			NotEmpty().
			Comment("宠物名称"),
		field.String("type").
			Default(PetTypeDog).
			Comment("宠物类型").
			Validate(func(s string) error {
				switch s {
				case PetTypeDog, PetTypeCat, PetTypeBird, PetTypeFish, PetTypeOthers:
					return nil
				default:
					return fmt.Errorf("invalid pet type %q", s)
				}
			}),
		field.String("breed").
			Optional().
			Comment("宠物品种"),
		field.Int("age").
			Optional().
			Comment("宠物年龄"),
		field.Float("weight").
			Optional().
			Comment("宠物体重(kg)"),
		field.String("gender").
			Optional().
			Comment("宠物性别"),
		field.Text("description").
			Optional().
			Comment("宠物描述，包括性格特征、习惯等"),
		field.Text("care_instructions").
			Optional().
			Comment("照护说明，包括喂食要求、运动需求等"),
		field.Strings("photos").
			Optional().
			Comment("宠物照片URL列表"),
		field.Bool("vaccinated").
			Default(false).
			Comment("是否已接种疫苗"),
		field.Time("created_at").
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}
