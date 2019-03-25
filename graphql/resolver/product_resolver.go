package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/musobarlab/gorengan/modules/product/domain"
)

// ProductInput struct
type ProductInput struct {
	ID           string
	Name         string
	Quantity     int32
	Category     string
	CreatorID    string
	CreatorIP    string
	Created      *time.Time
	EditorID     string
	EditorIP     string
	LastModified *time.Time
	IsDeleted    *bool
	Deleted      *time.Time
}

// ProductResolver resolver
type ProductResolver struct {
	p *domain.Product
}

// ID function
func (p *ProductResolver) ID() graphql.ID {
	return graphql.ID(p.p.ID)
}

// Name function
func (p *ProductResolver) Name() string {
	return p.p.Name
}

// Quantity function
func (p *ProductResolver) Quantity() int32 {
	return int32(p.p.Quantity)
}

// Category function
func (p *ProductResolver) Category() *CategoryResolver {
	return &CategoryResolver{&p.p.Category}
}

// CreatorID function
func (p *ProductResolver) CreatorID() string {
	return p.p.CreatorID
}

// CreatorIP function
func (p *ProductResolver) CreatorIP() string {
	return p.p.CreatorIP
}

// Created function
func (p *ProductResolver) Created() *graphql.Time {
	return &graphql.Time{Time: p.p.Created}
}

// EditorID function
func (p *ProductResolver) EditorID() string {
	return p.p.EditorID
}

// EditorIP function
func (p *ProductResolver) EditorIP() string {
	return p.p.EditorIP
}

// LastModified function
func (p *ProductResolver) LastModified() *graphql.Time {
	return &graphql.Time{Time: p.p.LastModified}
}

// IsDeleted function
func (p *ProductResolver) IsDeleted() *bool {
	return &p.p.IsDeleted
}

// Deleted function
func (p *ProductResolver) Deleted() *graphql.Time {
	return &graphql.Time{Time: p.p.Deleted}
}
