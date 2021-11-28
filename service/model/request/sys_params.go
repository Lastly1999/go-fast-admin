package request

type PageInfo struct {
	PageSize int    `json:"pageSize"`
	Page     int    `json:"page"`
	KeyWords string `json:"keyWords"`
}
