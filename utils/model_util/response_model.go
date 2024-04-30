package modelutil

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}

type ListResponse struct {
}

type PagedResponse struct {
	Message string `json:"message"`
	Data    []any  `json:"data"`
	// Paging
}
