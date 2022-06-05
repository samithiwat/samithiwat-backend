package model

type PaginationMetadata struct {
	ItemsPerPage int64
	ItemCount    int64
	TotalItem    int64
	CurrentPage  int64
	TotalPage    int64
}

func (p *PaginationMetadata) GetOffset() int64 {
	return (p.GetCurrentPage() - 1) * p.GetItemPerPage()
}

func (p *PaginationMetadata) GetItemPerPage() int64 {
	if p.ItemsPerPage < 10 {
		p.ItemsPerPage = 10
	}
	if p.ItemsPerPage > 100 {
		p.ItemsPerPage = 100
	}

	return p.ItemsPerPage
}

func (p *PaginationMetadata) GetCurrentPage() int64 {
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}
