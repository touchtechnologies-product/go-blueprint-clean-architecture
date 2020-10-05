package util

type Paginator struct {
	Total   int
	CurPage int
	PerPage int
	HasMore bool
}

type PageOption struct {
	Page    int
	PerPage int
	Filters map[string]interface{}
}

type SetOpParam struct {
	ID           string
	SetFieldName string
	Item         interface{}
}
