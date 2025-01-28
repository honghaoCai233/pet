package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 帖子类型枚举
const (
	PostTypeMoment  = "moment"  // 生活动态
	PostTypeArticle = "article" // 文章
	PostTypeQA      = "qa"      // 问答
)

type Community struct {
	ent.Schema
}

func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.Int("author_id").
			Comment("作者ID"),
		field.Int("pet_id").
			Optional().
			Comment("宠物ID，可选，用于宠物相关的动态"),
		field.String("title").
			NotEmpty().
			Comment("帖子标题"),
		field.String("type").
			Default(PostTypeMoment).
			Comment("帖子类型").
			Validate(func(s string) error {
				switch s {
				case PostTypeMoment, PostTypeArticle, PostTypeQA:
					return nil
				default:
					return fmt.Errorf("invalid post type %q", s)
				}
			}),
		field.Text("content").
			NotEmpty().
			Comment("帖子内容"),
		field.Strings("images").
			Optional().
			Comment("图片URL列表"),
		field.Int("likes").
			Default(0).
			Comment("点赞数"),
		field.Int("comments").
			Default(0).
			Comment("评论数"),
		field.Int("views").
			Default(0).
			Comment("浏览数"),
		field.Bool("is_pinned").
			Default(false).
			Comment("是否置顶"),
		field.Strings("tags").
			Optional().
			Comment("帖子标签"),
		field.Time("created_at").
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}
