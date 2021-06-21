package xjutils

type Page struct {
	Current int         `json:"current""` //当前页
	Size    int         `json:"size"`     //每页条数
	Total   int64       `json:"total"`    //总条数
	Pages   int64       `json:"pages"`    //总页数
	Records interface{} `json:"records"`  //结果集
}

func (p *Page) SetTotal(total int64) {
	p.Total = total
	if total%int64(p.Size) == 0 {
		p.Pages = total / int64(p.Size)
	} else {
		p.Pages = total/int64(p.Size) + 1
	}
}

func (p *Page) SetPageSize(current, size int, total int64) {
	if current < 1 {
		current = 1
	}
	if size < 1 {
		size = 1
	}
	p.Current = current
	p.Size = size
	p.Total = total
	if total%int64(size) == 0 {
		p.Pages = total / int64(size)
	} else {
		p.Pages = total/int64(size) + 1
	}
}

func (p *Page) MakePage(records interface{}, current, size int, total int64) Page {
	if current < 1 {
		current = 1
	}
	if size < 1 {
		size = 1
	}
	p.Records = records
	p.Current = current
	p.Size = size
	p.Total = total
	if total%int64(size) == 0 {
		p.Pages = total / int64(size)
	} else {
		p.Pages = total/int64(size) + 1
	}
	return *p
}
