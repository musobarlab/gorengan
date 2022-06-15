package schema

// ProductListResolver struct
type ProductListResolver struct {
	ProductsField []*ProductSchema
	MetaField     *MetaResolver
}

// Products function
func (p *ProductListResolver) Products() []*ProductSchema {
	return p.ProductsField
}

// Meta function
func (p *ProductListResolver) Meta() *MetaResolver {
	return p.MetaField
}
