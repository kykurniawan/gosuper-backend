package requests

type (
	IndexRequest struct {
		Page    int    `form:"page"`
		PerPage int    `form:"per_page"`
		Search  string `form:"search"`
	}
)
