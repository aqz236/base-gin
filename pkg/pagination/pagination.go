package pagination

import "math"

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page" form:"page" binding:"min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}

// PageResponse 分页响应
type PageResponse struct {
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
	Data       interface{} `json:"data"`
}

// NewPageRequest 创建分页请求
func NewPageRequest(page, pageSize int) *PageRequest {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	return &PageRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

// GetOffset 获取偏移量
func (p *PageRequest) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// NewPageResponse 创建分页响应
func NewPageResponse(total int64, page, pageSize int, data interface{}) *PageResponse {
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &PageResponse{
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
		Data:       data,
	}
}
