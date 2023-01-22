package entities

type PagingFilter struct {
	Page  int
	Limit int
	Sort  string
}

func (r *PagingFilter) GetPage() int {
	if r.Page == 0 {
		return 1
	}
	return r.Page
}

func (r *PagingFilter) GetLimit() int {
	if r.Limit == 0 {
		return 100
	}
	return r.Limit
}

func (r *PagingFilter) GetSort() string {
	if r.Sort == "" {
		return "timestamp DESC"
	}
	return r.Sort
}
