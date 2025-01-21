package response

import (
	"pet/pkg/pager"
)

type List interface {
}

type WithListResp[T List] struct {
	Data []*T `json:"data"`
}

type WithListAndPageResp[T List] struct {
	Data []*T `json:"data"`
	pager.Pager
}

type DailyResp struct {
	Category  string `json:"category"`
	Value     int64  `json:"value"`
	CreatedAt int64  `json:"created_at"`
}
