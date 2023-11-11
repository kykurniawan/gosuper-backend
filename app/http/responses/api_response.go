package responses

type (
	PaginationMetaResponse struct {
		PerPage int `json:"perPage"`
		Page    int `json:"page"`
	}
)
