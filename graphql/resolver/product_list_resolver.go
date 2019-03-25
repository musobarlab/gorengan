package resolver

// MetaResolver struct
type MetaResolver struct {
	page         *int32
	limit        *int32
	totalRecords *int32
	totalPages   *int32
}

// Page function
func (p *MetaResolver) Page() *int32 {
	return p.page
}

// Limit function
func (p *MetaResolver) Limit() *int32 {
	return p.limit
}

// TotalRecords function
func (p *MetaResolver) TotalRecords() *int32 {
	return p.totalRecords
}

// TotalPages function
func (p *MetaResolver) TotalPages() *int32 {
	return p.totalPages
}

// ProductListResolver struct
type ProductListResolver struct {
	products []*ProductResolver
	meta     *MetaResolver
}

// Products function
func (p *ProductListResolver) Products() []*ProductResolver {
	return p.products
}

// Meta function
func (p *ProductListResolver) Meta() *MetaResolver {
	return p.meta
}
