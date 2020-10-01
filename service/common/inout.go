package common

type List struct {
	Total   int
	Items   *[]interface{}
	CurPage int
	PerPage int
	HasMore bool
}

type ListOption struct {
	Page    int
	PerPage int
	Filters  map[string]interface{}
}

type SetOpParam struct {
	ID           string
	SetFieldName string
	Item         interface{}
}

func MakeTestListOption() (opt *ListOption) {
	return &ListOption{
		Page:    1,
		PerPage: 10,
		Filters:  map[string]interface{}{},
	}
}