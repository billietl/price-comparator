package dao

type Paginator struct {
	PageNumber  int
	PageSize    int
	HasPrevious bool
	HasNext     bool
}
