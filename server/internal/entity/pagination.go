package entity

type Meta struct {
	Page    uint `json:"page"`
	PerPage uint `json:"per_page"`
	Total   uint `json:"total"`
}

func (m *Meta) SetTotal(total uint) {
	m.Total = total
}

func (m *Meta) ToSQL() (limit uint, offset uint) {
	if m.PerPage <= 0 {
		m.PerPage = 20
	}
	if m.Page <= 0 {
		m.Page = 1
	}
	limit = m.PerPage
	offset = (m.Page - 1) * limit
	return m.PerPage, offset
}
