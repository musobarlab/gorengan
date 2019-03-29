package schema

// MetaResolver struct
type MetaResolver struct {
	PageField         *int32
	LimitField        *int32
	TotalRecordsField *int32
	TotalPagesField   *int32
}

// Page function
func (p *MetaResolver) Page() *int32 {
	return p.PageField
}

// Limit function
func (p *MetaResolver) Limit() *int32 {
	return p.LimitField
}

// TotalRecords function
func (p *MetaResolver) TotalRecords() *int32 {
	return p.TotalRecordsField
}

// TotalPages function
func (p *MetaResolver) TotalPages() *int32 {
	return p.TotalPagesField
}
