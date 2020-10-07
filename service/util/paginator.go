package util

type Paginator struct {
	Total   int
	CurPage int
	PerPage int
	HasMore bool
}

type PageOption struct {
	Page    int      `form:"page"`
	PerPage int      `form:"per_page"`
	Filters []string `form:"filters"`
}

type SetOpParam struct {
	Filters      []string
	SetFieldName string
	Item         interface{}
}
