package modelutil

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ListResponse struct {
}

type PagedResponse struct {
	Message string `json:"message"`
	Data    []any  `json:"data"`
	// Paging
}
