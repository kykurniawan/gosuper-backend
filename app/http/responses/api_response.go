package responses

type (
	PaginationMetaResponse struct {
		Total     int `json:"total"`
		PerPage   int `json:"perPage"`
		Page      int `json:"page"`
		TotalPage int `json:"totalPage"`
	}
)
