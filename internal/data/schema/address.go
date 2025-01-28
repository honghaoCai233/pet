package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Address 地址信息表
type Address struct {
	ent.Schema
}

// Fields 字段定义
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Optional().
			Comment("关联的用户ID，可选"),
		field.String("name").
			NotEmpty().
			Comment("地址名称，如'家'、'公司'等"),
		field.String("province").
			NotEmpty().
			Comment("省份"),
		field.String("city").
			NotEmpty().
			Comment("城市"),
		field.String("district").
			NotEmpty().
			Comment("区/县"),
		field.String("street").
			NotEmpty().
			Comment("街道地址"),
		field.String("detailed_info").
			NotEmpty().
			Comment("详细地址，如门牌号等"),
		field.String("contact_name").
			NotEmpty().
			Comment("联系人姓名"),
		field.String("contact_phone").
			NotEmpty().
			Comment("联系人电话"),
		field.Bool("is_default").
			Default(false).
			Comment("是否默认地址"),
		field.Time("created_at").
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}
