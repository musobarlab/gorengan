package schema

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	categorySchema "github.com/musobarlab/gorengan/modules/category/graphql/schema"
	"github.com/musobarlab/gorengan/modules/product/domain"
)

// ProductSchemaInput struct
type ProductSchemaInput struct {
	ID           string
	Name         string
	Quantity     int32
	Category     string
	CreatorID    *string
	CreatorIP    *string
	Created      *time.Time
	EditorID     *string
	EditorIP     *string
	LastModified *time.Time
	IsDeleted    *bool
	Deleted      *time.Time
}

// ProductSchema resolver
type ProductSchema struct {
	Product *domain.Product
}

// ID function
func (p *ProductSchema) ID() graphql.ID {
	return graphql.ID(p.Product.ID)
}

// Name function
func (p *ProductSchema) Name() string {
	return p.Product.Name
}

// Quantity function
func (p *ProductSchema) Quantity() int32 {
	return int32(p.Product.Quantity)
}

// Category function
func (p *ProductSchema) Category() *categorySchema.CategorySchema {
	return &categorySchema.CategorySchema{Category: &p.Product.Category}
}

// CreatorID function
func (p *ProductSchema) CreatorID() string {
	return p.Product.CreatorID
}

// CreatorIP function
func (p *ProductSchema) CreatorIP() string {
	return p.Product.CreatorIP
}

// Created function
func (p *ProductSchema) Created() *graphql.Time {
	return &graphql.Time{Time: p.Product.Created}
}

// EditorID function
func (p *ProductSchema) EditorID() string {
	return p.Product.EditorID
}

// EditorIP function
func (p *ProductSchema) EditorIP() string {
	return p.Product.EditorIP
}

// LastModified function
func (p *ProductSchema) LastModified() *graphql.Time {
	return &graphql.Time{Time: p.Product.LastModified}
}

// IsDeleted function
func (p *ProductSchema) IsDeleted() *bool {
	return &p.Product.IsDeleted
}

// Deleted function
func (p *ProductSchema) Deleted() *graphql.Time {
	return &graphql.Time{Time: p.Product.Deleted}
}
