package pager

/*
	{
	 list: [],
	 current?: number,
	 pageSize?: number,
	 total?: number,
	}
*/
//Pager See Doc https://pro.ant.design/zh-CN/docs/request
type Pager struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"pageSize" form:"pageSize"`
	Total    int `json:"total" form:"-"`
}

func NewPager(page, pageSize int) *Pager {
	return &Pager{
		Current:  page,
		PageSize: pageSize,
		Total:    0,
	}
}

func (p *Pager) Offset() int {
	if p.Current <= 0 || p.Current >= 1000 {
		p.Current = 1
	}
	return (p.Current - 1) * p.Limit()
}

func (p *Pager) Limit() int {
	if p.PageSize <= 0 || p.PageSize >= 1000 {
		p.PageSize = 20
	}
	return p.PageSize
}
